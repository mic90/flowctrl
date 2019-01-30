package main

import (
	"flowctrl"
	"flowctrl/_example/nodes"
	"flowctrl/buffer/adapter"
	"log"
)

func main() {
	// create nodes
	nodeA := nodes.NewPass("nodeA")
	nodeB := nodes.NewPass("nodeB")
	nodeC := nodes.NewPass("nodeC")
	nodeD := nodes.NewPass("nodeD")
	nodeE := nodes.NewPass("nodeE")
	nodeF := nodes.NewPass("nodeF")
	nodeG := nodes.NewPass("nodeG")
	delay := nodes.NewDelay(flowctrl.NewTimer())

	// init 'nodeA' with some value, prepare final reader on nodeF output
	initWriter := adapter.NewInt8(nodeA.Input)
	initWriter.Set(123)
	endReader := adapter.NewInt8(nodeF.Output)

	// setup connections
	nodeA.Output.Connect(nodeB.Input)
	nodeB.Output.Connect(nodeC.Input)
	nodeC.Output.Connect(delay.Input)
	delay.Output.Connect(nodeD.Input)
	nodeD.Output.Connect(nodeE.Input)
	nodeE.Output.Connect(nodeF.Input)
	// nodeG is not connected, it will be processed as separate graph

	graph := flowctrl.New(nodeD, nodeA, nodeC, nodeB, nodeE, nodeF, nodeG, delay)
	runner := flowctrl.NewSerialRunner(graph)

	log.Println("Flow started")
	for {
		err := runner.Process()
		if err != nil {
			panic(err)
		}
		// wait for delayed node to propagate data
		if endReader.Get() != 0 {
			break
		}
	}
	log.Println("Final read:", endReader.Get())
}
