// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package components

import (
	"resultra/tracker/webui/common/form/components/attachment"
	"resultra/tracker/webui/common/form/components/checkBox"
	"resultra/tracker/webui/common/form/components/comment"
	"resultra/tracker/webui/common/form/components/datePicker"
	"resultra/tracker/webui/common/form/components/emailAddr"
	"resultra/tracker/webui/common/form/components/file"
	"resultra/tracker/webui/common/form/components/gauge"
	"resultra/tracker/webui/common/form/components/htmlEditor"
	"resultra/tracker/webui/common/form/components/image"
	"resultra/tracker/webui/common/form/components/label"
	"resultra/tracker/webui/common/form/components/numberInput"
	"resultra/tracker/webui/common/form/components/progress"
	"resultra/tracker/webui/common/form/components/rating"
	"resultra/tracker/webui/common/form/components/selection"
	"resultra/tracker/webui/common/form/components/socialButton"
	"resultra/tracker/webui/common/form/components/textBox"
	"resultra/tracker/webui/common/form/components/toggle"
	"resultra/tracker/webui/common/form/components/urlLink"
	"resultra/tracker/webui/common/form/components/userSelection"
	"resultra/tracker/webui/common/form/components/userTag"
)

type ComponentViewTemplateParams struct {
	CheckBoxParams      checkBox.CheckboxViewTemplateParams
	ToggleParams        toggle.ToggleViewTemplateParams
	DatePickerParams    datePicker.DatePickerViewTemplateParams
	TextBoxParams       textBox.TextboxViewTemplateParams
	AttachmentParams    attachment.ImageViewTemplateParams
	HTMLEditorParams    htmlEditor.HTMLEditorViewTemplateParams
	RatingParams        rating.RatingViewTemplateParams
	CommentParams       comment.CommentViewTemplateParams
	SelectionParams     selection.SelectionViewTemplateParams
	UserSelectionParams userSelection.UserSelectionViewTemplateParams
	UserTagParams       userTag.UserTagViewTemplateParams
	ProgressParams      progress.ProgressViewTemplateParams
	GaugeParams         gauge.GaugeViewTemplateParams
	NumberInputParams   numberInput.NumberInputViewTemplateParams
	SocialButtonParams  socialButton.SocialButtonViewTemplateParams
	LabelParams         label.LabelViewTemplateParams
	EmailAddrParams     emailAddr.EmailAddrViewTemplateParams
	UrlLinkParams       urlLink.UrlLinkViewTemplateParams
	FileParams          file.FileViewTemplateParams
	ImageParams         image.ImageViewTemplateParams
}

//var DesignTemplateParams ComponentDesignTemplateParams
var ViewTemplateParams ComponentViewTemplateParams

func init() {

	ViewTemplateParams = ComponentViewTemplateParams{
		CheckBoxParams:      checkBox.ViewTemplateParams,
		ToggleParams:        toggle.ViewTemplateParams,
		DatePickerParams:    datePicker.ViewTemplateParams,
		TextBoxParams:       textBox.ViewTemplateParams,
		AttachmentParams:    attachment.ViewTemplateParams,
		HTMLEditorParams:    htmlEditor.ViewTemplateParams,
		RatingParams:        rating.ViewTemplateParams,
		CommentParams:       comment.ViewTemplateParams,
		SelectionParams:     selection.ViewTemplateParams,
		UserSelectionParams: userSelection.ViewTemplateParams,
		UserTagParams:       userTag.ViewTemplateParams,
		ProgressParams:      progress.ViewTemplateParams,
		GaugeParams:         gauge.ViewTemplateParams,
		NumberInputParams:   numberInput.ViewTemplateParams,
		SocialButtonParams:  socialButton.ViewTemplateParams,
		LabelParams:         label.ViewTemplateParams,
		EmailAddrParams:     emailAddr.ViewTemplateParams,
		UrlLinkParams:       urlLink.ViewTemplateParams,
		FileParams:          file.ViewTemplateParams,
		ImageParams:         image.ViewTemplateParams}

}
