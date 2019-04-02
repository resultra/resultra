// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package form

import (
	"database/sql"

	"github.com/resultra/resultra/server/form/components/attachment"
	"github.com/resultra/resultra/server/form/components/caption"
	"github.com/resultra/resultra/server/form/components/checkBox"
	"github.com/resultra/resultra/server/form/components/comment"
	"github.com/resultra/resultra/server/form/components/datePicker"
	"github.com/resultra/resultra/server/form/components/emailAddr"
	"github.com/resultra/resultra/server/form/components/file"
	"github.com/resultra/resultra/server/form/components/formButton"
	"github.com/resultra/resultra/server/form/components/gauge"
	"github.com/resultra/resultra/server/form/components/header"
	"github.com/resultra/resultra/server/form/components/htmlEditor"
	"github.com/resultra/resultra/server/form/components/image"
	"github.com/resultra/resultra/server/form/components/label"
	"github.com/resultra/resultra/server/form/components/numberInput"
	"github.com/resultra/resultra/server/form/components/progress"
	"github.com/resultra/resultra/server/form/components/rating"
	"github.com/resultra/resultra/server/form/components/selection"
	"github.com/resultra/resultra/server/form/components/socialButton"
	"github.com/resultra/resultra/server/form/components/textBox"
	"github.com/resultra/resultra/server/form/components/toggle"
	"github.com/resultra/resultra/server/form/components/urlLink"
	"github.com/resultra/resultra/server/form/components/userSelection"
	"github.com/resultra/resultra/server/form/components/userTag"
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
