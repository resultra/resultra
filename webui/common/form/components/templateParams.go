package components

import (
	"resultra/datasheet/webui/common/form/components/checkBox"
	"resultra/datasheet/webui/common/form/components/comment"
	"resultra/datasheet/webui/common/form/components/datePicker"
	"resultra/datasheet/webui/common/form/components/gauge"
	"resultra/datasheet/webui/common/form/components/htmlEditor"
	"resultra/datasheet/webui/common/form/components/image"
	"resultra/datasheet/webui/common/form/components/progress"
	"resultra/datasheet/webui/common/form/components/rating"
	"resultra/datasheet/webui/common/form/components/selection"
	"resultra/datasheet/webui/common/form/components/textBox"
	"resultra/datasheet/webui/common/form/components/userSelection"
)

type ComponentViewTemplateParams struct {
	CheckBoxParams      checkBox.CheckboxViewTemplateParams
	DatePickerParams    datePicker.DatePickerViewTemplateParams
	TextBoxParams       textBox.TextboxViewTemplateParams
	ImageParams         image.ImageViewTemplateParams
	HTMLEditorParams    htmlEditor.HTMLEditorViewTemplateParams
	RatingParams        rating.RatingViewTemplateParams
	CommentParams       comment.CommentViewTemplateParams
	SelectionParams     selection.SelectionViewTemplateParams
	UserSelectionParams userSelection.UserSelectionViewTemplateParams
	ProgressParams      progress.ProgressViewTemplateParams
	GaugeParams         gauge.GaugeViewTemplateParams
}

//var DesignTemplateParams ComponentDesignTemplateParams
var ViewTemplateParams ComponentViewTemplateParams

func init() {

	ViewTemplateParams = ComponentViewTemplateParams{
		CheckBoxParams:      checkBox.ViewTemplateParams,
		DatePickerParams:    datePicker.ViewTemplateParams,
		TextBoxParams:       textBox.ViewTemplateParams,
		ImageParams:         image.ViewTemplateParams,
		HTMLEditorParams:    htmlEditor.ViewTemplateParams,
		RatingParams:        rating.ViewTemplateParams,
		CommentParams:       comment.ViewTemplateParams,
		SelectionParams:     selection.ViewTemplateParams,
		UserSelectionParams: userSelection.ViewTemplateParams,
		ProgressParams:      progress.ViewTemplateParams,
		GaugeParams:         gauge.ViewTemplateParams}

}
