package graph

import (
	"flowctrl/node"
	"flowctrl/uuid"
)

type Ordered struct {
	root    *node.Runner
	Runners []*node.Runner
}

func NewOrdered(root *node.Runner, mapped map[uuid.Value]*node.Runner) *Ordered {
	graph := &Ordered{root, make([]*node.Runner, 0, 1)}
	graph.build(mapped)
	return graph
}

func (ordered *Ordered) build(mapped map[uuid.Value]*node.Runner) {
	ordered.buildRecursive(ordered.root, mapped)
}

func (ordered *Ordered) buildRecursive(node *node.Runner, mapped map[uuid.Value]*node.Runner) {
	connectors := node.Connectors()
	for i := 0; i < len(connectors); i++ {
		baseRunner := mapped[connectors[i].FromUUID()]
		ordered.buildRecursive(baseRunner, mapped)
	}
	ordered.Runners = append(ordered.Runners, node)
}
