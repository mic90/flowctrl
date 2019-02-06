package node

type NodeInfo struct {
	Name        string
	Description string
	Version     string
}

type Node interface {
	Process()
	Cleanup()
	Info() NodeInfo
}
