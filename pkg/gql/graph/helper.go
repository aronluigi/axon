package graph

import (
	"reflect"

	"github.com/aronluigi/axon/pkg/gql/graph/model"
	"github.com/aronluigi/axon/pkg/node"
	"github.com/aronluigi/axon/pkg/state"
)

func nodeToModel(n *node.Node) *model.Node {
	return &model.Node{
		UUID:        n.UUID.String(),
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

func stateNodesToResp(sn []*state.Node) []*model.FlowStateNode {
	var res []*model.FlowStateNode

	for _, n := range sn {
		var (
			ip []*model.Port
			op []*model.Port
		)

		for _, nip := range n.Data.InputPorts {
			np := model.Port(nip)
			ip = append(ip, &np)
		}

		for _, nop := range n.Data.OutputPorts {
			np := model.Port(nop)
			op = append(op, &np)
		}

		position := model.FlowNodePosition(n.Position)
		res = append(res, &model.FlowStateNode{
			ID:       n.ID,
			Type:     n.Type,
			Position: &position,
			Data: &model.Node{
				UUID:        n.Data.UUID.String(),
				DisplayName: n.Data.DisplayName,
				PackageName: n.Data.PackageName,
				InputPorts:  ip,
				OutputPorts: op,
			},
		})
	}

	return res
}
