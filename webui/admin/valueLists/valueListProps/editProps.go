package valueListProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/userRole"
	"resultra/datasheet/server/valueList"
	"resultra/datasheet/webui/thirdParty"

	"resultra/datasheet/server/common/runtimeConfig"
	adminCommon "resultra/datasheet/webui/admin/common"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
)

var formLinkTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/valueLists/valueListProps/editProps.html",
		"static/admin/valueLists/valueListProps/valueListProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		adminCommon.TemplateFileList,
		thirdParty.TemplateFileList,
		generic.TemplateFileList,
		common.TemplateFileList}

	formLinkTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FormLinkTemplParams struct {
	ElemPrefix      string
	Title           string
	DatabaseID      string
	DatabaseName    string
	ValueListID     string
	ValueListName   string
	SiteBaseURL     string
	CurrUserIsAdmin bool
}

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/admin/valueList/{valueListID}", editPropsPage)
}

func editPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	valueListID := vars["valueListID"]

	valueListInfo, err := valueList.GetValueList(valueListID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	dbInfo, err := databaseController.GetDatabaseInfo(valueListInfo.ParentDatabaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	elemPrefix := "valueList_"

	templParams := FormLinkTemplParams{
		ElemPrefix:      elemPrefix,
		Title:           "Value List Settings",
		DatabaseID:      dbInfo.DatabaseID,
		DatabaseName:    dbInfo.DatabaseName,
		ValueListID:     valueListID,
		ValueListName:   valueListInfo.Name,
		SiteBaseURL:     runtimeConfig.GetSiteBaseURL(),
		CurrUserIsAdmin: isAdmin}

	if err := formLinkTemplates.ExecuteTemplate(w, "editValueListPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
