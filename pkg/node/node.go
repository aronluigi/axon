package node

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type Node struct {
	UUID        uuid.UUID
	Function    func(any) (any, error)
	Input       any
	InputTypes  map[string]reflect.Type
	Output      any
	OutputTypes map[string]reflect.Type
}

func getIODefinition(obj any) (map[string]reflect.Type, error) {
	instanceType := reflect.TypeOf(obj)
	fields := make(map[string]reflect.Type)

	if instanceType.Kind() == reflect.Pointer {
		return fields, errors.New("input/output can't be a pointer")
	}

	for i := 0; i < instanceType.NumField(); i++ {
		field := instanceType.Field(i)
		fields[field.Name] = field.Type
	}

	return fields, nil
}

func NewNode[T any, K any](fn func(T) (K, error), input T, output K) (*Node, error) {
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

	uid, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("UUID: %w", err)
	}

	node := &Node{
		UUID:        uid,
		Function:    wrapper,
		Input:       input,
		InputTypes:  inputDef,
		Output:      output,
		OutputTypes: outputDef,
	}

	return node, nil
}

func (o *Node) Process(in any) (out any, err error) {
	return o.Function(in)
}
