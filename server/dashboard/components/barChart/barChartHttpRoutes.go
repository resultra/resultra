package barChart

import (
	"github.com/gorilla/mux"
)

func RegisterHTTPHandlers(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/api/newBarChart", newBarChart)
	apiRouter.HandleFunc("/api/getBarChartData", getBarChartData)
	apiRouter.HandleFunc("/api/setBarChartTitle", setBarChartTitle)
	apiRouter.HandleFunc("/api/setBarChartDimensions", setBarChartDimensions)

}
