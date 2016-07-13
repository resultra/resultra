package common

import (
	"resultra/datasheet/webui/common/breadCrumb"
	"resultra/datasheet/webui/common/databaseTOC"
	"resultra/datasheet/webui/common/formulaEditor"
	"resultra/datasheet/webui/common/helper"
	"resultra/datasheet/webui/common/objectSelection"
	"resultra/datasheet/webui/common/recordFilter"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{"static/common/include.html"}

	TemplateFileList = append(TemplateFileList, objectSelection.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, formulaEditor.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, breadCrumb.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, helper.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, recordFilter.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, databaseTOC.TemplateFileList...)

}
