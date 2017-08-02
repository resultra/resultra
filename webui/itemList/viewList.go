package itemList

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"

	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"
	"resultra/datasheet/server/itemList"
	"resultra/datasheet/server/userRole"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/common/form/components"
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/generic/propertiesSidebar"
	itemListCommon "resultra/datasheet/webui/itemList/common"
	"resultra/datasheet/webui/thirdParty"
)

var viewListTemplates *template.Template

func init() {
	baseTemplateFiles := []string{"static/itemList/viewList.html",
		"static/itemList/listItems.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		thirdParty.TemplateFileList,
		generic.TemplateFileList,
		common.TemplateFileList,
		itemListCommon.TemplateFileList}
	viewListTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ViewListTemplateParams struct {
	Title                   string
	ListID                  string
	DatabaseID              string
	DatabaseName            string
	CurrUserIsAdmin         bool
	ListName                string
	ListPrivileges          string
	DisplayPanelParams      propertiesSidebar.PanelTemplateParams
	FilteringPanelParams    propertiesSidebar.PanelTemplateParams
	SortPanelParams         propertiesSidebar.PanelTemplateParams
	ComponentParams         components.ComponentViewTemplateParams
	FilterConfigPanelParams recordFilter.FilterPanelTemplateParams
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
		log.Println("view list: listID ID = %v", listID)

		listInfo, err := itemList.GetItemList(listID)
		if err != nil {
			api.WriteErrorResponse(w, err)
			return
		}

		listPrivs, privsErr := userRole.GetCurrentUserItemListPrivs(r, listInfo.ParentDatabaseID, listID)
		if privsErr != nil {
			api.WriteErrorResponse(w, privsErr)
			return
		}
		if listPrivs == userRole.ListRolePrivsNone {
			api.WriteErrorResponse(w, fmt.Errorf("Invalid permissions loading page. No permissions to view or edit this list."))
			return
		}

		dbInfo, getErr := databaseController.GetDatabaseInfo(listInfo.ParentDatabaseID)
		if getErr != nil {
			api.WriteErrorResponse(w, getErr)
			return
		}

		elemPrefix := "form_"

		isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

		templParams := ViewListTemplateParams{Title: "View List",
			ListID:          listID,
			DatabaseID:      dbInfo.DatabaseID,
			DatabaseName:    dbInfo.DatabaseName,
			CurrUserIsAdmin: isAdmin,
			ListName:        listInfo.Name,
			ListPrivileges:  listPrivs,
			DisplayPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "View",
				PanelID: "viewListDisplay"},
			FilteringPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Filtering",
				PanelID: "viewFormFiltering"},
			FilterConfigPanelParams: recordFilter.NewFilterPanelTemplateParams(elemPrefix),
			SortPanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Sorting",
				PanelID: "viewFormSorting"},
			ComponentParams: components.ViewTemplateParams}

		if err := viewListTemplates.ExecuteTemplate(w, "viewList", templParams); err != nil {
			api.WriteErrorResponse(w, err)
		}

	}
}
