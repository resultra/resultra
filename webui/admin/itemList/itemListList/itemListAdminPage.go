// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package itemListList

import (
	"html/template"
	"net/http"
	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/databaseController"
	"resultra/tracker/server/userRole"
	"resultra/tracker/server/workspace"
	adminCommon "resultra/tracker/webui/admin/common"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"

	"github.com/gorilla/mux"
)

var formsTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/admin/itemList/itemListList/itemListAdminPage.html",
		"static/admin/itemList/itemListList/listList.html",
		"static/admin/itemList/itemListList/newListDialog.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	formsTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FieldTemplParams struct {
	Title                 string
	DatabaseID            string
	DatabaseName          string
	WorkspaceName         string
	CurrUserIsAdmin       bool
	IsSingleUserWorkspace bool
}

func itemListAdminPage(w http.ResponseWriter, r *http.Request) {

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

	currUserIsAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	templParams := FieldTemplParams{
		Title:                 "Item Lists",
		DatabaseID:            databaseID,
		DatabaseName:          dbInfo.DatabaseName,
		WorkspaceName:         workspaceName,
		CurrUserIsAdmin:       currUserIsAdmin,
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace()}

	if err := formsTemplates.ExecuteTemplate(w, "itemListAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
