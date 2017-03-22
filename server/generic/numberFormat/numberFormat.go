package numberFormat

type NumberFormatProperties struct {
	Format string `json:"format"`
}

func DefaultNumberFormatProperties() NumberFormatProperties {
	return NumberFormatProperties{Format: "general"}
}
