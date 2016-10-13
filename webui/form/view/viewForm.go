package view

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/form/components"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

var viewFormTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/form/view/viewForm.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList,
		components.TemplateFileList}
	viewFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ViewFormTemplateParams struct {
	Title                string
	FormID               string
	TableID              string
	DatabaseID           string
	DatabaseName         string
	FormName             string
	FilteringPanelParams propertiesSidebar.PanelTemplateParams
	SortPanelParams      propertiesSidebar.PanelTemplateParams
}

func ViewForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	formID := vars["formID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		err := viewFormTemplates.ExecuteTemplate(w, "userSignInPage", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		log.Println("view form: : form ID = %v", formID)

		formDBInfo, getErr := databaseController.GetFormDatabaseInfo(formID)
		if getErr != nil {
			api.WriteErrorResponse(w, getErr)
			return
		}

		templParams := ViewFormTemplateParams{Title: "View Form",
			FormID:       formID,
			TableID:      formDBInfo.TableID,
			DatabaseID:   formDBInfo.DatabaseID,
			DatabaseName: formDBInfo.DatabaseName,
			FormName:     formDBInfo.FormName,
			FilteringPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering",
				PanelID: "viewFormFiltering"},
			SortPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Sorting",
				PanelID: "viewFormSorting"}}

		err := viewFormTemplates.ExecuteTemplate(w, "viewForm", templParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}
