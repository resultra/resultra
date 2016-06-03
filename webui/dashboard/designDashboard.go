package dashboard

import (
	"appengine"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/generic/api"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/dashboard/components"
	"resultra/datasheet/webui/generic"
)

var designDashboardTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/dashboard/dashboardCommon.html",
		"static/dashboard/dashboardProps.html",
		"static/dashboard/designDashboard.html"}

	templateFileLists := [][]string{
		baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList,
		components.TemplateFileList}
	designDashboardTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func designDashboard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	dashboardID := vars["dashboardID"]
	log.Println("Design Dashboard: editing for dashboard with ID = ", dashboardID)

	appEngContext := appengine.NewContext(r)
	dashboardForDesign, getErr := dashboard.GetDashboard(appEngContext, dashboardID)
	if getErr != nil {
		api.WriteErrorResponse(w, getErr)
		return
	}

	templParams := createDashboardTemplateParams(dashboardForDesign)

	err := designDashboardTemplates.ExecuteTemplate(w, "designDashboard", templParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
