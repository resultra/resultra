// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package alertProps

import (
	"html/template"
	"net/http"
	"resultra/tracker/server/alert"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/databaseController"
	adminCommon "resultra/tracker/webui/admin/common"

	"github.com/gorilla/mux"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/userRole"
	"resultra/tracker/server/workspace"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/common/field"
	"resultra/tracker/webui/common/recordFilter"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
)

var alertTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/alerts/alertProps/alertProps.html",
		"static/admin/alerts/alertProps/events.html",
		"static/admin/alerts/alertProps/recipients.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}

	alertTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type AlertTemplParams struct {
	Title                           string
	DatabaseID                      string
	DatabaseName                    string
	WorkspaceName                   string
	AlertID                         string
	AlertName                       string
	CurrUserIsAdmin                 bool
	IsSingleUserWorkspace           bool
	FieldSelectionParams            field.FieldSelectionDropdownTemplateParams
	TriggerConditionPropPanelParams recordFilter.FilterPanelTemplateParams
}

func editAlertPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	alertID := vars["alertID"]

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

	alertInfo, alertErr := alert.GetAlert(trackerDBHandle, alertID)
	if alertErr != nil {
		http.Error(w, alertErr.Error(), http.StatusInternalServerError)
		return

	}

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(trackerDBHandle, alertInfo.ParentDatabaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
		return
	}

	fieldSelectionParams := field.FieldSelectionDropdownTemplateParams{
		ElemPrefix:     "alertCondition_",
		ButtonTitle:    "Add Trigger Event",
		ButtonIconName: "glyphicon-plus-sign"}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	triggerConditionElemPrefix_ := "alertTriggerCondition_"
	templParams := AlertTemplParams{
		Title:                           "Alert Settings",
		DatabaseID:                      dbInfo.DatabaseID,
		DatabaseName:                    dbInfo.DatabaseName,
		WorkspaceName:                   workspaceName,
		AlertID:                         alertInfo.AlertID,
		AlertName:                       alertInfo.Name,
		CurrUserIsAdmin:                 isAdmin,
		IsSingleUserWorkspace:           runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace(),
		FieldSelectionParams:            fieldSelectionParams,
		TriggerConditionPropPanelParams: recordFilter.NewFilterPanelTemplateParams(triggerConditionElemPrefix_)}

	if err := alertTemplates.ExecuteTemplate(w, "alertPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
