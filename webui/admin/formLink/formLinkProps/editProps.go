package formLinkProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/tracker/server/common/runtimeConfig"
	"resultra/tracker/server/databaseController"
	"resultra/tracker/server/formLink"
	adminCommon "resultra/tracker/webui/admin/common"

	"resultra/tracker/server/common/databaseWrapper"
	"resultra/tracker/server/common/userAuth"
	"resultra/tracker/server/userRole"
	"resultra/tracker/server/workspace"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/common/defaultValues"
	"resultra/tracker/webui/generic"
	"resultra/tracker/webui/thirdParty"
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
		IsSingleUserWorkspace:   runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace,
		SiteBaseURL:             runtimeConfig.GetSiteBaseURL(),
		DefaultValuePanelParams: defaultValues.NewDefaultValuesTemplateParams(elemPrefix)}

	if err := formLinkTemplates.ExecuteTemplate(w, "editFormLinkPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
