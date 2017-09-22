package form

import (
	"resultra/datasheet/server/form/components/attachment"
	"resultra/datasheet/server/form/components/caption"
	"resultra/datasheet/server/form/components/checkBox"
	"resultra/datasheet/server/form/components/comment"
	"resultra/datasheet/server/form/components/datePicker"
	"resultra/datasheet/server/form/components/emailAddr"
	"resultra/datasheet/server/form/components/file"
	"resultra/datasheet/server/form/components/formButton"
	"resultra/datasheet/server/form/components/gauge"
	"resultra/datasheet/server/form/components/header"
	"resultra/datasheet/server/form/components/htmlEditor"
	"resultra/datasheet/server/form/components/image"
	"resultra/datasheet/server/form/components/label"
	"resultra/datasheet/server/form/components/numberInput"
	"resultra/datasheet/server/form/components/progress"
	"resultra/datasheet/server/form/components/rating"
	"resultra/datasheet/server/form/components/selection"
	"resultra/datasheet/server/form/components/socialButton"
	"resultra/datasheet/server/form/components/textBox"
	"resultra/datasheet/server/form/components/toggle"
	"resultra/datasheet/server/form/components/urlLink"
	"resultra/datasheet/server/form/components/userSelection"
)

type FormInfo struct {
	Form               Form                          `json:"form"`
	TextBoxes          []textBox.TextBox             `json:"textBoxes"`
	NumberInputs       []numberInput.NumberInput     `json:"numberInputs"`
	CheckBoxes         []checkBox.CheckBox           `json:"checkBoxes"`
	Toggles            []toggle.Toggle               `json:"toggles"`
	DatePickers        []datePicker.DatePicker       `json:"datePickers"`
	HtmlEditors        []htmlEditor.HtmlEditor       `json:"htmlEditors"`
	Ratings            []rating.Rating               `json:"ratings"`
	Comments           []comment.Comment             `json:"comments"`
	Attachments        []attachment.Image            `json:"attachments"`
	Headers            []header.Header               `json:"headers"`
	FormButtons        []formButton.FormButton       `json:"formButtons"`
	Selections         []selection.Selection         `json:"selections"`
	UserSelections     []userSelection.UserSelection `json:"userSelections"`
	ProgressIndicators []progress.Progress           `json:"progressIndicators"`
	Gauges             []gauge.Gauge                 `json:"gauges"`
	Captions           []caption.Caption             `json:"captions"`
	SocialButtons      []socialButton.SocialButton   `json:"socialButtons"`
	Labels             []label.Label                 `json:"labels"`
	EmailAddrs         []emailAddr.EmailAddr         `json:"emailAddrs"`
	UrlLinks           []urlLink.UrlLink             `json:"urlLinks"`
	Files              []file.File                   `json:"files"`
	Images             []image.Image                 `json:"images"`
}

type GetFormInfoParams struct {
	FormID string `json:"formID"`
}

func GetFormInfo(formID string) (*FormInfo, error) {

	form, err := GetForm(formID)
	if err != nil {
		return nil, err
	}

	textBoxes, err := textBox.GetTextBoxes(formID)
	if err != nil {
		return nil, err
	}

	numberInputs, err := numberInput.GetNumberInputs(formID)
	if err != nil {
		return nil, err
	}

	checkBoxes, err := checkBox.GetCheckBoxes(formID)
	if err != nil {
		return nil, err
	}

	toggles, err := toggle.GetToggles(formID)
	if err != nil {
		return nil, err
	}

	datePickers, err := datePicker.GetDatePickers(formID)
	if err != nil {
		return nil, err
	}

	htmlEditors, err := htmlEditor.GetHtmlEditors(formID)
	if err != nil {
		return nil, err
	}

	attachments, err := attachment.GetImages(formID)
	if err != nil {
		return nil, err
	}

	headers, err := header.GetHeaders(formID)
	if err != nil {
		return nil, err
	}

	formButtons, err := formButton.GetButtons(formID)
	if err != nil {
		return nil, err
	}

	ratings, err := rating.GetRatings(formID)
	if err != nil {
		return nil, err
	}

	comments, err := comment.GetComments(formID)
	if err != nil {
		return nil, err
	}

	selections, err := selection.GetSelections(formID)
	if err != nil {
		return nil, err
	}

	userSelections, err := userSelection.GetUserSelections(formID)
	if err != nil {
		return nil, err
	}

	progressIndicators, err := progress.GetProgressIndicators(formID)
	if err != nil {
		return nil, err
	}

	captions, err := caption.GetCaptions(formID)
	if err != nil {
		return nil, err
	}

	gauges, err := gauge.GetGauges(formID)
	if err != nil {
		return nil, err
	}

	socialButtons, err := socialButton.GetSocialButtons(formID)
	if err != nil {
		return nil, err
	}

	labels, err := label.GetLabels(formID)
	if err != nil {
		return nil, err
	}

	emailAddrs, err := emailAddr.GetEmailAddrs(formID)
	if err != nil {
		return nil, err
	}

	urlLinks, err := urlLink.GetUrlLinks(formID)
	if err != nil {
		return nil, err
	}

	files, err := file.GetFiles(formID)
	if err != nil {
		return nil, err
	}

	images, err := image.GetImages(formID)
	if err != nil {
		return nil, err
	}

	formInfo := FormInfo{
		Form:               *form,
		TextBoxes:          textBoxes,
		NumberInputs:       numberInputs,
		CheckBoxes:         checkBoxes,
		Toggles:            toggles,
		DatePickers:        datePickers,
		HtmlEditors:        htmlEditors,
		Attachments:        attachments,
		Headers:            headers,
		FormButtons:        formButtons,
		Ratings:            ratings,
		Comments:           comments,
		Selections:         selections,
		UserSelections:     userSelections,
		ProgressIndicators: progressIndicators,
		Captions:           captions,
		Gauges:             gauges,
		SocialButtons:      socialButtons,
		Labels:             labels,
		EmailAddrs:         emailAddrs,
		UrlLinks:           urlLinks,
		Files:              files,
		Images:             images}

	return &formInfo, nil
}
