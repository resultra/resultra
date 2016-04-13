package values

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"resultra/datasheet/server/field"
)

const valSummaryCount string = "count"
const valSummarySum string = "sum"
const valSummaryAvg string = "average"

// DashboardValGrouping represents a grouping of field values for purposes of summarizing
// values in bar charts, lines charts, pie charts, and summary tables.
type ValSummary struct {

	// SummaryField is the field used to summarize values.
	SummarizeByField *datastore.Key

	// SummarizeValsWith configures how values from SummarizeByField are summarized.
	//
	// Depending on the data type of the field, different options are
	// available to summarize the values, including:
	//
	// Number: average, sum, min, max, stdDev
	// Date: count
	// Text: count
	// Bool: count
	SummarizeValsWith string
}

// DashboardValGroupingRef is an opaque reference to the dashboard value grouping parameters.
type ValSummaryRef struct {
	SummarizeByFieldRef field.FieldRef `json:"summarizeByFieldRef"`
	SummarizeValsWith   string         `json:"summarizeValsWith"`
}

type NewValSummaryParams struct {
	FieldID           string `json:"fieldID"`
	SummarizeValsWith string `json:"summarizeValsWith"`
}

func (valSummaryRef ValSummaryRef) toNewValSummaryParams() NewValSummaryParams {
	return NewValSummaryParams{
		FieldID:           valSummaryRef.SummarizeByFieldRef.FieldID,
		SummarizeValsWith: valSummaryRef.SummarizeValsWith}
}

func validateFieldTypeWithSummary(fieldType string, summarizeValsWith string) error {
	switch summarizeValsWith {
	case valSummaryCount:
		return nil
	case valSummarySum, valSummaryAvg:
		if fieldType != field.FieldTypeNumber {
			return fmt.Errorf("Invalid summary = %v for field type = %v", summarizeValsWith, fieldType)
		}
	case valGroupByDay:
		if fieldType != field.FieldTypeDate {
			return fmt.Errorf("Invalid summary = %v for field type = %v", summarizeValsWith, fieldType)
		}
	default:
		return fmt.Errorf("Invalid summary = %v for field type = %v", summarizeValsWith, fieldType)
	}

	return nil
}

func NewValSummary(appEngContext appengine.Context, params NewValSummaryParams) (*ValSummary, *ValSummaryRef, error) {

	fieldKey, fieldRef, fieldErr := field.GetExistingFieldRefAndKey(appEngContext, field.GetFieldParams{params.FieldID})
	if fieldErr != nil {
		return nil, nil, fmt.Errorf("NewValGrouping: Can't get field value grouping: datastore error = %v", fieldErr)
	}

	if summarizeErr := validateFieldTypeWithSummary(fieldRef.FieldInfo.Type, params.SummarizeValsWith); summarizeErr != nil {
		return nil, nil, fmt.Errorf("NewValGrouping: Invalid value summary: %v", summarizeErr)
	}

	valSummaryRef := ValSummaryRef{*fieldRef, params.SummarizeValsWith}
	valSummary := ValSummary{fieldKey, params.SummarizeValsWith}

	return &valSummary, &valSummaryRef, nil

}

func NewValSummaryFromRef(appEngContext appengine.Context, valSummaryRef ValSummaryRef) (*ValSummary, *ValSummaryRef, error) {
	return NewValSummary(appEngContext, valSummaryRef.toNewValSummaryParams())
}

func (valSummary ValSummary) GetValSummaryRef(appEngContext appengine.Context) (*ValSummaryRef, error) {

	fieldRef, fieldErr := field.GetFieldFromKey(appEngContext, valSummary.SummarizeByField)
	if fieldErr != nil {
		return nil, fmt.Errorf("GetValGroupingRef: Can't get field  for value grouping: datastore error = %v", fieldErr)
	}
	valSummaryRef := ValSummaryRef{*fieldRef, valSummary.SummarizeValsWith}

	return &valSummaryRef, nil

}
