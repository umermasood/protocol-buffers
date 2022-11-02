package main

import (
	pb "08/proto-go/proto"
	"fmt"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func main() {
	fmt.Println(doSimple())
}
