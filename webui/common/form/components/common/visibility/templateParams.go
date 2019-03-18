// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package visibility

import (
	"resultra/tracker/webui/common/recordFilter"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type VisibilityPropertyTemplateParams struct {
	PanelParams                          propertiesSidebar.PanelTemplateParams
	ElemPrefix                           string
	VisibilityFilterConditionPanelParams recordFilter.FilterPanelTemplateParams
}

func NewComponentVisibilityTempalteParams(elemPrefix string, panelID string) VisibilityPropertyTemplateParams {

	panelParams := propertiesSidebar.PanelTemplateParams{
		PanelHeaderLabel: "Conditional Visibility",
		PanelID:          panelID}

	params := VisibilityPropertyTemplateParams{ElemPrefix: elemPrefix,
		PanelParams:                          panelParams,
		VisibilityFilterConditionPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix)}

	return params
}
