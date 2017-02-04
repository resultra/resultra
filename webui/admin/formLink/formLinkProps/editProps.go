package formLinkProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/formLink"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/common/defaultValues"
	"resultra/datasheet/webui/generic"
)

var formLinkTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/formLink/formLinkProps/editProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList}

	formLinkTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FormLinkTemplParams struct {
	ElemPrefix              string
	Title                   string
	DatabaseID              string
	DatabaseName            string
	LinkID                  string
	LinkName                string
	DefaultValuePanelParams defaultValues.DefaultValuesPanelTemplateParams
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/formLink/{linkID}", editPropsPage)
}

func editPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	linkID := vars["linkID"]

	linkInfo, err := formLink.GetFormLink(linkID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	formDBInfo, err := databaseController.GetFormDatabaseInfo(linkInfo.FormID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Println("editPropsPage: viewing/editing admin settings for form link ID = ", linkID)

	elemPrefix := "formLink_"

	templParams := FormLinkTemplParams{
		ElemPrefix:              elemPrefix,
		Title:                   "Form Link Settings",
		DatabaseID:              formDBInfo.DatabaseID,
		DatabaseName:            formDBInfo.DatabaseName,
		LinkID:                  linkID,
		LinkName:                linkInfo.Name,
		DefaultValuePanelParams: defaultValues.NewDefaultValuesTemplateParams(elemPrefix)}

	if err := formLinkTemplates.ExecuteTemplate(w, "editFormLinkPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
