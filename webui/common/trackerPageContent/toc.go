package trackerPageContent

import (
	"html/template"
	"net/http"
	"resultra/tracker/webui/generic"
)

var tocTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/common/trackerPageContent/toc.html"}

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
