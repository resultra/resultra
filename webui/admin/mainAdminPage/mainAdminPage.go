// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package mainAdminPage

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"resultra/tracker/server/common/runtimeConfig"

	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/userRole"
	adminCommon "resultra/tracker/webui/admin/common"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"

	"resultra/tracker/webui/admin/common/inputProperties"

	"resultra/tracker/webui/admin/alerts/alertList"
	"resultra/tracker/webui/admin/alerts/alertProps"
	"resultra/tracker/webui/admin/collaborators/collaboratorList"
	"resultra/tracker/webui/admin/collaborators/collaboratorProps"
	"resultra/tracker/webui/admin/dashboards"
	"resultra/tracker/webui/admin/fields/fieldList"
	"resultra/tracker/webui/admin/fields/fieldProps"
	"resultra/tracker/webui/admin/formLink/formLinkList"
	"resultra/tracker/webui/admin/formLink/formLinkProps"
	"resultra/tracker/webui/admin/forms/formList"
	"resultra/tracker/webui/admin/general"
	"resultra/tracker/webui/admin/globals"
	"resultra/tracker/webui/admin/itemList/itemListList"
	"resultra/tracker/webui/admin/itemList/itemListProps"
	itemListUserRole "resultra/tracker/webui/admin/itemList/itemListProps/userRole"

	"resultra/tracker/webui/admin/tables/colProps"
	"resultra/tracker/webui/admin/tables/tableList"
	"resultra/tracker/webui/admin/tables/tableProps"
	"resultra/tracker/webui/admin/userRole/userRoleList"
	"resultra/tracker/webui/admin/userRole/userRoleProps"
	"resultra/tracker/webui/admin/valueLists/valueListList"
	"resultra/tracker/webui/admin/valueLists/valueListProps"
)

var mainAdminPageTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/mainAdminPage/mainAdminPage.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		inputProperties.TemplateFileList,
		common.TemplateFileList,
		general.TemplateFileList,
		formList.TemplateFileList,
		formLinkList.TemplateFileList,
		formLinkProps.TemplateFileList,
		tableList.TemplateFileList,
		colProps.TemplateFileList,
		tableProps.TemplateFileList,
		itemListList.TemplateFileList,
		itemListProps.TemplateFileList,
		itemListUserRole.TemplateFileList,
		fieldList.TemplateFileList,
		fieldProps.TemplateFileList,
		valueListList.TemplateFileList,
		valueListProps.TemplateFileList,
		dashboards.TemplateFileList,
		alertList.TemplateFileList,
		alertProps.TemplateFileList,
		userRoleList.TemplateFileList,
		userRoleProps.TemplateFileList,
		collaboratorList.TemplateFileList,
		collaboratorProps.TemplateFileList,
		globals.TemplateFileList}

	mainAdminPageTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	DatabaseID            string
	CurrUserIsAdmin       bool
	IsSingleUserWorkspace bool
}

func mainAdminPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	databaseID := vars["databaseID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	currUserIsAdmin := userRole.CurrUserIsDatabaseAdmin(r, databaseID)

	templParams := TemplParams{
		DatabaseID:            databaseID,
		CurrUserIsAdmin:       currUserIsAdmin,
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.SingleUserWorkspace()}

	if err := mainAdminPageTemplates.ExecuteTemplate(w, "mainAdminPageContent", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

type OffPageContentTemplParams struct{}

func mainAdminPageOffPageContent(w http.ResponseWriter, r *http.Request) {

	templParams := OffPageContentTemplParams{}

	if err := mainAdminPageTemplates.ExecuteTemplate(w, "mainAdminPageOffpageContent", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
