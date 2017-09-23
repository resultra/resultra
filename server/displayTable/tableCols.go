package displayTable

import (
	"fmt"
	"resultra/datasheet/server/displayTable/columns/attachment"
	"resultra/datasheet/server/displayTable/columns/checkBox"
	"resultra/datasheet/server/displayTable/columns/comment"
	"resultra/datasheet/server/displayTable/columns/datePicker"
	"resultra/datasheet/server/displayTable/columns/emailAddr"
	"resultra/datasheet/server/displayTable/columns/file"
	"resultra/datasheet/server/displayTable/columns/formButton"
	"resultra/datasheet/server/displayTable/columns/image"
	"resultra/datasheet/server/displayTable/columns/note"
	"resultra/datasheet/server/displayTable/columns/numberInput"
	"resultra/datasheet/server/displayTable/columns/progress"
	"resultra/datasheet/server/displayTable/columns/rating"
	"resultra/datasheet/server/displayTable/columns/socialButton"
	"resultra/datasheet/server/displayTable/columns/tag"
	"resultra/datasheet/server/displayTable/columns/textInput"
	"resultra/datasheet/server/displayTable/columns/toggle"
	"resultra/datasheet/server/displayTable/columns/urlLink"
	"resultra/datasheet/server/displayTable/columns/userSelection"
)

type TableColsInfo []interface{}
type TableColsByID map[string]interface{}

func getTableCols(parentTableID string) (TableColsInfo, TableColsByID, error) {

	tableColData := TableColsInfo{}
	tableColsByID := TableColsByID{}

	numberInputCols, err := numberInput.GetNumberInputs(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range numberInputCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	progressCols, err := progress.GetProgressIndicators(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range progressCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	textInputCols, err := textInput.GetTextInputs(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range textInputCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	datePickerCols, err := datePicker.GetDatePickers(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range datePickerCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	checkBoxCols, err := checkBox.GetCheckBoxes(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range checkBoxCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	ratingCols, err := rating.GetRatings(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range ratingCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	toggleCols, err := toggle.GetToggles(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range toggleCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	userSelectionCols, err := userSelection.GetUserSelections(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range userSelectionCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	noteCols, err := note.GetNotes(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range noteCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	commentCols, err := comment.GetComments(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range commentCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	attachCols, err := attachment.GetAttachments(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range attachCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	buttonCols, err := formButton.GetButtons(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range buttonCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	socialButtonCols, err := socialButton.GetSocialButtons(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range socialButtonCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	tagCols, err := tag.GetTags(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range tagCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	emailAddrCols, err := emailAddr.GetEmailAddrs(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range emailAddrCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	urlLinkCols, err := urlLink.GetUrlLinks(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range urlLinkCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	fileCols, err := file.GetFiles(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range fileCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	imageCols, err := image.GetImages(parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range imageCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	return tableColData, tableColsByID, nil
}
