package node_test

import (
	"fmt"
	"testing"

	"github.com/aronluigi/axon/pkg/node"
	. "github.com/smartystreets/goconvey/convey"
)

func NodeTest(a string, b int, c *testing.T) (string, int, *testing.T, error) {
	return a, b, c, nil
}

type Input struct {
	FirstName string
	LastName  string
	T         *testing.T
}

type Out struct {
	Test string
}

func TestNode(t *testing.T) {
	Convey("node creation", t, func() {
		node, err := node.NewNode(func(in Input) (Out, error) {
			a, _, _, err := NodeTest(in.FirstName, 10, t)
			if err != nil {
				return Out{}, err
			}

			return Out{Test: a}, nil
		}, Input{}, Out{})

		So(err, ShouldBeNil)
		fmt.Println(node, "------------")

		_, err = node.Process(Input{FirstName: "test"})
		So(err, ShouldBeNil)
	})
}
