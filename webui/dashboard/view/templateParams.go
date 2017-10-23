package view

import (
	"resultra/datasheet/webui/dashboard/components"
)

type ViewDashboardTemplateParams struct {
	ComponentParams components.ComponentViewTemplateParams
}

var ViewTemplateParams ViewDashboardTemplateParams

func init() {
	ViewTemplateParams = ViewDashboardTemplateParams{
		ComponentParams: components.ViewTemplateParams}

}
