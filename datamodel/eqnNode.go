package datamodel

import (
	"encoding/json"
	"fmt"
)

type EquationNode struct {

	// Internal/root nodes - functions which point to arguments of
	// other functions and/or leaf nodes.
	FuncName string         `json:"funcName,omitempty"`
	FuncArgs []EquationNode `json:"funcArgs,omitempty"`

	// Leaf nodes - values
	FieldID string `json:"fieldID,omitempty"`

	// Literal values -  Use pointers to the values, which allows the use of
	// empty strings or zero numbers in the values.
	// If not using a string pointer, an empty string won't be
	// omitted from JSON encoding output.
	TextVal   *string  `json:"textVal,omitempty"`
	NumberVal *float64 `json:"numberVal,omitempty"`
}

func fieldRefEqnNode(fieldID string) *EquationNode {
	// TODO - Verify the field with fieldID exists
	return &EquationNode{FieldID: fieldID}
}

func funcEqnNode(funcName string, funcArgs []EquationNode) *EquationNode {
	// TODO - Verify the function with the given name exists
	return &EquationNode{FuncName: funcName, FuncArgs: funcArgs}
}

func decodeEquation(encodedEqn string) (*EquationNode, error) {

	decodedEqnNode := EquationNode{}
	encodedBytes := []byte(encodedEqn)
	if err := json.Unmarshal(encodedBytes, &decodedEqnNode); err != nil {
		return nil, fmt.Errorf("Failure decoding equation: encoded eqn = %v, decode error=%v",
			encodedEqn, err)
	} else {
		return &decodedEqnNode, nil
	}
}
