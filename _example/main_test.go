package main

import (
	"flowctrl"
	"flowctrl/_example/nodes"
	"flowctrl/buffer/adapter"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	b.ReportAllocs()

	nodeA := nodes.NewPass("nodeA")
	nodeB := nodes.NewPass("nodeB")
	nodeC := nodes.NewPass("nodeC")
	nodeD := nodes.NewPass("nodeD")
	nodeE := nodes.NewPass("nodeE")
	nodeF := nodes.NewPass("nodeF")
	nodeG := nodes.NewPass("nodeG")

	initWriter := adapter.NewInt8(nodeA.Input)
	initWriter.Set(123)

	// setup connections
	nodeA.Output.Connect(nodeB.Input)
	nodeB.Output.Connect(nodeC.Input)
	nodeC.Output.Connect(nodeD.Input)
	nodeD.Output.Connect(nodeE.Input)
	nodeE.Output.Connect(nodeF.Input)

	graph := flowctrl.New(nodeD, nodeA, nodeC, nodeB, nodeE, nodeF, nodeG)
	runner := flowctrl.NewSerialRunner(graph)

	for i:=0; i<b.N; i++ {
		err := runner.Process()
		if err != nil {
			b.Fail()
		}
	}
}