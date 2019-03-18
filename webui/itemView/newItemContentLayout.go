package itemView

import (
	"html/template"
	"net/http"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
)

var contentTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/itemView/newItemContentLayout.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		common.TemplateFileList}

	contentTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ContentTemplParams struct{}

func newItemContentLayout(w http.ResponseWriter, r *http.Request) {

	templParams := ContentTemplParams{}

	if err := contentTemplates.ExecuteTemplate(w, "newItemContentLayout", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func newItemOffPageContent(w http.ResponseWriter, r *http.Request) {

	templParams := ContentTemplParams{}

	if err := contentTemplates.ExecuteTemplate(w, "newItemOffPageContent", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
