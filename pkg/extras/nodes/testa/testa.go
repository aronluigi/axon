package testa

import (
	"github.com/aronluigi/axon/pkg/node"
	"github.com/google/uuid"
)

type Input struct {
	LastName  *string
	FirstName string
	UUID      uuid.UUID
	Test      []string
}

type Output struct {
	Test string
}

func handler(input Input) (Output, error) {
	return Output{Test: input.FirstName}, nil
}

func TestANode() *node.Node {
	n, err := node.NewNode("Test A Node", "axon.extras", handler, Input{}, Output{})
	if err != nil {
		panic(err)
	}

	return n
}
