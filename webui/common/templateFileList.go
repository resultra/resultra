package common

import (
	"resultra/datasheet/webui/common/breadCrumb"
	"resultra/datasheet/webui/common/databaseTOC"
	"resultra/datasheet/webui/common/field"
	"resultra/datasheet/webui/common/formulaEditor"
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/common/recordSort"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/common/include.html"}

	TemplateFileList = append(TemplateFileList, formulaEditor.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, breadCrumb.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, recordFilter.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, databaseTOC.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, recordSort.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, field.TemplateFileList...)

}
