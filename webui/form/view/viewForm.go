package view

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"

	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/itemList"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/form/common/timeline"
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
		components.TemplateFileList,
		timeline.TemplateFileList}
	viewFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ViewFormTemplateParams struct {
	Title                string
	FormID               string
	ListID               string
	TableID              string
	DatabaseID           string
	DatabaseName         string
	ListName             string
	FilteringPanelParams propertiesSidebar.PanelTemplateParams
	SortPanelParams      propertiesSidebar.PanelTemplateParams
	ComponentParams      components.ComponentViewTemplateParams
}

func ViewList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	listID := vars["listID"]

	_, authErr := userAuth.GetCurrentUserInfo(r)
	if authErr != nil {
		err := viewFormTemplates.ExecuteTemplate(w, "userSignInPage", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		log.Println("view list: : listID ID = %v", listID)

		listInfo, err := itemList.GetItemList(listID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		formID := listInfo.FormID

		formDBInfo, getErr := databaseController.GetFormDatabaseInfo(formID)
		if getErr != nil {
			api.WriteErrorResponse(w, getErr)
			return
		}

		templParams := ViewFormTemplateParams{Title: "View List",
			FormID:       formID,
			ListID:       listID,
			TableID:      formDBInfo.TableID,
			DatabaseID:   formDBInfo.DatabaseID,
			DatabaseName: formDBInfo.DatabaseName,
			ListName:     listInfo.Name,
			FilteringPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering",
				PanelID: "viewFormFiltering"},
			SortPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Sorting",
				PanelID: "viewFormSorting"},
			ComponentParams: components.ViewTemplateParams}

		if err := viewFormTemplates.ExecuteTemplate(w, "viewForm", templParams); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}
