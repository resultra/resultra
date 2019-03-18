package itemList

import (
	"html/template"
	"net/http"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
)

var contentTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/itemList/contentLayout.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		common.TemplateFileList}

	contentTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ListContentTemplParams struct{}

func listViewContentLayout(w http.ResponseWriter, r *http.Request) {

	templParams := ListContentTemplParams{}

	if err := contentTemplates.ExecuteTemplate(w, "listViewContentLayout", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func listViewOffPageContent(w http.ResponseWriter, r *http.Request) {

	templParams := ListContentTemplParams{}

	if err := contentTemplates.ExecuteTemplate(w, "listViewOffPageContent", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
