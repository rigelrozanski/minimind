package main

import "fmt"

func main() {

	els := make(map[uint64]Element)
	conns := make(map[uint64]ElementConnection)

	conns[0] = NewElementConnection(0, 0, 1, "forward")
	conns[1] = NewElementConnection(1, 0, 2, "backward")

	els[0] = NewElement(0, "some element sick cool sweet element woohoo", []uint64{0, 1})
	els[1] = NewElement(1, "some extra element", []uint64{0})
	els[2] = NewElement(2, "some other extra element", []uint64{1})

	var outputStr string
	outputStr = els[0].RenderBordered()

	fmt.Println(outputStr)
}

// rendering strategy
// determine the biggest X, biggest Y among a layer,
