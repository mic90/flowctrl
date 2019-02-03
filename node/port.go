package node

import (
	"errors"
	"github.com/mic90/flowctrl/uuid"
)

var (
	ErrIncompatiblePorts = errors.New("incompatible ports provided")
)

type Direction int

const (
	DirectionInput  Direction = 0
	DirectionOutput Direction = 1
)

type Port interface {
	UUID() uuid.Value
	Direction() Direction
	Connectors() []*Connector
}
