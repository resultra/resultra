package calcField

import (
	"appengine"
	"bytes"
	"encoding/json"
	"fmt"
	"resultra/datasheet/server/field"
	"strconv"
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

func FieldRefEqnNode(fieldID string) *EquationNode {
	// TODO - Verify the field with fieldID exists
	return &EquationNode{FieldID: fieldID}
}

func FuncEqnNode(funcName string, funcArgs []EquationNode) *EquationNode {
	// TODO - Verify the function with the given name exists
	return &EquationNode{FuncName: funcName, FuncArgs: funcArgs}
}

func NumberEqnNode(numberVal float64) *EquationNode {
	theVal := numberVal
	return &EquationNode{NumberVal: &theVal}
}

func TextEqnNode(textVal string) *EquationNode {
	theVal := textVal
	return &EquationNode{TextVal: &theVal}
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

// This function needs to use a FieldRefIDIndex instead of get for every field. However,
// since the Fields are currently stored without an ancestore, they are not consistent.
// TODO - migrate this function back to using FieldRefIDIndex once Fields are setup
// with an ancestore.
func (equation EquationNode) UserText(appEngContext appengine.Context) (string, error) {

	var resultBuf bytes.Buffer

	if len(equation.FuncName) > 0 {

		resultBuf.WriteString(equation.FuncName)
		resultBuf.WriteString("(")

		handledFirstArg := false
		for _, arg := range equation.FuncArgs {
			if handledFirstArg {
				// After the first argument has been processed, prefix
				// the next argument with a comma.
				resultBuf.WriteString(",")
			}
			handledFirstArg = true

			argEquationText, err := arg.UserText(appEngContext)
			if err != nil {
				return "", err
			} else {
				resultBuf.WriteString(argEquationText)
			}

		} // for each argument

		resultBuf.WriteString(")")
	} else if len(equation.FieldID) > 0 {

		fieldRef, err := field.GetField(appEngContext, equation.FieldID)
		if err != nil {
			return "", err
		} else {
			resultBuf.WriteString(fieldRef.RefName)
		}

	} else if equation.TextVal != nil {
		resultBuf.WriteString("\"")
		resultBuf.WriteString(*equation.TextVal)
		resultBuf.WriteString("\"")
	} else if equation.NumberVal != nil {
		resultBuf.WriteString(strconv.FormatFloat(*equation.NumberVal, 'f', 6, 64))
	}
	return resultBuf.String(), nil
}
