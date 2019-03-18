package itemView

import (
	"html/template"
	"net/http"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
)

var existingContentTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/itemView/existingItemContentLayout.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		common.TemplateFileList}

	existingContentTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ExistingContentTemplParams struct{}

func existingItemContentLayout(w http.ResponseWriter, r *http.Request) {

	templParams := ExistingContentTemplParams{}

	if err := existingContentTemplates.ExecuteTemplate(w, "existingItemContentLayout", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func existingItemOffPageContent(w http.ResponseWriter, r *http.Request) {

	templParams := ExistingContentTemplParams{}

	if err := existingContentTemplates.ExecuteTemplate(w, "existingItemOffPageContent", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
