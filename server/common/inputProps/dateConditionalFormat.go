package inputProps

import "time"

type DateConditionalFormat struct {
	Condition   string     `json:"condition"`
	ColorScheme string     `json:"colorScheme"`
	NumberParam *float64   `json:"numberParam,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
}
