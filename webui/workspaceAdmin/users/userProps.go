package users

import (
	"github.com/gorilla/mux"
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
	UserID        string
}

func userPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userID := vars["userID"]

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
		WorkspaceName: workspaceName,
		UserID:        userID}

	if err := userTemplates.ExecuteTemplate(w, "userPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
