package nodes

import (
	"context"
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
	info := node.NodeInfo{name, "description", "1.0.0"}
	return &Pass{info, input, output, adapter.NewInt8(input), adapter.NewInt8(output)}
}

func (node *Pass) Info() node.NodeInfo {
	return node.info
}

func (node *Pass) Process(ctx context.Context) {
	read := node.reader.Get()
	node.writer.Set(read)
}

func (node *Pass) Cleanup() {}