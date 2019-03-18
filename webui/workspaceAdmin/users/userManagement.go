package users

import (
	"html/template"
	"net/http"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/workspace"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
	adminCommon "resultra/tracker/webui/workspaceAdmin/common"
)

var formsTemplates *template.Template

func init() {

	baseTemplateFiles := []string{
		"static/workspaceAdmin/users/userManagement.html",
		"static/workspaceAdmin/users/userRegistrationProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		adminCommon.TemplateFileList,
		common.TemplateFileList}

	formsTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type UserTemplParams struct {
	Title         string
	WorkspaceName string
}

func userAdminPage(w http.ResponseWriter, r *http.Request) {

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

	templParams := UserTemplParams{
		Title:         "User Management",
		WorkspaceName: workspaceName}

	if err := formsTemplates.ExecuteTemplate(w, "userMgmtAdminPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
