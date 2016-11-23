package calcField

import (
	"encoding/json"
	"fmt"
	"resultra/datasheet/server/generic"
)

type EquationNode struct {

	// Internal/root nodes - functions which point to arguments of
	// other functions and/or leaf nodes.
	FuncName string          `json:"funcName,omitempty"`
	FuncArgs []*EquationNode `json:"funcArgs,omitempty"`

	// Leaf nodes - values
	FieldID  string `json:"fieldID,omitempty"`
	GlobalID string `json:"globalID,omitempty"`

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

func GlobalRefEqnNode(globalID string) *EquationNode {
	return &EquationNode{GlobalID: globalID}
}

func FuncEqnNode(funcName string, funcArgs []*EquationNode) *EquationNode {
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

// Traverse the equation node tree, mapping the field and global IDs from source to the re-mapped IDs. This
// is used when copying an existing database to a template or copying a template to create a new database.
func remapEquationToClonedIDs(remappedIDs map[string]string, eqnNode *EquationNode) error {

	if len(eqnNode.FieldID) > 0 {

		remappedFieldID, idFound := remappedIDs[eqnNode.FieldID]
		if !idFound {
			return fmt.Errorf("RemapEquationNodeToClonedIDs: Can't find remapped ID for field ID = %v", eqnNode.FieldID)
		}

		eqnNode.FieldID = remappedFieldID
	}

	if len(eqnNode.GlobalID) > 0 {

		remappedGlobalID, idFound := remappedIDs[eqnNode.GlobalID]
		if !idFound {
			return fmt.Errorf("RemapEquationNodeToClonedIDs: Can't find remapped ID for global ID = %v", eqnNode.GlobalID)
		}
		eqnNode.GlobalID = remappedGlobalID

	}

	for eqnNode.FuncArgs != nil {
		for _, currArg := range eqnNode.FuncArgs {
			if err := remapEquationToClonedIDs(remappedIDs, currArg); err != nil {
				return fmt.Errorf("RemapEquationNodeToClonedIDs: %v", err)
			}

		}
	}

	return nil
}

func CloneEquation(remappedIDs map[string]string, encodedEqn string) (string, error) {

	rootEqnNode, err := decodeEquation(encodedEqn)
	if err != nil {
		return "", fmt.Errorf("CloneEquation: failure decoding source equation: %v", err)
	}

	if err := remapEquationToClonedIDs(remappedIDs, rootEqnNode); err != nil {
		return "", fmt.Errorf("CloneEquation: %v", err)
	}

	encodedClonedEqn, encodeErr := generic.EncodeJSONString(rootEqnNode)
	if encodeErr != nil {
		return "", encodeErr
	}

	return encodedClonedEqn, nil

}
