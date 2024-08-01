package graph

import (
	"reflect"

	"github.com/aronluigi/axon/pkg/gql/graph/model"
	"github.com/aronluigi/axon/pkg/node"
)

func nodeToModel(n node.Node) *model.Node {
	return &model.Node{
		DisplayName: n.DisplayName,
		PackageName: n.PackageName,
		InputPorts:  getNodePortsToModel(n.InputTypes),
		OutputPorts: getNodePortsToModel(n.OutputTypes),
	}
}

func getNodePortsToModel(ports map[string]reflect.Type) []*model.Port {
	var res []*model.Port

	for k, v := range ports {
		res = append(res, &model.Port{
			Name:    k,
			Type:    v.String(),
			PkgPath: v.PkgPath(),
		})
	}

	return res
}
