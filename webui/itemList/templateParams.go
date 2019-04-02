// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package itemList

import (
	"github.com/resultra/resultra/webui/common/form/components"
	"github.com/resultra/resultra/webui/common/recordFilter"
	"github.com/resultra/resultra/webui/generic/propertiesSidebar"
)

type ViewListTemplateParams struct {
	DisplayPanelParams      propertiesSidebar.PanelTemplateParams
	FilteringPanelParams    propertiesSidebar.PanelTemplateParams
	SortPanelParams         propertiesSidebar.PanelTemplateParams
	ComponentParams         components.ComponentViewTemplateParams
	FilterConfigPanelParams recordFilter.FilterPanelTemplateParams
}

var ViewListTemplParams ViewListTemplateParams

func init() {

	elemPrefix := "form_"

	ViewListTemplParams = ViewListTemplateParams{
		DisplayPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "View",
			PanelID: "viewListDisplay"},
		FilteringPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering",
			PanelID: "viewFormFiltering"},
		FilterConfigPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix),
		SortPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Sorting",
			PanelID: "viewFormSorting"},
		ComponentParams: components.ViewTemplateParams}

}
