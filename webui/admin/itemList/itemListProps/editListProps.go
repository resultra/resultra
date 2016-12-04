package itemListProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"
	itemListDataModel "resultra/datasheet/server/itemList"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
)

var itemListTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/admin/itemList/itemListProps/editListProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList}

	itemListTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ItemListTemplParams struct {
	Title        string
	DatabaseID   string
	DatabaseName string
	ListID       string
}

func editListPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	listID := vars["listID"]

	databaseID, err := itemListDataModel.GetItemListDatabaseID(listID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(databaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	log.Println("editListPropsPage: viewing/editing admin settings for list ID = ", listID)

	templParams := ItemListTemplParams{
		Title:        "Item List Settings",
		DatabaseID:   databaseID,
		DatabaseName: dbInfo.DatabaseName,
		ListID:       listID}

	if err := itemListTemplates.ExecuteTemplate(w, "editItemListPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
