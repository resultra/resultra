// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package fieldList

import (
	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/runtimeConfig"
	"github.com/resultra/resultra/server/databaseController"
	"github.com/resultra/resultra/server/userRole"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/workspace"
	adminCommon "github.com/resultra/resultra/webui/admin/common"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/thirdParty"
)

var fieldListTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{
		"static/admin/fields/fieldList/fieldListAdminPage.html",
		"static/admin/fields/fieldList/fieldList.html",
		"static/admin/fields/fieldList/newFieldDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	fieldListTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FieldTemplParams struct {
	Title                 string
	DatabaseID            string
	DatabaseName          string
	WorkspaceName         string
	FieldID               string
	FieldName             string
	CurrUserIsAdmin       bool
	IsSingleUserWorkspace bool
}

func fieldListAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

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

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(trackerDBHandle, databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	templParams := FieldTemplParams{
		Title:                 "Field Settings",
		DatabaseID:            databaseID,
		DatabaseName:          dbInfo.DatabaseName,
		WorkspaceName:         workspaceName,
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace(),
		CurrUserIsAdmin:       isAdmin}

	if err := fieldListTemplates.ExecuteTemplate(w, "fieldAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func fieldListOffpageContent(w http.ResponseWriter, r *http.Request) {

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Error(w, authErr.Error(), http.StatusInternalServerError)
		return
	}

	if err := fieldListTemplates.ExecuteTemplate(w, "fieldAdminPageOffpageContent", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
