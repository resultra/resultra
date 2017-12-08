package fieldProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/databaseController"
	adminCommon "resultra/datasheet/webui/admin/common"

	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/field"
	"resultra/datasheet/server/common/userAuth"
	"resultra/datasheet/server/userRole"
	"resultra/datasheet/server/workspace"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var fieldTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/admin/fields/fieldProps/fieldProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		adminCommon.TemplateFileList,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}

	fieldTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type FieldTemplParams struct {
	Title           string
	DatabaseID      string
	DatabaseName    string
	WorkspaceName   string
	FieldID         string
	FieldName       string
	CurrUserIsAdmin bool
}

func editFieldPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fieldID := vars["fieldID"]

	log.Println("editFieldPropsPage: viewing/editing admin settings for field ID = ", fieldID)

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

	workspaceName, workspaceErr := workspace.GetWorkspaceName(trackerDBHandle)
	if workspaceErr != nil {
		http.Error(w, workspaceErr.Error(), http.StatusInternalServerError)
		return
	}

	fieldInfo, err := field.GetField(trackerDBHandle, fieldID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(trackerDBHandle, fieldInfo.ParentDatabaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
	}

	isAdmin := userRole.CurrUserIsDatabaseAdmin(r, dbInfo.DatabaseID)

	//	elemPrefix := "userRole_"
	templParams := FieldTemplParams{
		Title:           "Field Settings",
		DatabaseID:      fieldInfo.ParentDatabaseID,
		DatabaseName:    dbInfo.DatabaseName,
		WorkspaceName:   workspaceName,
		FieldID:         fieldID,
		FieldName:       fieldInfo.Name,
		CurrUserIsAdmin: isAdmin}

	if err := fieldTemplates.ExecuteTemplate(w, "editFieldPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
