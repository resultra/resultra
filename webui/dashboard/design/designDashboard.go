package design

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/generic/api"
	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/admin/common/inputProperties"
	"resultra/datasheet/webui/common"
	dashboardCommon "resultra/datasheet/webui/dashboard/common"
	"resultra/datasheet/webui/dashboard/components"
	"resultra/datasheet/webui/dashboard/design/properties"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var designDashboardTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/dashboard/design/designDashboard.html",
		"static/dashboard/design/designDashboardPalette.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList,
		dashboardCommon.TemplateFileList,
		inputProperties.TemplateFileList,
		adminCommon.TemplateFileList,
		components.TemplateFileList,
		properties.TemplateFileList}
	designDashboardTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func DesignDashboard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	dashboardID := vars["dashboardID"]
	log.Println("Design Dashboard: editing for dashboard with params = %+v", vars)

	dashboardForDesign, getErr := dashboard.GetDashboard(dashboardID)
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	}

	templParams, templErr := createDashboardTemplateParams(r, dashboardForDesign)
	if templErr != nil {
		api.WriteErrorResponse(w, templErr)
		return
	}

	err := designDashboardTemplates.ExecuteTemplate(w, "designDashboard", *templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
