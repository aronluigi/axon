package node_test

import (
	"testing"

	"github.com/aronluigi/axon/pkg/node"
	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

type Input struct {
	LastName  *string
	FirstName string
	UUID      uuid.UUID
}

type Output struct {
	Test string
}

func HandlerDummy(input Input) (Output, error) {
	return Output{Test: input.FirstName}, nil
}

func TestNode(t *testing.T) {
	Convey("node creation", t, func() {
		node, err := node.NewNode(HandlerDummy, Input{}, Output{})
		So(err, ShouldBeNil)
		So(node.InputTypes["UUID"], ShouldHaveSameTypeAs, uuid.UUID)

		r, err := node.Process(Input{FirstName: "test"})
		So(err, ShouldBeNil)
		So(r, ShouldHaveSameTypeAs, Output{})
	})
}
