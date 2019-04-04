// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package design

import (
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/dashboard"
	"github.com/resultra/resultra/server/generic/api"
	"github.com/resultra/resultra/server/workspace"
	adminCommon "github.com/resultra/resultra/webui/admin/common"
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/dashboard/components"
	"github.com/resultra/resultra/webui/dashboard/design/properties"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/thirdParty"
	"html/template"
	"log"
	"net/http"
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
