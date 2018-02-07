package design

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/workspace"
	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/admin/common/inputProperties"
	"resultra/datasheet/webui/common"
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

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		api.WriteErrorResponse(w, dbErr)
		return
	}

	workspaceName, workspaceErr := workspace.GetWorkspaceName(trackerDBHandle)
	if workspaceErr != nil {
		http.Error(w, workspaceErr.Error(), http.StatusInternalServerError)
		return
	}

	dashboardForDesign, getErr := dashboard.GetDashboard(trackerDBHandle, dashboardID)
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	}

	templParams, templErr := createDashboardTemplateParams(r, dashboardForDesign, workspaceName)
	if templErr != nil {
		api.WriteErrorResponse(w, templErr)
		return
	}

	err := designDashboardTemplates.ExecuteTemplate(w, "designDashboard", *templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
