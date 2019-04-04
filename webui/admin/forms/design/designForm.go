// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package design

import (
	"github.com/gorilla/mux"
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/userAuth"
	"html/template"
	"log"
	"net/http"

	"github.com/resultra/resultra/server/databaseController"
	"github.com/resultra/resultra/server/generic/api"
	"github.com/resultra/resultra/server/workspace"
	adminCommon "github.com/resultra/resultra/webui/admin/common"
	"github.com/resultra/resultra/webui/admin/common/inputProperties"
	"github.com/resultra/resultra/webui/admin/forms/design/properties"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/common/form/components"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/thirdParty"
)

// Parse all the HTML templates at once. Individual templates can then
// be referenced throughout this package using htmlTemplates.ExecuteTemplate(...)
var designFormTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/admin/forms/design/designForm.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList,
		inputProperties.TemplateFileList,
		components.TemplateFileList,
		properties.TemplateFileList,
		adminCommon.TemplateFileList}
	designFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func designFormPageContent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	formID := vars["formID"]
	log.Println("Design Form: editing for form with ID = ", formID)

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	workspaceName, workspaceErr := workspace.GetWorkspaceName(trackerDBHandle)
	if workspaceErr != nil {
		http.Error(w, workspaceErr.Error(), http.StatusInternalServerError)
		return
	}

	formDBInfo, getErr := databaseController.GetFormDatabaseInfo(trackerDBHandle, formID)
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	}

	templParams := createDesignFormTemplateParams(r, formDBInfo, workspaceName)

	err := designFormTemplates.ExecuteTemplate(w, "designFormContent", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func designFormOffpageContent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	formID := vars["formID"]
	log.Println("Design Form: editing for form with ID = ", formID)

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	workspaceName, workspaceErr := workspace.GetWorkspaceName(trackerDBHandle)
	if workspaceErr != nil {
		http.Error(w, workspaceErr.Error(), http.StatusInternalServerError)
		return
	}

	formDBInfo, getErr := databaseController.GetFormDatabaseInfo(trackerDBHandle, formID)
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	}

	templParams := createDesignFormTemplateParams(r, formDBInfo, workspaceName)

	err := designFormTemplates.ExecuteTemplate(w, "designFormOffpageContent", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func designFormSidebarContent(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	formID := vars["formID"]
	log.Println("Design Form: editing for form with ID = ", formID)

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
	if dbErr != nil {
		http.Error(w, dbErr.Error(), http.StatusInternalServerError)
		return
	}

	workspaceName, workspaceErr := workspace.GetWorkspaceName(trackerDBHandle)
	if workspaceErr != nil {
		http.Error(w, workspaceErr.Error(), http.StatusInternalServerError)
		return
	}

	formDBInfo, getErr := databaseController.GetFormDatabaseInfo(trackerDBHandle, formID)
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	}

	templParams := createDesignFormTemplateParams(r, formDBInfo, workspaceName)

	err := designFormTemplates.ExecuteTemplate(w, "formDesignPropertiesSidebar", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
