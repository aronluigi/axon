package node

import (
	"errors"
	"fmt"
	"reflect"
)

type NodeInterface interface {
	Process(in any) (out any, err error)
}

type Node struct {
	Function         func(any) (any, error)
	Input            any
	InputDefinition  NodeIODef
	Output           any
	OutputDefinition NodeIODef
}

type NodeIODef struct {
	Fields map[string]string
	Type   string
}

func getIODefinition(obj any) (NodeIODef, error) {
	def := NodeIODef{}
	instanceType := reflect.TypeOf(obj)

	if instanceType.Kind() == reflect.Pointer {
		return def, errors.New("input/output can't be a pointer")
	}

	fields := make(map[string]string)
	for i := 0; i < instanceType.NumField(); i++ {
		field := instanceType.Field(i)
		fields[field.Name] = field.Type.String()
	}

	def.Type = instanceType.Name()
	def.Fields = fields

	return def, nil
}

func NewNode[T any, K any](fn func(T) (K, error), input T, output K) (NodeInterface, error) {
	wrapper := func(in any) (any, error) {
		if v, ok := (in).(T); ok {
			return fn(v)
		}

		return nil, fmt.Errorf("input is not of type: %T but %T", input, in)
	}

	inputDef, err := getIODefinition(input)
	if err != nil {
		return nil, fmt.Errorf("input definition: getIODefinition: %w", err)
	}

	outputDef, err := getIODefinition(output)
	if err != nil {
		return nil, fmt.Errorf("output definition: getIODefinition: %w", err)
	}

	node := &Node{
		Function:         wrapper,
		Input:            input,
		InputDefinition:  inputDef,
		Output:           output,
		OutputDefinition: outputDef,
	}

	return node, nil
}

func (o *Node) Process(in any) (out any, err error) {
	return o.Function(in)
}
