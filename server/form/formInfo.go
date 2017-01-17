package form

import (
	"resultra/datasheet/server/form/components/checkBox"
	"resultra/datasheet/server/form/components/comment"
	"resultra/datasheet/server/form/components/datePicker"
	"resultra/datasheet/server/form/components/formButton"
	"resultra/datasheet/server/form/components/header"
	"resultra/datasheet/server/form/components/htmlEditor"
	"resultra/datasheet/server/form/components/image"
	"resultra/datasheet/server/form/components/rating"
	"resultra/datasheet/server/form/components/selection"
	"resultra/datasheet/server/form/components/textBox"
	"resultra/datasheet/server/form/components/userSelection"
)

type FormInfo struct {
	Form           Form                          `json:"form"`
	TextBoxes      []textBox.TextBox             `json:"textBoxes"`
	CheckBoxes     []checkBox.CheckBox           `json:"checkBoxes"`
	DatePickers    []datePicker.DatePicker       `json:"datePickers"`
	HtmlEditors    []htmlEditor.HtmlEditor       `json:"htmlEditors"`
	Ratings        []rating.Rating               `json:"ratings"`
	Comments       []comment.Comment             `json:"comments"`
	Images         []image.Image                 `json:"images"`
	Headers        []header.Header               `json:"headers"`
	FormButtons    []formButton.FormButton       `json:"formButtons"`
	Selections     []selection.Selection         `json:"selections"`
	UserSelections []userSelection.UserSelection `json:"userSelections"`
}

type GetFormInfoParams struct {
	FormID string `json:"formID"`
}

func getFormInfo(params GetFormInfoParams) (*FormInfo, error) {

	form, err := GetForm(params.FormID)
	if err != nil {
		return nil, err
	}

	textBoxes, err := textBox.GetTextBoxes(params.FormID)
	if err != nil {
		return nil, err
	}

	checkBoxes, err := checkBox.GetCheckBoxes(params.FormID)
	if err != nil {
		return nil, err
	}

	datePickers, err := datePicker.GetDatePickers(params.FormID)
	if err != nil {
		return nil, err
	}

	htmlEditors, err := htmlEditor.GetHtmlEditors(params.FormID)
	if err != nil {
		return nil, err
	}

	images, err := image.GetImages(params.FormID)
	if err != nil {
		return nil, err
	}

	headers, err := header.GetHeaders(params.FormID)
	if err != nil {
		return nil, err
	}

	formButtons, err := formButton.GetButtons(params.FormID)
	if err != nil {
		return nil, err
	}

	ratings, err := rating.GetRatings(params.FormID)
	if err != nil {
		return nil, err
	}

	comments, err := comment.GetComments(params.FormID)
	if err != nil {
		return nil, err
	}

	selections, err := selection.GetSelections(params.FormID)
	if err != nil {
		return nil, err
	}

	userSelections, err := userSelection.GetUserSelections(params.FormID)
	if err != nil {
		return nil, err
	}

	formInfo := FormInfo{
		Form:           *form,
		TextBoxes:      textBoxes,
		CheckBoxes:     checkBoxes,
		DatePickers:    datePickers,
		HtmlEditors:    htmlEditors,
		Images:         images,
		Headers:        headers,
		FormButtons:    formButtons,
		Ratings:        ratings,
		Comments:       comments,
		Selections:     selections,
		UserSelections: userSelections}

	return &formInfo, nil
}
