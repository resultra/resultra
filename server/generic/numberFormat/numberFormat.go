package numberFormat

const NumberFormatGeneral string = "general"
const NumberFormatCurrency string = "currency"
const NumberFormatCurrency0Prec string = "currency0prec"
const NumberFormatPercent0 string = "percent0"
const NumberFormatPercent1 string = "percent1"
const NumberFormatPercent string = "percent"

type NumberFormatProperties struct {
	Format string `json:"format"`
}

func DefaultNumberFormatProperties() NumberFormatProperties {
	return NumberFormatProperties{Format: "general"}
}
