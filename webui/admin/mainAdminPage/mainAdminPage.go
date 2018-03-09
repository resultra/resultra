package mainAdminPage

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"

	"resultra/datasheet/server/common/runtimeConfig"

	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/server/userRole"
	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"

	"resultra/datasheet/webui/admin/common/inputProperties"

	"resultra/datasheet/webui/admin/alerts/alertList"
	"resultra/datasheet/webui/admin/alerts/alertProps"
	"resultra/datasheet/webui/admin/collaborators/collaboratorList"
	"resultra/datasheet/webui/admin/collaborators/collaboratorProps"
	"resultra/datasheet/webui/admin/dashboards"
	"resultra/datasheet/webui/admin/fields/fieldList"
	"resultra/datasheet/webui/admin/fields/fieldProps"
	"resultra/datasheet/webui/admin/formLink/formLinkList"
	"resultra/datasheet/webui/admin/formLink/formLinkProps"
	"resultra/datasheet/webui/admin/forms/formList"
	"resultra/datasheet/webui/admin/general"
	"resultra/datasheet/webui/admin/globals"
	"resultra/datasheet/webui/admin/itemList/itemListList"
	"resultra/datasheet/webui/admin/itemList/itemListProps"
	itemListUserRole "resultra/datasheet/webui/admin/itemList/itemListProps/userRole"

	"resultra/datasheet/webui/admin/tables/colProps"
	"resultra/datasheet/webui/admin/tables/tableList"
	"resultra/datasheet/webui/admin/tables/tableProps"
	"resultra/datasheet/webui/admin/userRole/userRoleList"
	"resultra/datasheet/webui/admin/userRole/userRoleProps"
	"resultra/datasheet/webui/admin/valueLists/valueListList"
	"resultra/datasheet/webui/admin/valueLists/valueListProps"
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
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace}

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
