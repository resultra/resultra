package alertProps

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/server/alert"
	"resultra/datasheet/server/databaseController"
	adminCommon "resultra/datasheet/webui/admin/common"

	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/common/field"
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

var alertTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/alerts/alertProps/alertProps.html",
		"static/admin/alerts/alertProps/events.html",
		"static/admin/alerts/alertProps/recipients.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		adminCommon.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList}

	alertTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type AlertTemplParams struct {
	Title                           string
	DatabaseID                      string
	DatabaseName                    string
	AlertID                         string
	AlertName                       string
	FieldSelectionParams            field.FieldSelectionDropdownTemplateParams
	TriggerConditionPropPanelParams recordFilter.FilterPanelTemplateParams
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

	fieldSelectionParams := field.FieldSelectionDropdownTemplateParams{
		ElemPrefix:     "alertCondition_",
		ButtonTitle:    "Add Trigger Event",
		ButtonIconName: "glyphicon-plus-sign"}

	triggerConditionElemPrefix_ := "alertTriggerCondition_"
	templParams := AlertTemplParams{
		Title:                           "Alert Settings",
		DatabaseID:                      dbInfo.DatabaseID,
		DatabaseName:                    dbInfo.DatabaseName,
		AlertID:                         alertInfo.AlertID,
		AlertName:                       alertInfo.Name,
		FieldSelectionParams:            fieldSelectionParams,
		TriggerConditionPropPanelParams: recordFilter.NewFilterPanelTemplateParams(triggerConditionElemPrefix_)}

	if err := alertTemplates.ExecuteTemplate(w, "alertPropsPage", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
