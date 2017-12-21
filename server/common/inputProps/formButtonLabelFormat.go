package inputProps

const LabelFormatForm string = "form"

type FormButtonLabelFormatProperties struct {
	LabelType   string `json:"labelType"`
	CustomLabel string `json:"customLabel"`
}

func NewDefaultFormButtonLabelFormatProperties() FormButtonLabelFormatProperties {
	return FormButtonLabelFormatProperties{
		LabelType:   LabelFormatForm,
		CustomLabel: ""}
}
