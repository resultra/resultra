package trackerTOC

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"resultra/datasheet/webui/generic"
)

func RegisterHTTPHandlers(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/common/trackerTOC/toc", tocPageContent)
}

var tocTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/common/trackerTOC/toc.html"}

	templateFileLists := [][]string{baseTemplateFiles}

	tocTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TOCTemplParams struct{}

func tocPageContent(w http.ResponseWriter, r *http.Request) {

	templParams := TOCTemplParams{}

	if err := tocTemplates.ExecuteTemplate(w, "databaseTOC", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
