// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package colProps

import (
	"github.com/resultra/resultra/webui/common/form/components/common/label"
	"github.com/resultra/resultra/webui/common/form/components/common/permissions"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type SocialButtonColPropsTemplateParams struct {
	ElemPrefix            string
	LabelPanelParams      label.LabelPropertyTemplateParams
	PermissionPanelParams permissions.PermissionsPropertyTemplateParams
}

func newSocialButtonTemplateParams() SocialButtonColPropsTemplateParams {

	elemPrefix := "socialButton_"

	templParams := SocialButtonColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "socialButtonLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "socialButtonPerms")}

	return templParams

}
