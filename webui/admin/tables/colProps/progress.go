// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package colProps

import (
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common/form/components/common/label"
	"github.com/resultra/resultra/webui/common/valueThreshold"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type ProgressColPropsTemplateParams struct {
	ElemPrefix           string
	ThresholdValueParams valueThreshold.ThresholdValuesPanelTemplateParams
	LabelPanelParams     label.LabelPropertyTemplateParams
	HelpPopupParams      inputProperties.HelpPopupPropertyTemplateParams
}

func newProgressTemplateParams() ProgressColPropsTemplateParams {

	elemPrefix := "progress_"

	templParams := ProgressColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "progressLabel"}},
		HelpPopupParams:      inputProperties.NewHelpPopupTemplateParams(elemPrefix, "progressHelp"),
		ThresholdValueParams: valueThreshold.NewThresholdValuesTemplateParams(elemPrefix)}

	return templParams

}
