package view

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/webui/common"
	dashboardCommon "resultra/datasheet/webui/dashboard/common"
	"resultra/datasheet/webui/dashboard/components"
	"resultra/datasheet/webui/dashboard/design/properties"
	"resultra/datasheet/webui/form/submit"
	"resultra/datasheet/webui/generic"
)

var viewDashboardTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/dashboard/view/viewDashboard.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList,
		dashboardCommon.TemplateFileList,
		components.TemplateFileList,
		properties.TemplateFileList,
		submit.TemplateFileList}
	viewDashboardTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ViewDashboardTemplateParams struct {
	DatabaseID      string
	DatabaseName    string
	DashboardID     string
	DashboardName   string
	Title           string
	ComponentParams components.ComponentViewTemplateParams
}

func ViewDashboard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	dashboardID := vars["dashboardID"]
	log.Println("ViewDashboard: viewing dashboard with params = %+v", vars)

	dashboardDbInfo, getErr := databaseController.GetDashboardDatabaseInfo(dashboardID)
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	}

	templParams := ViewDashboardTemplateParams{
		DatabaseID:      dashboardDbInfo.DatabaseID,
		DatabaseName:    dashboardDbInfo.DatabaseName,
		DashboardID:     dashboardDbInfo.DashboardID,
		DashboardName:   dashboardDbInfo.DashboardName,
		Title:           "View Dashboard",
		ComponentParams: components.ViewTemplateParams}

	err := viewDashboardTemplates.ExecuteTemplate(w, "viewDashboard", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
