// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package valueListProps

import (
	"github.com/resultra/resultra/server/databaseController"
	"github.com/resultra/resultra/server/userRole"
	"github.com/resultra/resultra/server/valueList"
	"github.com/resultra/resultra/webui/thirdParty"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/resultra/resultra/server/common/runtimeConfig"
	"github.com/resultra/resultra/server/workspace"
	adminCommon "github.com/resultra/resultra/webui/admin/common"

	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
)

var formLinkTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/valueLists/valueListProps/editProps.html",
		"static/admin/valueLists/valueListProps/valueListProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		adminCommon.TemplateFileList,
		thirdParty.TemplateFileList,
		generic.TemplateFileList,
		common.TemplateFileList}

	formLinkTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FormLinkTemplParams struct {
	ElemPrefix            string
	Title                 string
	DatabaseID            string
	DatabaseName          string
	WorkspaceName         string
	ValueListID           string
	ValueListName         string
	SiteBaseURL           string
	CurrUserIsAdmin       bool
	IsSingleUserWorkspace bool
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/valueList/{valueListID}", editPropsPage)
}

func editPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	valueListID := vars["valueListID"]

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

	valueListInfo, err := valueList.GetValueList(trackerDBHandle, valueListID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	dbInfo, err := databaseController.GetDatabaseInfo(trackerDBHandle, valueListInfo.ParentDatabaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	elemPrefix := "valueList_"

	templParams := FormLinkTemplParams{
		ElemPrefix:            elemPrefix,
		Title:                 "Value List Settings",
		DatabaseID:            dbInfo.DatabaseID,
		DatabaseName:          dbInfo.DatabaseName,
		WorkspaceName:         workspaceName,
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace(),
		ValueListID:           valueListID,
		ValueListName:         valueListInfo.Name,
		SiteBaseURL:           runtimeConfig.GetSiteBaseURL(),
		CurrUserIsAdmin:       isAdmin}

	if err := formLinkTemplates.ExecuteTemplate(w, "editValueListPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
