package trackerTOC

import (
	"html/template"
	"net/http"
	"resultra/datasheet/webui/common/alert"
	"resultra/datasheet/webui/generic"
)

var headerTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/common/trackerTOC/headerButtons.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		alert.TemplateFileList}

	headerTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type HeaderTemplParams struct{}

func headerButtonsContent(w http.ResponseWriter, r *http.Request) {

	templParams := HeaderTemplParams{}

	if err := headerTemplates.ExecuteTemplate(w, "trackerHeaderButtons", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
