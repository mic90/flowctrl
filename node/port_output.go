package node

import (
	"github.com/mic90/flowctrl/buffer"
	"github.com/mic90/flowctrl/uuid"
)

type OutputPort struct {
	*buffer.Const
	uuid       uuid.Value
	direction  Direction
	connectors []*Connector
}

func NewOutputPort(_type buffer.Type) *OutputPort {
	return &OutputPort{buffer.NewConst(_type), uuid.New(), DirectionOutput, make([]*Connector, 0, 1)}
}

func (p *OutputPort) Write(data []byte) (int, error) {
	written, err := p.Const.Write(data)
	if err != nil {
		return written, err
	}
	for i := 0; i < len(p.connectors); i++ {
		conn := p.connectors[i]
		conn.Notify()
	}
	return written, nil
}

func (p *OutputPort) UUID() uuid.Value {
	return p.uuid
}

func (p *OutputPort) Direction() Direction {
	return p.direction
}

func (p *OutputPort) Connectors() []*Connector {
	return p.connectors
}

func (p *OutputPort) Copy(other *InputPort) (int, error) {
	return p.Const.Copy(other.Const)
}

func (p *OutputPort) Connect(inputs ...*InputPort) {
	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		if err := p.connectInput(input); err != nil {
			panic(err)
		}
	}
}

func (p *OutputPort) connectInput(input *InputPort) error {
	if input.connector != nil {
		return ErrIncompatiblePorts
	}
	connector := NewConnector(p, input)
	// add connector to the destination port
	input.connector = connector
	// add connector to output port to enable notifiers
	p.connectors = append(p.connectors, connector)
	return nil
}
