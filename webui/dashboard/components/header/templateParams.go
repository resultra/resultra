// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package header

import (
	"github.com/resultra/resultra/webui/dashboard/components/common/delete"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type HeaderDesignTemplateParams struct {
	ElemPrefix        string
	TitlePanelParams  propertiesSidebar.PanelTemplateParams
	FormatPanelParams propertiesSidebar.PanelTemplateParams
	DeletePanelParams delete.DeletePropertyPanelTemplateParams
}

var DesignTemplateParams HeaderDesignTemplateParams

func init() {

	elemPrefix := "header_"

	DesignTemplateParams = HeaderDesignTemplateParams{
		ElemPrefix:        elemPrefix,
		TitlePanelParams:  propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Title", PanelID: "headerTitle"},
		DeletePanelParams: delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "headerDelete", "Delete Header"),
		FormatPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "headerFormat"}}

}
