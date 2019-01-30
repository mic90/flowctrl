package nodes

import (
	"flowctrl"
	"flowctrl/buffer"
	"flowctrl/buffer/adapter"
	"flowctrl/node"
	"time"
)

type Delay struct {
	info node.NodeInfo

	Input  *node.InputPort
	Output *node.OutputPort

	reader *adapter.Int8
	writer *adapter.Int8

	timer flowctrl.Waiter
}

func NewDelay(timer flowctrl.Waiter) *Delay {
	input := node.NewInputPort(buffer.Int8)
	output := node.NewOutputPort(buffer.Int8)
	info := node.NodeInfo{Name:"delay", Description:"passed data to another node with delay", Version:"1.0.0"}
	return &Delay{info, input, output, adapter.NewInt8(input), adapter.NewInt8(output), timer}
}

func (node *Delay) Info() node.NodeInfo {
	return node.info
}

func (node *Delay) Process() {
	if !node.timer.WaitFor(5*time.Second) {
		return
	}

	read := node.reader.Get()
	node.writer.Set(read)
}

func (node *Delay) Cleanup() {}