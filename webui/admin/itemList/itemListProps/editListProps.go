package itemListProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"
	itemListDataModel "resultra/datasheet/server/itemList"

	"resultra/datasheet/webui/admin/itemList/itemListProps/userRole"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/generic"
)

var itemListTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/admin/itemList/itemListProps/editListProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList,
		userRole.TemplateFileList}

	itemListTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ItemListTemplParams struct {
	Title                    string
	DatabaseID               string
	DatabaseName             string
	ListID                   string
	ListName                 string
	FilterPropPanelParams    recordFilter.FilterPanelTemplateParams
	PreFilterPropPanelParams recordFilter.FilterPanelTemplateParams
}

func editListPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	listID := vars["listID"]

	listInfo, err := itemListDataModel.GetItemList(listID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(listInfo.ParentDatabaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	log.Println("editListPropsPage: viewing/editing admin settings for list ID = ", listID)

	elemPrefix := "itemList_"
	preFilterElemPrefix := "itemListPreFilter_"
	templParams := ItemListTemplParams{
		Title:                    "Item List Settings",
		DatabaseID:               listInfo.ParentDatabaseID,
		DatabaseName:             dbInfo.DatabaseName,
		ListID:                   listID,
		ListName:                 listInfo.Name,
		FilterPropPanelParams:    recordFilter.NewFilterPanelTemplateParams(elemPrefix),
		PreFilterPropPanelParams: recordFilter.NewFilterPanelTemplateParams(preFilterElemPrefix),
	}

	if err := itemListTemplates.ExecuteTemplate(w, "editItemListPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
