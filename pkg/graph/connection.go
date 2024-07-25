package graph

import "github.com/aronluigi/axon/pkg/node"

type Connection struct {
	FromNode *node.Node
	FromPort string
	ToNode   *node.Node
	ToPort   string
}

func NewConnection(fromNode *node.Node, fromPort string, toNode *node.Node, toPort string) *Connection {
	conn := &Connection{
		FromNode: fromNode,
		FromPort: fromPort,
		ToNode:   toNode,
		ToPort:   toPort,
	}

	return conn
}
