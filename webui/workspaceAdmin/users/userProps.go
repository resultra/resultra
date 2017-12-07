package users

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

var userTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/workspaceAdmin/users/userProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	userTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type UserPropsTemplParams struct {
	Title         string
	WorkspaceName string
}

func userPropsPage(w http.ResponseWriter, r *http.Request) {

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

	templParams := UserPropsTemplParams{
		Title:         "User Account Settings",
		WorkspaceName: workspaceName}

	if err := userTemplates.ExecuteTemplate(w, "userPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
