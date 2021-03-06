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

	"github.com/resultra/resultra/server/common/runtimeConfig"

	"github.com/resultra/resultra/server/common/userAuth"
	"github.com/resultra/resultra/server/userRole"
	adminCommon "github.com/resultra/resultra/webui/admin/common"
	"github.com/resultra/resultra/webui/common"
	"github.com/resultra/resultra/webui/generic"
	"github.com/resultra/resultra/webui/thirdParty"

	"github.com/resultra/resultra/webui/admin/common/inputProperties"

	"github.com/resultra/resultra/webui/admin/alerts/alertList"
	"github.com/resultra/resultra/webui/admin/alerts/alertProps"
	"github.com/resultra/resultra/webui/admin/collaborators/collaboratorList"
	"github.com/resultra/resultra/webui/admin/collaborators/collaboratorProps"
	"github.com/resultra/resultra/webui/admin/dashboards"
	"github.com/resultra/resultra/webui/admin/fields/fieldList"
	"github.com/resultra/resultra/webui/admin/fields/fieldProps"
	"github.com/resultra/resultra/webui/admin/formLink/formLinkList"
	"github.com/resultra/resultra/webui/admin/formLink/formLinkProps"
	"github.com/resultra/resultra/webui/admin/forms/formList"
	"github.com/resultra/resultra/webui/admin/general"
	"github.com/resultra/resultra/webui/admin/globals"
	"github.com/resultra/resultra/webui/admin/itemList/itemListList"
	"github.com/resultra/resultra/webui/admin/itemList/itemListProps"
	itemListUserRole "github.com/resultra/resultra/webui/admin/itemList/itemListProps/userRole"

	"github.com/resultra/resultra/webui/admin/tables/colProps"
	"github.com/resultra/resultra/webui/admin/tables/tableList"
	"github.com/resultra/resultra/webui/admin/tables/tableProps"
	"github.com/resultra/resultra/webui/admin/userRole/userRoleList"
	"github.com/resultra/resultra/webui/admin/userRole/userRoleProps"
	"github.com/resultra/resultra/webui/admin/valueLists/valueListList"
	"github.com/resultra/resultra/webui/admin/valueLists/valueListProps"
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
