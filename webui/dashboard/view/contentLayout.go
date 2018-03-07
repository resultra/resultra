package view

import (
	"html/template"
	"net/http"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/dashboard/components"
	"resultra/datasheet/webui/generic"
)

var contentTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/dashboard/view/contentLayout.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList,
		components.TemplateFileList}

	contentTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type ContentTemplParams struct{}

func dashboardContentLayout(w http.ResponseWriter, r *http.Request) {

	templParams := ContentTemplParams{}

	if err := contentTemplates.ExecuteTemplate(w, "dashboardContentLayout", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func dashboardSidebarLayout(w http.ResponseWriter, r *http.Request) {

	if err := contentTemplates.ExecuteTemplate(w, "dashboardSidebarContent", ViewTemplateParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
