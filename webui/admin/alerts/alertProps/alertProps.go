package alertProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/alert"
	"resultra/datasheet/server/databaseController"
	adminCommon "resultra/datasheet/webui/admin/common"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var alertTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/alerts/alertProps/alertProps.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}

	alertTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type UserRoleTemplParams struct {
	Title        string
	DatabaseID   string
	DatabaseName string
	AlertID      string
	AlertName    string
}

func editAlertPropsPage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	alertID := vars["alertID"]

	alertInfo, alertErr := alert.GetAlert(alertID)
	if alertErr != nil {
		http.Error(w, alertErr.Error(), http.StatusInternalServerError)
		return

	}

	dbInfo, dbInfoErr := databaseController.GetDatabaseInfo(alertInfo.ParentDatabaseID)
	if dbInfoErr != nil {
		http.Error(w, dbInfoErr.Error(), http.StatusInternalServerError)
		return
	}

	//	elemPrefix := "userRole_"
	templParams := UserRoleTemplParams{
		Title:        "Alert Settings",
		DatabaseID:   dbInfo.DatabaseID,
		DatabaseName: dbInfo.DatabaseName,
		AlertID:      alertInfo.AlertID,
		AlertName:    alertInfo.Name}

	if err := alertTemplates.ExecuteTemplate(w, "alertPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
