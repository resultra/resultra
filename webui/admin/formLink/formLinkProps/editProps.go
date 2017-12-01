package formLinkProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/formLink"
	adminCommon "resultra/datasheet/webui/admin/common"

	"resultra/datasheet/server/common/runtimeConfig"

	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/userRole"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/common/defaultValues"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
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
	CurrUserIsAdmin         bool
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
		LinkID:                  linkID,
		LinkName:                linkInfo.Name,
		CurrUserIsAdmin:         isAdmin,
		SiteBaseURL:             runtimeConfig.GetSiteBaseURL(),
		DefaultValuePanelParams: defaultValues.NewDefaultValuesTemplateParams(elemPrefix)}

	if err := formLinkTemplates.ExecuteTemplate(w, "editFormLinkPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
