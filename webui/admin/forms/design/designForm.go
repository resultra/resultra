package design

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/common/databaseWrapper"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/server/generic/userAuth"
	adminCommon "resultra/datasheet/webui/admin/common"
	"resultra/datasheet/webui/admin/common/inputProperties"
	"resultra/datasheet/webui/admin/forms/design/properties"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/common/form/components"
	"resultra/datasheet/webui/generic"
	"resultra/datasheet/webui/thirdParty"
)

// Parse all the HTML templates at once. Individual templates can then
// be referenced throughout this package using htmlTemplates.ExecuteTemplate(...)
var designFormTemplates *template.Template

func init() {
	//	designFormTemplateFiles := []string{}

	baseTemplateFiles := []string{"static/admin/forms/design/designForm.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		thirdParty.TemplateFileList,
		common.TemplateFileList,
		inputProperties.TemplateFileList,
		components.TemplateFileList,
		properties.TemplateFileList,
		adminCommon.TemplateFileList}
	designFormTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func DesignForm(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	formID := vars["formID"]
	log.Println("Design Form: editing for form with ID = ", formID)

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

	formDBInfo, getErr := databaseController.GetFormDatabaseInfo(trackerDBHandle, formID)
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	}

	templParams := createDesignFormTemplateParams(r, formDBInfo)

	err := designFormTemplates.ExecuteTemplate(w, "designForm", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
