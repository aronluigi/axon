package graph

import (
	"fmt"

	"github.com/aronluigi/axon/pkg/node"
	"github.com/google/uuid"
)

type GraphNodeConnection struct {
	FromNode *node.Node
	FromPort string
	ToNode   *node.Node
	ToPort   string
}

type Graph struct {
	Nodes       map[uuid.UUID]*node.Node
	Connections []*GraphNodeConnection
}

func NewGraph() GraphInterface {
	return &Graph{
		Nodes:       make(map[uuid.UUID]*node.Node),
		Connections: []*GraphNodeConnection{},
	}
}

type GraphInterface interface {
	AddNode(...*node.Node)
	AddConnection(fromNode *node.Node, fromPort string, toNode *node.Node, toPort string) error
}

func (o *Graph) AddNode(nodes ...*node.Node) {
	for _, node := range nodes {
		if _, ok := o.Nodes[node.UUID]; ok {
			continue
		}

		o.Nodes[node.UUID] = node
	}
}

func (o *Graph) AddConnection(fromNode *node.Node, fromPort string, toNode *node.Node, toPort string) error {
	if fromNode.UUID == toNode.UUID {
		return fmt.Errorf("FromNode [%s] is the same as ToNode", fromNode.UUID.String())
	}

	if _, ok := o.Nodes[fromNode.UUID]; !ok {
		return fmt.Errorf("FromNode [%s] is not present in the graph", fromNode.UUID.String())
	}

	if _, ok := o.Nodes[toNode.UUID]; !ok {
		return fmt.Errorf("ToNode [%s] is not present in the graph", toNode.UUID.String())
	}

	if _, ok := o.Nodes[fromNode.UUID].OutputTypes[fromPort]; !ok {
		return fmt.Errorf("FromPort field [%s] is not present in FromNode OutputTypes", fromPort)
	}

	if _, ok := o.Nodes[toNode.UUID].InputTypes[toPort]; !ok {
		return fmt.Errorf("ToPort field [%s] is not present in ToNode InputTypes", toPort)
	}

	o.Connections = append(o.Connections, &GraphNodeConnection{
		FromNode: fromNode,
		ToNode:   toNode,
		FromPort: fromPort,
		ToPort:   toPort,
	})

	o.createTree()
	return nil
}

func (o *Graph) Execute() {
}

func (o *Graph) createTree() {
	fmt.Println("---------")
	fmt.Println(len(o.Connections))
	for _, node := range o.Connections {
		fmt.Println(node)
	}
}
