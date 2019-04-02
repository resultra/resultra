// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package progress

import (
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common/form/components/common/delete"
	"github.com/resultra/resultra/webui/common/form/components/common/label"
	"github.com/resultra/resultra/webui/common/form/components/common/newFormElemDialog"
	"github.com/resultra/resultra/webui/common/form/components/common/visibility"
	"github.com/resultra/resultra/webui/common/valueThreshold"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type ProgressDesignTemplateParams struct {
	ElemPrefix               string
	FormatPanelParams        propertiesSidebar.PanelTemplateParams
	RangePanelParams         propertiesSidebar.PanelTemplateParams
	ThresholdPanelParams     propertiesSidebar.PanelTemplateParams
	ThresholdValueParams     valueThreshold.ThresholdValuesPanelTemplateParams
	NewComponentDialogParams newFormElemDialog.TemplateParams
	LabelPanelParams         label.LabelPropertyTemplateParams
	VisibilityPanelParams    visibility.VisibilityPropertyTemplateParams
	DeletePanelParams        delete.DeletePropertyPanelTemplateParams
	HelpPopupParams          inputProperties.HelpPopupPropertyTemplateParams
}

type ProgressViewTemplateParams struct {
	ElemPrefix string
}

var DesignTemplateParams ProgressDesignTemplateParams
var ViewTemplateParams ProgressViewTemplateParams

func init() {

	elemPrefix := "progress_"

	DesignTemplateParams = ProgressDesignTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "progressLabel"}},
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "progressDelete", "Delete Progress Indicator"),
		HelpPopupParams:       inputProperties.NewHelpPopupTemplateParams(elemPrefix, "progressHelp"),
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "progressVisibility"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "progressFormat"},
		RangePanelParams:      propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Range", PanelID: "progressRange"},
		ThresholdPanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Value Thresholds", PanelID: "progressThreshold"},
		ThresholdValueParams:  valueThreshold.NewThresholdValuesTemplateParams(elemPrefix),
		NewComponentDialogParams: newFormElemDialog.TemplateParams{
			ElemPrefix:         elemPrefix,
			DialogTitle:        "New Progress Indicator",
			FieldInfoPrompt:    `Progress indicators use a field value to determine the progress level.`,
			NewFieldInfoPrompt: `Enter the parameters for the new field to store this progress indicator's values.`}}

	ViewTemplateParams = ProgressViewTemplateParams{
		ElemPrefix: elemPrefix}

}
