package view

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/userRole"
	"resultra/datasheet/webui/common"
	dashboardCommon "resultra/datasheet/webui/dashboard/common"
	"resultra/datasheet/webui/dashboard/components"
	"resultra/datasheet/webui/dashboard/design/properties"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var viewDashboardTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/dashboard/view/viewDashboard.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList,
		dashboardCommon.TemplateFileList,
		components.TemplateFileList,
		properties.TemplateFileList}
	viewDashboardTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ViewDashboardTemplateParams struct {
	DatabaseID      string
	DatabaseName    string
	DashboardID     string
	DashboardName   string
	CurrUserIsAdmin bool
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

	hasViewPrivs, privsErr := userRole.CurrentUserHasDashboardViewPrivs(r, dashboardDbInfo.DatabaseID, dashboardID)
	if privsErr != nil {
		api.WriteErrorResponse(w, privsErr)
		return
	}
	if !hasViewPrivs {
		api.WriteErrorResponse(w, fmt.Errorf("ERROR: No permissions to view this dashboard"))
		return
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dashboardDbInfo.DatabaseID)

	templParams := ViewDashboardTemplateParams{
		DatabaseID:      dashboardDbInfo.DatabaseID,
		DatabaseName:    dashboardDbInfo.DatabaseName,
		DashboardID:     dashboardDbInfo.DashboardID,
		DashboardName:   dashboardDbInfo.DashboardName,
		CurrUserIsAdmin: isAdmin,
		Title:           "View Dashboard",
		ComponentParams: components.ViewTemplateParams}

	err := viewDashboardTemplates.ExecuteTemplate(w, "viewDashboard", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
