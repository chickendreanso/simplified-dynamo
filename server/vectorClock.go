package main

import (
	pb "dynamoSimplified/pb"
)

// Find out the concurrent key value pair from the vector clocks
// Return the concurrent key value pair, or value array
func CompareVectorClocks(data []*pb.KeyValue) []*pb.KeyValue {
	concurrent := []*pb.KeyValue{}
	currentClock := data[0]
	data = data[1:]
	for len(data) > 0 {
		lt, mt := compareClock(currentClock, data[0])
		if !lt { //there are no elements that are less than the current clock, the first clock is ahead of current
			currentClock = data[0]
			data = data[1:]
		} else if mt { //there are elements that are mt and lt the current thus they are concurrent
			concurrent = append(concurrent, data[0])
			data = data[1:]
		} else { //no elements mt the current so it is behind the current
			data = data[1:]
		}
	}
	concurrent = append(concurrent, currentClock)
	current := concurrent[0]
	concurrent = concurrent[1:]
	i := 0
	for i < len(concurrent) {
		lt, mt := compareClock(current, concurrent[i])
		if !lt { //there are no elements that are less than the current clock, the first clock is ahead of current
			current = concurrent[i]
			concurrent = append(concurrent[:i], concurrent[i+1:]...)
			i = 0
		} else if mt { //there are elements that are mt and lt the current thus they are concurrent
			i += 1
		} else { //no elements mt the current so it is behind the current
			data = data[1:]
		}
	}
	concurrent = append(concurrent, current)
	return concurrent
}

func compareClock(KV1 *pb.KeyValue, KV2 *pb.KeyValue) (bool, bool) {
	lt := false                                            //check if there are clocks1 less than clocks2
	mt := false                                            //check if there are clocks1 greater than clocks2
	for node, clock1 := range KV1.VectorClock.Timestamps { //extract all nodes in compare clock
		clock2, found := KV2.VectorClock.Timestamps[node] //check against the clocks in the next machine
		if !found {
			clock2.ClokcVal = 0 // as if not written before.
		}
		if clock1.ClokcVal < clock2.ClokcVal {
			lt = true
		} else if clock1.ClokcVal > clock2.ClokcVal {
			mt = true
		}
	}
	return lt, mt
}
