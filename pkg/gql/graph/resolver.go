package graph

import (
	"github.com/aronluigi/axon/pkg/extras"
	"github.com/aronluigi/axon/pkg/node"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DefaultNodes []*node.Node
}

func NewResolver() *Resolver {
	return &Resolver{
		DefaultNodes: extras.GetNods(),
	}
}
