package webui

import "html/template"

// Parse all the HTML templates at once. Individual templates can then
// be referenced throughout this package using htmlTemplates.ExecuteTemplate(...)
var htmlTemplates = template.Must(template.ParseFiles(

	"dashboard/barChart/barChartProps.html",
	"dashboard/barChart/newBarChartDialog.html",
	"dashboard/dashboardCommon.html",
	"dashboard/dashboardProps.html",
	"dashboard/designDashboard.html",

	"template/calcField.html",
	"template/common.html",
	"template/filterRecords.html",
	"template/home.html",
	"form/designForm.html",
	"form/checkBox/newCheckBoxDialog.html",
	"form/textBox/newTextBoxDialog.html",
	"template/tableProps.html",
	"form/viewForm.html",
	"form/common/newFormElemDialog.html",
	"form/checkBox/checkboxProp.html"))
