package displayTable

import (
	"database/sql"
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
	"resultra/datasheet/server/displayTable/columns/textSelection"
	"resultra/datasheet/server/displayTable/columns/toggle"
	"resultra/datasheet/server/displayTable/columns/urlLink"
	"resultra/datasheet/server/displayTable/columns/userSelection"
	"resultra/datasheet/server/displayTable/columns/userTag"
)

type TableColsInfo []interface{}
type TableColsByID map[string]interface{}

func getTableCols(trackerDBHandle *sql.DB, parentTableID string) (TableColsInfo, TableColsByID, error) {

	tableColData := TableColsInfo{}
	tableColsByID := TableColsByID{}

	numberInputCols, err := numberInput.GetNumberInputs(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range numberInputCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	progressCols, err := progress.GetProgressIndicators(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range progressCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	textInputCols, err := textInput.GetTextInputs(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range textInputCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	textSelectionCols, err := textSelection.GetTextSelections(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range textSelectionCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	datePickerCols, err := datePicker.GetDatePickers(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range datePickerCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	checkBoxCols, err := checkBox.GetCheckBoxes(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range checkBoxCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	ratingCols, err := rating.GetRatings(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range ratingCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	toggleCols, err := toggle.GetToggles(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range toggleCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	userSelectionCols, err := userSelection.GetUserSelections(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range userSelectionCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	userTagCols, err := userTag.GetUserTags(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range userTagCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	noteCols, err := note.GetNotes(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range noteCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	commentCols, err := comment.GetComments(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range commentCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	attachCols, err := attachment.GetAttachments(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range attachCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	buttonCols, err := formButton.GetButtons(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range buttonCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	socialButtonCols, err := socialButton.GetSocialButtons(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range socialButtonCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	tagCols, err := tag.GetTags(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range tagCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	emailAddrCols, err := emailAddr.GetEmailAddrs(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range emailAddrCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	urlLinkCols, err := urlLink.GetUrlLinks(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range urlLinkCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	fileCols, err := file.GetFiles(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range fileCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	imageCols, err := image.GetImages(trackerDBHandle, parentTableID)
	if err != nil {
		return nil, nil, fmt.Errorf("getTableCols: %v", err)
	}
	for _, col := range imageCols {
		tableColData = append(tableColData, col)
		tableColsByID[col.ColumnID] = col
	}

	return tableColData, tableColsByID, nil
}
