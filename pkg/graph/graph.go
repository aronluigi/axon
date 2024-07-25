package graph

import (
	"fmt"

	"github.com/aronluigi/axon/pkg/node"
	"github.com/google/uuid"
)

type Graph struct {
	Nodes       map[uuid.UUID]*node.Node
	Connections []*Connection
}

func NewGraph() GraphInterface {
	return &Graph{
		Nodes:       make(map[uuid.UUID]*node.Node),
		Connections: []*Connection{},
	}
}

type GraphInterface interface {
	AddNode(*node.Node)
	AddConnection(*Connection) error
}

func (o *Graph) AddNode(n *node.Node) {
	o.Nodes[n.UUID] = n
}

func (o *Graph) AddConnection(con *Connection) error {
	if con.FromNode.UUID == con.ToNode.UUID {
		return fmt.Errorf("FromNode [%s] is the same as ToNode", con.FromNode.UUID.String())
	}

	if _, ok := o.Nodes[con.FromNode.UUID]; !ok {
		return fmt.Errorf("FromNode [%s] is not present in the graph", con.FromNode.UUID.String())
	}

	if _, ok := o.Nodes[con.ToNode.UUID]; !ok {
		return fmt.Errorf("ToNode [%s] is not present in the graph", con.ToNode.UUID.String())
	}

	if _, ok := o.Nodes[con.FromNode.UUID].OutputTypes[con.FromPort]; !ok {
		return fmt.Errorf("FromPort field [%s] is not present in FromNode OutputTypes", con.FromPort)
	}

	if _, ok := o.Nodes[con.ToNode.UUID].InputTypes[con.ToPort]; !ok {
		return fmt.Errorf("ToPort field [%s] is not present in ToNode InputTypes", con.ToPort)
	}

	o.Connections = append(o.Connections, con)

	return nil
}

func (o *Graph) Execute() {
}
