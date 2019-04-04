// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package formLinkProps

import (
	"github.com/resultra/resultra/server/common/runtimeConfig"
	"github.com/resultra/resultra/server/databaseController"
	"github.com/resultra/resultra/server/formLink"
	adminCommon "github.com/resultra/resultra/webui/admin/common"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/resultra/resultra/server/common/databaseWrapper"
	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/userRole"
	"github.com/resultra/resultra/server/workspace"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/common/defaultValues"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/thirdParty"
)

var formLinkTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/formLink/formLinkProps/editProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		adminCommon.TemplateFileList,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}

	formLinkTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FormLinkTemplParams struct {
	ElemPrefix              string
	Title                   string
	DatabaseID              string
	DatabaseName            string
	WorkspaceName           string
	CurrUserIsAdmin         bool
	IsSingleUserWorkspace   bool
	LinkID                  string
	LinkName                string
	SiteBaseURL             string
	DefaultValuePanelParams defaultValues.DefaultValuesPanelTemplateParams
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/formLink/{linkID}", editPropsPage)
}

func editPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	linkID := vars["linkID"]

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

	linkInfo, err := formLink.GetFormLink(trackerDBHandle, linkID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	formDBInfo, err := databaseController.GetFormDatabaseInfo(trackerDBHandle, linkInfo.FormID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("editPropsPage: viewing/editing admin settings for form link ID = ", linkID)

	elemPrefix := "formLink_"

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, formDBInfo.DatabaseID)

	templParams := FormLinkTemplParams{
		ElemPrefix:              elemPrefix,
		Title:                   "Form Link Settings",
		DatabaseID:              formDBInfo.DatabaseID,
		DatabaseName:            formDBInfo.DatabaseName,
		WorkspaceName:           workspaceName,
		LinkID:                  linkID,
		LinkName:                linkInfo.Name,
		CurrUserIsAdmin:         isAdmin,
		IsSingleUserWorkspace:   runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace(),
		SiteBaseURL:             runtimeConfig.GetSiteBaseURL(),
		DefaultValuePanelParams: defaultValues.NewDefaultValuesTemplateParams(elemPrefix)}

	if err := formLinkTemplates.ExecuteTemplate(w, "editFormLinkPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
