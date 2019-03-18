// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package permissions

import (
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type PermissionsPropertyTemplateParams struct {
	PanelParams propertiesSidebar.PanelTemplateParams
	ElemPrefix  string
}

func NewPermissionTemplateParams(elemPrefix string, panelID string) PermissionsPropertyTemplateParams {

	panelParams := propertiesSidebar.PanelTemplateParams{
		PanelHeaderLabel: "Permissions",
		PanelID:          panelID}

	params := PermissionsPropertyTemplateParams{ElemPrefix: elemPrefix,
		PanelParams: panelParams}

	return params
}
