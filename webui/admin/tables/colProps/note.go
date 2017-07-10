package colProps

import (
	"resultra/datasheet/webui/common/form/components/common/label"
	"resultra/datasheet/webui/common/form/components/common/permissions"
	"resultra/datasheet/webui/generic/propertiesSidebar"
)

type NoteColPropsTemplateParams struct {
	ElemPrefix            string
	LabelPanelParams      label.LabelPropertyTemplateParams
	PermissionPanelParams permissions.PermissionsPropertyTemplateParams
}

func newNoteTemplateParams() NoteColPropsTemplateParams {

	elemPrefix := "note_"

	templParams := NoteColPropsTemplateParams{
		ElemPrefix: elemPrefix,
		LabelPanelParams: label.LabelPropertyTemplateParams{ElemPrefix: elemPrefix,
			PanelParams: propertiesSidebar.PanelTemplateParams{PanelHeaderLabel: "Label", PanelID: "noteLabel"}},
		PermissionPanelParams: permissions.NewPermissionTemplateParams(elemPrefix, "notePerms")}

	return templParams

}
