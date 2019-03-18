package label

import (
	"resultra/tracker/webui/generic/propertiesSidebar"
)

type LabelPropertyTemplateParams struct {
	PanelParams       propertiesSidebar.PanelTemplateParams
	ElemPrefix        string
	HideNoLabelOption bool
}
