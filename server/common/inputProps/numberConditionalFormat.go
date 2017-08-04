package inputProps

type NumberConditionalFormat struct {
	Condition   string   `json:"condition"`
	ColorScheme string   `json:"colorScheme"`
	Param       *float64 `json:"param,omitempty"`
}
