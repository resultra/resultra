package common

import (
	"resultra/datasheet/webui/common/alert"
	"resultra/datasheet/webui/common/attachment"
	"resultra/datasheet/webui/common/breadCrumb"
	"resultra/datasheet/webui/common/conditionalFormat"
	"resultra/datasheet/webui/common/database"
	"resultra/datasheet/webui/common/defaultValues"
	"resultra/datasheet/webui/common/field"
	"resultra/datasheet/webui/common/form"
	"resultra/datasheet/webui/common/formulaEditor"
	"resultra/datasheet/webui/common/helpPopup"
	"resultra/datasheet/webui/common/itemList"
	"resultra/datasheet/webui/common/recordFilter"
	"resultra/datasheet/webui/common/recordSort"
	"resultra/datasheet/webui/common/valueList"
	"resultra/datasheet/webui/common/valueThreshold"
)

var TemplateFileList []string

func init() {

	TemplateFileList = []string{}

	TemplateFileList = append(TemplateFileList, attachment.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, formulaEditor.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, breadCrumb.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, recordFilter.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, database.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, recordSort.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, field.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, form.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, defaultValues.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueThreshold.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, valueList.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, itemList.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, conditionalFormat.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, helpPopup.TemplateFileList...)
	TemplateFileList = append(TemplateFileList, alert.TemplateFileList...)

}
