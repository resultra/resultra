// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package delete

import (
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type DeletePropertyPanelTemplateParams struct {
	PanelParams       propertiesSidebar.PanelTemplateParams
	ElemPrefix        string
	DeleteButtonLabel string
}

func NewDeletePropertyPanelTemplateParams(elemPrefix string, panelID string, deleteButtonLabel string) DeletePropertyPanelTemplateParams {

	panelParams := propertiesSidebar.PanelTemplateParams{
		PanelHeaderLabel: "Delete",
		PanelID:          panelID}

	params := DeletePropertyPanelTemplateParams{ElemPrefix: elemPrefix,
		PanelParams:       panelParams,
		DeleteButtonLabel: deleteButtonLabel}

	return params
}
