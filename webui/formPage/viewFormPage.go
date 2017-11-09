package formPage

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"

	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/userRole"

	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

type ViewFormPageTemplateParams struct {
	Title           string
	FormID          string
	FormName        string
	DatabaseID      string
	DatabaseName    string
	CurrUserIsAdmin bool
	RecordID        string
	SrcFormButtonID string
	SrcButtonColID  string
}

var viewFormTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/formPage/viewFormPage.html", "static/formPage/common.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}
	viewFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func viewFormPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	formID := vars["formID"]
	recordID := vars["recordID"]

	srcColID, colFound := vars["col"]
	if !colFound {
		srcColID = ""
	}
	srcFrmButtonID, buttonFound := vars["frm"]
	if !buttonFound {
		srcFrmButtonID = ""
	}

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		err := submitFormTemplates.ExecuteTemplate(w, "userSignInPage", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {

		trackerDBHandle, dbErr := databaseWrapper.GetTrackerDatabaseHandle(r)
		if dbErr != nil {
			http.Error(w, dbErr.Error(), http.StatusInternalServerError)
			return
		}

		formDBInfo, getErr := databaseController.GetFormDatabaseInfo(trackerDBHandle, formID)
		if getErr != nil {
			api.WriteErrorResponse(w, getErr)
			return
		}

		isAdmin := userRole.CurrUserIsDatabaseAdmin(r, formDBInfo.DatabaseID)

		templParams := ViewFormPageTemplateParams{Title: "View Form",
			FormID:          formDBInfo.FormID,
			FormName:        formDBInfo.FormName,
			DatabaseID:      formDBInfo.DatabaseID,
			DatabaseName:    formDBInfo.DatabaseName,
			CurrUserIsAdmin: isAdmin,
			RecordID:        recordID,
			SrcFormButtonID: srcFrmButtonID,
			SrcButtonColID:  srcColID}

		if err := viewFormTemplates.ExecuteTemplate(w, "viewFormPage", templParams); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}
