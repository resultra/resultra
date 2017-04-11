package values

import (
	"fmt"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/generic/uniqueID"
)

const ValSummaryCount string = "count"
const ValSummarySum string = "sum"
const ValSummaryAvg string = "average"

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
	SummarizeValsWith string `json:"summarizeValsWith"`
}

func (srcValSummary ValSummary) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*ValSummary, error) {

	destSummary := srcValSummary

	remappedFieldID, err := remappedIDs.GetExistingRemappedID(srcValSummary.SummarizeByFieldID)
	if err != nil {
		return nil, fmt.Errorf("ValSummary.Clone: %v", err)
	}
	destSummary.SummarizeByFieldID = remappedFieldID

	return &destSummary, nil

}

type NewValSummaryParams struct {
	FieldID           string `json:"fieldID"`
	SummarizeValsWith string `json:"summarizeValsWith"`
}

func validateFieldTypeWithSummary(fieldType string, summarizeValsWith string) error {
	switch summarizeValsWith {
	case ValSummaryCount:
		return nil
	case ValSummarySum, ValSummaryAvg:
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

func NewValSummary(params NewValSummaryParams) (*ValSummary, error) {

	summaryField, fieldErr := field.GetField(params.FieldID)
	if fieldErr != nil {
		return nil, fmt.Errorf("NewValGrouping: Can't get field value grouping: datastore error = %v", fieldErr)
	}

	if summarizeErr := validateFieldTypeWithSummary(summaryField.Type, params.SummarizeValsWith); summarizeErr != nil {
		return nil, fmt.Errorf("NewValGrouping: Invalid value summary: %v", summarizeErr)
	}

	valSummary := ValSummary{summaryField.FieldID, params.SummarizeValsWith}

	return &valSummary, nil

}

func (valSummary ValSummary) SummaryLabel() (string, error) {

	summaryField, fieldErr := field.GetField(valSummary.SummarizeByFieldID)
	if fieldErr != nil {
		return "", fmt.Errorf("SummaryLabel: Can't get field: %v", fieldErr)
	}

	switch valSummary.SummarizeValsWith {
	case ValSummaryCount:
		return fmt.Sprintf(`Count of '%v'`, summaryField.Name), nil
	case ValSummaryAvg:
		return fmt.Sprintf(`Average of '%v'`, summaryField.Name), nil
	case ValSummarySum:
		return fmt.Sprintf(`Sum of '%v'`, summaryField.Name), nil
	default:
		return "", fmt.Errorf("Unable to generate value label: unexpected summary = %v", valSummary.SummarizeValsWith)
	}
}
