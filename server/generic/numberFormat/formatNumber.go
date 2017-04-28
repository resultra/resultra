package numberFormat

import (
	"fmt"
)

func FormatNumber(val float64, format string) string {
	switch format {
	case NumberFormatGeneral:
		return fmt.Sprintf("%v", val)
	case NumberFormatCurrency:
		if val >= 0.0 {
			return fmt.Sprintf("$%.02f", val)
		} else {
			// Put the negative sign on the LHS of $
			positiveVal := val * -1.0
			return fmt.Sprintf("-$%.02f", positiveVal)
		}
	case NumberFormatCurrency0Prec:
		if val >= 0.0 {
			return fmt.Sprintf("$%.0f", val)
		} else {
			// Put the negative sign on the LHS of $
			positiveVal := val * -1.0
			return fmt.Sprintf("-$%.0f", positiveVal)
		}
	case NumberFormatPercent0:
		percVal := val * 100.0
		return fmt.Sprintf("%.0f%%", percVal)
	case NumberFormatPercent1:
		percVal := val * 100.0
		return fmt.Sprintf("%.01f%%", percVal)
	case NumberFormatPercent:
		percVal := val * 100.0
		return fmt.Sprintf("%.02f%%", percVal)
	default:
		return fmt.Sprintf("%v", val)
	}
}
