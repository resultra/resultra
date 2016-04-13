package calcField

import (
	"appengine"
	"encoding/json"
	"fmt"
	"resultra/datasheet/server/field"
)

// Parameters for creating a new calculated field. FieldEqn needs to be converted to
// JSON before being saved. TBD - Should the parameters instead be an equation in
// end-user format? If so, this code will need an update once equation parsing is
// done.
type NewCalcFieldParams struct {
	Name     string       `json:"name"`
	Type     string       `json:"type"`
	RefName  string       `json:"refName"`
	FieldEqn EquationNode `json:"fieldEqn"`
}

func encodeEqnJSONString(val interface{}) (string, error) {
	b, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("Error encoding calculated field equaton: %v", err)
	}
	return string(b), nil
}

func NewCalcField(appEngContext appengine.Context, calcFieldParams NewCalcFieldParams) (string, error) {

	jsonEncodeEqn, encodeErr := encodeEqnJSONString(calcFieldParams.FieldEqn)
	if encodeErr != nil {
		return "", encodeErr
	}

	// Create the actual field. All the parameters are the same as calcFieldParams, except
	// the equation which is encoded in JSON.
	newField := field.Field{
		Name:         calcFieldParams.Name,
		Type:         calcFieldParams.Type,
		RefName:      calcFieldParams.RefName,
		CalcFieldEqn: jsonEncodeEqn,
		IsCalcField:  true}

	return field.CreateNewFieldFromRawInputs(appEngContext, newField)
}
