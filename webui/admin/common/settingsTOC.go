package common

import (
	"html/template"
	"net/http"
	"resultra/datasheet/server/common/runtimeConfig"
	"resultra/datasheet/webui/common"
	"resultra/datasheet/webui/generic"
)

var contentTemplates *template.Template

func init() {

	baseTemplateFiles := []string{"static/admin/common/settingsTOC.html"}

	templateFileLists := [][]string{baseTemplateFiles,
		generic.TemplateFileList,
		common.TemplateFileList}

	contentTemplates = generic.ParseTemplatesFromFileLists(templateFileLists)
}

type TOCContentTemplParams struct {
	IsSingleUserWorkspace bool
}

func tocContentLayout(w http.ResponseWriter, r *http.Request) {

	templParams := TOCContentTemplParams{
		IsSingleUserWorkspace: runtimeConfig.CurrRuntimeConfig.IsSingleUserWorkspace}

	if err := contentTemplates.ExecuteTemplate(w, "adminSettingsTOC", templParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
