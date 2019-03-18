// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package header

import (
	"resultra/tracker/webui/common/form/components/common/delete"
	"resultra/tracker/webui/common/form/components/common/visibility"
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type HeaderTemplateParams struct {
	ElemPrefix            string
	FormatPanelParams     propertiesSidebar.PanelTemplateParams
	LabelPanelParams      propertiesSidebar.PanelTemplateParams
	VisibilityPanelParams visibility.VisibilityPropertyTemplateParams
	DeletePanelParams     delete.DeletePropertyPanelTemplateParams
}

var TemplateParams HeaderTemplateParams

func init() {

	elemPrefix := "header_"

	TemplateParams = HeaderTemplateParams{
		ElemPrefix:            elemPrefix,
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "headerFormat"},
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "headerDelete", "Delete Header"),
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "headerVisibility"),
		LabelPanelParams:      propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Header Text", PanelID: "headerLabel"}}

}
