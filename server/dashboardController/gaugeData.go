package dashboardController

import (
	//	"fmt"
	//	"resultra/datasheet/server/common/recordSortDataModel"
	//	"resultra/datasheet/server/dashboard"
	"resultra/datasheet/server/dashboard/components/gauge"
	//	"resultra/datasheet/server/dashboard/values"
	//	"resultra/datasheet/server/recordFilter"
	//	"resultra/datasheet/server/recordReadController"
)

type GaugeData struct {
	GaugeID               string                `json:"gaugeID"`
	Gauge                 gauge.Gauge           `json:"gauge"`
	Title                 string                `json:"title"`
	GroupedSummarizedVals GroupedSummarizedVals `json:"groupedSummarizedVals"`
}
