package main

import (
	"context"
	"dynamoSimplified/config"
	hash "dynamoSimplified/hash"
	pb "dynamoSimplified/pb"
	"encoding/json"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"time"
)

// server is used to implement dynamo.KeyValueStoreServer.
// TODO: store data in memory first
type Server struct {
	pb.UnimplementedKeyValueStoreServer
	id             uint32
	addr           string
	mu             *sync.RWMutex // protects the following
	store          map[string]pb.KeyValue
	membershipList *pb.MembershipList
	vectorClocks   map[string]pb.VectorClock
}

type GetResponse struct {
	Message string `json:"message"`
}

func NewServer(addr string) *Server {
	return &Server{
		id:    hash.GenHash(addr),
		addr:  addr,
		mu:    &sync.RWMutex{},
		store: make(map[string]pb.KeyValue),
		membershipList: &pb.MembershipList{Nodes: []*pb.Node{
			&pb.Node{Id: hash.GenHash(addr), Address: addr, Timestamp: timestamppb.Now(), IsAlive: true},
		}},
		vectorClocks: make(map[string]pb.VectorClock),
	}
}

// Write implements dynamo.KeyValueStoreServer
func (s *Server) Write(ctx context.Context, in *pb.WriteRequest) (*pb.WriteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// this should be the ID of the current node
	nodeID := s.id
	key := in.KeyValue.Key

	// Use appropriate time for your use case
	// s.vectorClocks[key] = currentClock

	// Store the new value

	// Replicate write to W-1 other nodes (assuming the first write is the current node)

	// Make a gRPC call to Write method of the other node
	// ...
	if !in.IsReplica {
		var currentClock *pb.VectorClock
		// Update vector clock
		kv, found := s.store[key]
		if !found {
			currentClock = &pb.VectorClock{Timestamps: make(map[uint32]*pb.ClockStruct)}
			currentClock.Timestamps[nodeID] = &pb.ClockStruct{ClokcVal: 1, Timestamp: timestamppb.Now()}
		} else {
			currentClock = kv.VectorClock
			currentClock.Timestamps[nodeID].ClokcVal += 1
			currentClock.Timestamps[nodeID].Timestamp = timestamppb.Now()
		}
		in.KeyValue.VectorClock = currentClock
		s.store[key] = *in.KeyValue
		value, _ := s.store[key]
		replicaResult := SendRequestToReplica(kv, s.membershipList.Nodes, config.WRITE, s.addr) //replica result is an array
		result := append(replicaResult, &value)
		return &pb.WriteResponse{KeyValue: result, Success: true}, nil
		//TODO: implement timeout when waited to long to get write success. or detect write failure
	} else {
		value, _ := s.store[key]
		return &pb.WriteResponse{KeyValue: []*pb.KeyValue{&value}, Success: true}, nil
	}

	return &pb.WriteResponse{Success: true}, nil
}

func (s *Server) HintedHandoffWriteRequest(ctx context.Context, in *pb.HintedHandoffWriteRequest) (*pb.WriteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// TODO
	return &pb.WriteResponse{Success: true}, nil
}

// Read implements dynamo.KeyValueStoreServer
func (s *Server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	log.Printf("read request received for %v", in.Key)

	key := in.Key
	kv := pb.KeyValue{Key: key}

	value, ok := s.store[key]
	if !ok {
		return &pb.ReadResponse{Success: false, Message: "Key not found"}, nil
	}
	if !in.IsReplica {
		replicaResult := SendRequestToReplica(kv, s.membershipList.Nodes, config.READ, s.addr)
		result := append(replicaResult, &value) //contains the addresses of all stores
		//compare vector clocks
		result = CompareVectorClocks(result)
		return &pb.ReadResponse{KeyValue: result, Success: true}, nil
	}

	return &pb.ReadResponse{KeyValue: []*pb.KeyValue{&value}, Success: true}, nil
}

// Join implements dynamo.KeyValueStoreServer
func (s *Server) Join(ctx context.Context, in *pb.Node) (*pb.MembershipList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Join request received from %v", in.Address)

	// Update membership list
	s.membershipList.Nodes = append(s.membershipList.Nodes, in)

	// Send membership list to joining node
	return &pb.MembershipList{Nodes: s.membershipList.Nodes}, nil
}

// Gossip implements dynamo.KeyValueStoreServer
func (s *Server) Gossip(ctx context.Context, in *pb.GossipMessage) (*pb.GossipAck, error) {
	// Update nodes based on the received gossip message
	s.mu.Lock()
	s.membershipList = ReconcileMembershipList(s.membershipList, in.MembershipList)
	log.Println("Membership list:")
	for _, node := range s.membershipList.Nodes {
		if node.IsAlive {
			log.Printf("Node %v is alive", node.Address)
		} else {
			log.Printf("Node %v is dead", node.Address)
		}
	}
	s.mu.Unlock()

	return &pb.GossipAck{Success: true}, nil
}

func ReconcileMembershipList(list1 *pb.MembershipList, list2 *pb.MembershipList) *pb.MembershipList {
	var mp = make(map[uint32]*pb.Node)

	for _, node := range list1.Nodes {
		mp[node.Id] = node
	}

	for _, node := range list2.Nodes {
		if _, ok := mp[node.Id]; !ok {
			mp[node.Id] = node
		} else {
			if mp[node.Id].Timestamp.Seconds < node.Timestamp.Seconds {
				mp[node.Id] = node
			}
		}
	}

	var newList = pb.MembershipList{Nodes: []*pb.Node{}}
	for _, node := range mp {
		newList.Nodes = append(newList.Nodes, node)
	}

	return &pb.MembershipList{Nodes: newList.Nodes}
}

// create a method to periodically send gossip message to other nodes
func (s *Server) SendGossip(ctx context.Context) {
	for {
		// randomly pick one other node from membership list
		// send gossip to that node
		s.mu.RLock()
		targetNode := s.membershipList.Nodes[rand.Intn(len(s.membershipList.Nodes))]
		s.mu.RUnlock()
		if targetNode.Address == s.addr {
			continue
		}

		// create grpc client
		conn, err := grpc.Dial(targetNode.Address, grpc.WithInsecure())
		if err != nil {
			log.Printf("fail to dial: %v", err)
			// update membership list to change isAlive to false
			s.mu.Lock()
			for _, node := range s.membershipList.Nodes {
				if node.Address == targetNode.Address {
					node.IsAlive = false
					node.Timestamp = timestamppb.Now()
					s.mu.Unlock()
					break
				}
			}
			continue
		}
		defer conn.Close()

		client := pb.NewKeyValueStoreClient(conn)

		// send gossip message
		s.mu.RLock()
		membershipList := s.membershipList
		s.mu.RUnlock()
		resp, err := client.Gossip(ctx, &pb.GossipMessage{MembershipList: membershipList})
		if err != nil {
			log.Printf("fail to send gossip to %v", targetNode.Address)
			// update membership list to change isAlive to false
			s.mu.Lock()
			for _, node := range s.membershipList.Nodes {
				if node.Address == targetNode.Address {
					node.IsAlive = false
					node.Timestamp = timestamppb.Now()
					s.mu.Unlock()
					break
				}
			}
			continue
		}

		if resp.Success {
			s.mu.Lock()
			for _, node := range s.membershipList.Nodes {
				if node.Address == targetNode.Address {
					node.IsAlive = true
					node.Timestamp = timestamppb.Now()
					s.mu.Unlock()
					break
				}
			}
		}
		time.Sleep(time.Second * 5)
	}
}

func (s *Server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{}, nil
}

func getSeedNodeAddr(webclient string) string {
	// call get rest api to webclient address
	// get the seed node address
	// Make a GET request to the API. {message: "addr"}
	resp, err := http.Get(webclient)
	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	log.Printf("Response body: %v", string(body))

	var data GetResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal the response body: %v", err)
	}

	return data.Message
}

var (
	addr      = flag.String("addr", "127.0.0.1:50051", "the addr to serve on")
	webclient = flag.String("webclient", "", "the addr of the seed node")
	sleep     = flag.Duration("sleep", time.Second*5, "duration between changes in health")

	system = "" // empty string represents the health of the system
)

func main() {
	flag.Parse()

	// create a server
	server := NewServer(*addr)

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register the server with the gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterKeyValueStoreServer(grpcServer, server)

	log.Printf("Server listening at %v", lis.Addr())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	addrToJoin := getSeedNodeAddr(*webclient)
	// join the seed node if not empty
	if addrToJoin != "" {
		// create grpc client
		conn, err := grpc.Dial(addrToJoin, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()

		client := pb.NewKeyValueStoreClient(conn)

		// join the seed node
		resp, err := client.Join(context.Background(), &pb.Node{Id: hash.GenHash(*addr), Address: *addr, Timestamp: timestamppb.Now(), IsAlive: true})
		if err != nil {
			log.Fatalf("%d failed to join %d at %v, retrying...", server.id, addrToJoin, addrToJoin)
		} else {
			log.Printf("%d joined successfully", server.id)
		}

		server.mu.Lock()
		server.membershipList = resp
		server.mu.Unlock()
	}

	log.Printf("Starting gossip...")

	// start gossiping
	go server.SendGossip(context.Background())
	for {
	}
}
