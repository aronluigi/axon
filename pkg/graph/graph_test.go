package graph_test

import (
	"testing"

	"github.com/aronluigi/axon/pkg/graph"
	"github.com/aronluigi/axon/pkg/node"
	. "github.com/smartystreets/goconvey/convey"
)

type Input struct {
	LastName  *string
	FirstName string
}

type Output struct {
	Test string
}

func HandlerDummy(input Input) (Output, error) {
	return Output{Test: input.FirstName}, nil
}

func TestNode(t *testing.T) {
	Convey("Graph", t, func() {
		node_a, err := node.NewNode(HandlerDummy, Input{}, Output{})
		So(err, ShouldBeNil)
		node_b, err := node.NewNode(HandlerDummy, Input{}, Output{})
		So(err, ShouldBeNil)
		node_c, err := node.NewNode(HandlerDummy, Input{}, Output{})
		So(err, ShouldBeNil)

		g := graph.NewGraph()
		g.AddNode(node_a)
		g.AddNode(node_b)

		Convey("add connection", func() {
			Convey("positive", func() {
				err := g.AddConnection(node_a, "Test", node_b, "FirstName")
				So(err, ShouldBeNil)
			})

			Convey("negative", func() {
				Convey("node missing from graph", func() {
					err := g.AddConnection(node_a, "Test", node_c, "FirstName")
					So(err, ShouldBeError)
					So(err.Error(), ShouldContainSubstring, "not present")
				})
				Convey("same input output node", func() {
					err := g.AddConnection(node_a, "Test", node_a, "FirstName")
					So(err, ShouldBeError)
				})
				Convey("missing field", func() {
					err := g.AddConnection(node_a, "Test", node_b, "FirstNameA")
					So(err, ShouldBeError)
					err = g.AddConnection(node_a, "TestA", node_b, "FirstName")
					So(err, ShouldBeError)
				})
			})
		})
	})
}
