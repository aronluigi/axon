package main

import "github.com/aronluigi/axon/pkg/gql"

//	type Flow struct {
//		Graph       graph.GraphInterface
//		Connections connection.Connection
//	}
//
//	type FlowInterface interface {
//		AddNode(*node.Node)
//	}
//
//	func NewFlow() FlowInterface {
//		return &Flow{
//			Graph: graph.NewGraph(),
//		}
//	}
//
//	func (o *Flow) AddNode(n *node.Node) {
//		o.Graph.AddNode(n)
//	}
func main() {
	gql.StartServer(8888)
}
