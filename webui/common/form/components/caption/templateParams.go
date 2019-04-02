// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package caption

import (
	"github.com/resultra/resultra/webui/common/form/components/common/delete"
	"github.com/resultra/resultra/webui/common/form/components/common/visibility"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type CaptionTemplateParams struct {
	ElemPrefix            string
	FormatPanelParams     propertiesSidebar.PanelTemplateParams
	CaptionPanelParams    propertiesSidebar.PanelTemplateParams
	VisibilityPanelParams visibility.VisibilityPropertyTemplateParams
	DeletePanelParams     delete.DeletePropertyPanelTemplateParams
}

var TemplateParams CaptionTemplateParams

func init() {

	elemPrefix := "caption_"

	TemplateParams = CaptionTemplateParams{
		ElemPrefix:            elemPrefix,
		VisibilityPanelParams: visibility.NewComponentVisibilityTempalteParams(elemPrefix, "captionVisibility"),
		DeletePanelParams:     delete.NewDeletePropertyPanelTemplateParams(elemPrefix, "captionDelete", "Delete Caption"),
		FormatPanelParams:     propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Format", PanelID: "captionFormat"},
		CaptionPanelParams:    propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Caption", PanelID: "headerCaption"}}
}
