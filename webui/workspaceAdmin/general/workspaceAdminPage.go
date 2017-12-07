package general

import (
	"html/template"
	"net/http"

	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/workspace"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
	adminCommon "resultra/datasheet/webui/workspaceAdmin/common"
)

var generalTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/workspaceAdmin/general/workspaceAdminPage.html",
		"static/workspaceAdmin/general/workspaceNameProperty.html",
		"static/workspaceAdmin/general/userAccountProperties.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList,
		adminCommon.TemplateFileList}

	generalTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TemplParams struct {
	Title         string
	WorkspaceName string
}

func workspaceAdminPage(w http.ResponseWriter, r *http.Request) {

	userInfo, authErr := userAuth.GetCurrentUserInfo(r)
	if (authErr != nil) || (!userInfo.IsWorkspaceAdmin) {
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

	templParams := TemplParams{
		Title:         "General Settings",
		WorkspaceName: workspaceName}

	if err := generalTemplates.ExecuteTemplate(w, "workspaceAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
