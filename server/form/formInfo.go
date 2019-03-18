package form

import (
	"database/sql"

	"resultra/tracker/server/form/components/attachment"
	"resultra/tracker/server/form/components/caption"
	"resultra/tracker/server/form/components/checkBox"
	"resultra/tracker/server/form/components/comment"
	"resultra/tracker/server/form/components/datePicker"
	"resultra/tracker/server/form/components/emailAddr"
	"resultra/tracker/server/form/components/file"
	"resultra/tracker/server/form/components/formButton"
	"resultra/tracker/server/form/components/gauge"
	"resultra/tracker/server/form/components/header"
	"resultra/tracker/server/form/components/htmlEditor"
	"resultra/tracker/server/form/components/image"
	"resultra/tracker/server/form/components/label"
	"resultra/tracker/server/form/components/numberInput"
	"resultra/tracker/server/form/components/progress"
	"resultra/tracker/server/form/components/rating"
	"resultra/tracker/server/form/components/selection"
	"resultra/tracker/server/form/components/socialButton"
	"resultra/tracker/server/form/components/textBox"
	"resultra/tracker/server/form/components/toggle"
	"resultra/tracker/server/form/components/urlLink"
	"resultra/tracker/server/form/components/userSelection"
	"resultra/tracker/server/form/components/userTag"
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
	UserTags           []userTag.UserTag             `json:"userTags"`
}

type GetFormInfoParams struct {
	FormID string `json:"formID"`
}

func GetFormInfo(trackerDBHandle *sql.DB, formID string) (*FormInfo, error) {

	form, err := GetForm(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	textBoxes, err := textBox.GetTextBoxes(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	numberInputs, err := numberInput.GetNumberInputs(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	checkBoxes, err := checkBox.GetCheckBoxes(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	toggles, err := toggle.GetToggles(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	datePickers, err := datePicker.GetDatePickers(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	htmlEditors, err := htmlEditor.GetHtmlEditors(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	attachments, err := attachment.GetImages(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	headers, err := header.GetHeaders(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	formButtons, err := formButton.GetButtons(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	ratings, err := rating.GetRatings(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	comments, err := comment.GetComments(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	selections, err := selection.GetSelections(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	userSelections, err := userSelection.GetUserSelections(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	userTags, err := userTag.GetUserTags(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	progressIndicators, err := progress.GetProgressIndicators(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	captions, err := caption.GetCaptions(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	gauges, err := gauge.GetGauges(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	socialButtons, err := socialButton.GetSocialButtons(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	labels, err := label.GetLabels(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	emailAddrs, err := emailAddr.GetEmailAddrs(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	urlLinks, err := urlLink.GetUrlLinks(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	files, err := file.GetFiles(trackerDBHandle, formID)
	if err != nil {
		return nil, err
	}

	images, err := image.GetImages(trackerDBHandle, formID)
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
		UserTags:           userTags,
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
