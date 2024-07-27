package extras

import (
	"github.com/aronluigi/axon/pkg/extras/nodes/testa"
	"github.com/aronluigi/axon/pkg/node"
)

func GetNods() []*node.Node {
	return []*node.Node{
		testa.TestANode(),
	}
}
