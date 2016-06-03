package values

import (
	"appengine"
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
	SummarizeByFieldID string `json:"summarizeByFieldID"`

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

type NewValSummaryParams struct {
	FieldID           string `json:"fieldID"`
	SummarizeValsWith string `json:"summarizeValsWith"`
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
		if fieldType != field.FieldTypeTime {
			return fmt.Errorf("Invalid summary = %v for field type = %v", summarizeValsWith, fieldType)
		}
	default:
		return fmt.Errorf("Invalid summary = %v for field type = %v", summarizeValsWith, fieldType)
	}

	return nil
}

func NewValSummary(appEngContext appengine.Context, params NewValSummaryParams) (*ValSummary, error) {

	summaryField, fieldErr := field.GetField(appEngContext, params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewValGrouping: Can't get field value grouping: datastore error = %v", fieldErr)
	}

	if summarizeErr := validateFieldTypeWithSummary(summaryField.Type, params.SummarizeValsWith); summarizeErr != nil {
		return nil, fmt.Errorf("NewValGrouping: Invalid value summary: %v", summarizeErr)
	}

	valSummary := ValSummary{summaryField.FieldID, params.SummarizeValsWith}

	return &valSummary, nil

}
