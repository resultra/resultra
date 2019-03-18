// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package design

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/dashboard"
	"resultra/tracker/server/generic/api"
	"resultra/tracker/server/workspace"
	adminCommon "resultra/tracker/webui/admin/common"
	"resultra/tracker/webui/admin/common/inputProperties"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/dashboard/components"
	"resultra/tracker/webui/dashboard/design/properties"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
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

func designDashboardMainContent(w http.ResponseWriter, r *http.Request) {

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

	err := designDashboardTemplates.ExecuteTemplate(w, "designDashboardMainContent", *templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func designDashboardOffpageContent(w http.ResponseWriter, r *http.Request) {

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

	err := designDashboardTemplates.ExecuteTemplate(w, "designDashboardOffpageContent", *templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func designDashboardSidebarContent(w http.ResponseWriter, r *http.Request) {

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

	err := designDashboardTemplates.ExecuteTemplate(w, "dashboardDesignPropertiesSidebar", *templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
