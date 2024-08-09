package graph

import (
	"github.com/aronluigi/axon/pkg/extras"
	"github.com/aronluigi/axon/pkg/node"
	"github.com/aronluigi/axon/pkg/state"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DefaultNodes []*node.Node
	StateService *state.State
}

func NewResolver(stateService *state.State) *Resolver {
	return &Resolver{
		DefaultNodes: extras.GetNods(),
		StateService: stateService,
	}
}
