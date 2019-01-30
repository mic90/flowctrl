package nodes

import (
	"flowctrl/buffer"
	"flowctrl/buffer/adapter"
	"flowctrl/node"
)

type Pass struct {
	info node.NodeInfo

	Input  *node.InputPort
	Output *node.OutputPort

	reader *adapter.Int8
	writer *adapter.Int8
}

func NewPass(name string) *Pass {
	input := node.NewInputPort(buffer.Int8)
	output := node.NewOutputPort(buffer.Int8)
	info := node.NodeInfo{Name: name, Description: "passes input dat to another node", Version: "1.0.0"}
	return &Pass{info, input, output, adapter.NewInt8(input), adapter.NewInt8(output)}
}

func (node *Pass) Info() node.NodeInfo {
	return node.info
}

func (node *Pass) Process() {
	read := node.reader.Get()
	node.writer.Set(read)
}

func (node *Pass) Cleanup() {}