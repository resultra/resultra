package itemList

import (
	"html/template"
	"net/http"
	"resultra/tracker/webui/common"
	"resultra/tracker/webui/generic"
)

var propertyTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/itemList/properties.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList}

	propertyTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

func propertySidebarContent(w http.ResponseWriter, r *http.Request) {

	if err := propertyTemplates.ExecuteTemplate(w, "listViewProps", ViewListTemplParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
