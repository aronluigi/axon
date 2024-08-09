package state

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type State struct {
	filePath string
}

type NodePosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type NodePort struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	PkgPath string `json:"pkgPath"`
}

type NodeData struct {
	DisplayName string     `json:"displayName"`
	PackageName string     `json:"packageName"`
	InputPorts  []NodePort `json:"inputPorts"`
	OutputPorts []NodePort `json:"outputPorts"`
	UUID        uuid.UUID  `json:"uuid"`
}

type Node struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"`
	Data     NodeData     `json:"data"`
	Position NodePosition `json:"position"`
}

type store struct {
	Edges string  `json:"edges"`
	Nodes []*Node `json:"nodes"`
}

func NewState(filePath string) *State {
	return &State{
		filePath: filePath,
	}
}

func (o *State) Update(_ context.Context, nodes []*Node, edges string) error {
	s := store{Nodes: nodes, Edges: edges}

	content, err := json.Marshal(s)
	fmt.Println(string(content))
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	err = os.WriteFile(o.filePath, content, 0755)
	if err != nil {
		return fmt.Errorf("os.OpenFile: %w", err)
	}

	return nil
}
