package view

import (
	"resultra/tracker/webui/dashboard/components"
)

type ViewDashboardTemplateParams struct {
	ComponentParams components.ComponentViewTemplateParams
}

var ViewTemplateParams ViewDashboardTemplateParams

func init() {
	ViewTemplateParams = ViewDashboardTemplateParams{
		ComponentParams: components.ViewTemplateParams}

}
