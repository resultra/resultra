// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package itemListProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/databaseController"
	itemListDataModel "resultra/tracker/server/itemList"
	overallUserRole "resultra/tracker/server/userRole"
	adminCommon "resultra/tracker/webui/admin/common"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/workspace"
	"resultra/tracker/webui/admin/itemList/itemListProps/userRole"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/common/recordFilter"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
)

var itemListTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/admin/itemList/itemListProps/editListProps.html",
		"static/admin/itemList/itemListProps/alternateViews.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList,
		userRole.TemplateFileList}

	itemListTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ItemListTemplParams struct {
	Title                    string
	DatabaseID               string
	DatabaseName             string
	WorkspaceName            string
	ListID                   string
	ListName                 string
	CurrUserIsAdmin          bool
	IsSingleUserWorkspace    bool
	FilterPropPanelParams    recordFilter.FilterPanelTemplateParams
	PreFilterPropPanelParams recordFilter.FilterPanelTemplateParams
}

func editListPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	listID := vars["listID"]

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

	listInfo, err := itemListDataModel.GetItemList(trackerDBHandle, listID)
	if err != nil {
		log.Println("Error retrieving item list info:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(trackerDBHandle, listInfo.ParentDatabaseID)
	if dbInfoErr != nil {
		log.Println("Error retrieving item list database info:", dbInfoErr)
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	log.Println("editListPropsPage: viewing/editing admin settings for list ID = ", listID)

	elemPrefix := "itemList_"
	preFilterElemPrefix := "itemListPreFilter_"

	currUserIsAdmin := overallUserRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	templParams := ItemListTemplParams{
		Title:                    "Item List Settings",
		DatabaseID:               listInfo.ParentDatabaseID,
		DatabaseName:             dbInfo.DatabaseName,
		WorkspaceName:            workspaceName,
		ListID:                   listID,
		ListName:                 listInfo.Name,
		CurrUserIsAdmin:          currUserIsAdmin,
		IsSingleUserWorkspace:    runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace,
		FilterPropPanelParams:    recordFilter.NewFilterPanelTemplateParams(elemPrefix),
		PreFilterPropPanelParams: recordFilter.NewFilterPanelTemplateParams(preFilterElemPrefix),
	}

	if err := itemListTemplates.ExecuteTemplate(w, "editItemListPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
