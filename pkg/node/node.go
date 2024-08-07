package node

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type Node struct {
	Input       any
	Output      any
	Function    func(any) (any, error)
	InputTypes  map[string]reflect.Type
	OutputTypes map[string]reflect.Type
	DisplayName string    `json:"displayName"`
	PackageName string    `json:"packageName"`
	UUID        uuid.UUID `json:"uuid"`
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

func NewNode[T any, K any](displayName, packageName string, fn func(T) (K, error), input T, output K) (*Node, error) {
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

	uid := uuid.NewSHA1(uuid.NameSpaceX500, []byte(displayName+"-"+packageName))

	node := &Node{
		UUID:        uid,
		Function:    wrapper,
		Input:       input,
		InputTypes:  inputDef,
		Output:      output,
		OutputTypes: outputDef,
		DisplayName: displayName,
		PackageName: packageName,
	}

	return node, nil
}

func (o *Node) Process(in any) (out any, err error) {
	return o.Function(in)
}
