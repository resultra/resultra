package barChart

import (
	"resultra/datasheet/webui/dashboard/components/common/newComponentDialog"
)

type BarChartTemplateParams struct {
	SelectTableParams newComponentDialog.SelectTableTemplateParams
}

var TemplateParams BarChartTemplateParams

func init() {
	TemplateParams = BarChartTemplateParams{
		SelectTableParams: newComponentDialog.SelectTableTemplateParams{"barChart_",
			"Select a table as the source of data for this bar chart."}}
}
