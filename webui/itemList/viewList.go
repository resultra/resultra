package itemList

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
	"resultra/datasheet/webui/common/form/components"
	"resultra/datasheet/webui/common/form/submit"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/generic/propertiesSidebar"
	"resultra/datasheet/webui/itemList/common/timeline"
)

var viewListTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/itemList/viewList.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList,
		components.TemplateFileList,
		submit.TemplateFileList,
		timeline.TemplateFileList}
	viewListTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ViewListTemplateParams struct {
	Title                string
	FormID               string
	ListID               string
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
		err := viewListTemplates.ExecuteTemplate(w, "userSignInPage", nil)
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

		templParams := ViewListTemplateParams{Title: "View List",
			FormID:       formID,
			ListID:       listID,
			DatabaseID:   formDBInfo.DatabaseID,
			DatabaseName: formDBInfo.DatabaseName,
			ListName:     listInfo.Name,
			FilteringPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering",
				PanelID: "viewFormFiltering"},
			SortPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Sorting",
				PanelID: "viewFormSorting"},
			ComponentParams: components.ViewTemplateParams}

		if err := viewListTemplates.ExecuteTemplate(w, "viewList", templParams); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}
