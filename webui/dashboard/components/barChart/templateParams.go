package barChart

import (
	"resultra/datasheet/webui/dashboard/components/common/newComponentDialog"
)

type BarChartTemplateParams struct {
	ElemPrefix        string
	SelectTableParams newComponentDialog.SelectTableTemplateParams
}

var TemplateParams BarChartTemplateParams

func init() {

	elemPrefix := "barChart_"

	TemplateParams = BarChartTemplateParams{
		ElemPrefix: elemPrefix,
		SelectTableParams: newComponentDialog.SelectTableTemplateParams{elemPrefix,
			"Select a table as the source of data for this bar chart."}}
}
