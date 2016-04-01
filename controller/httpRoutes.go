package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterAPIHTTPHandlers() {

	apiRouter := mux.NewRouter()

	apiRouter.HandleFunc("/api/newLayout", newLayout)
	apiRouter.HandleFunc("/api/newLayoutContainer", newLayoutContainer)
	apiRouter.HandleFunc("/api/resizeLayoutContainer", resizeLayoutContainer)
	apiRouter.HandleFunc("/api/getLayoutContainers", getLayoutContainers)

	apiRouter.HandleFunc("/api/newDashboard", newDashboard)
	apiRouter.HandleFunc("/api/newBarChart", newBarChart)
	apiRouter.HandleFunc("/api/getBarChartData", getBarChartData)
	apiRouter.HandleFunc("/api/updateBarChartProps", updateBarChartProps)

	apiRouter.HandleFunc("/api/newField", newField)
	apiRouter.HandleFunc("/api/newCalcField", newCalcField)
	apiRouter.HandleFunc("/api/getFieldsByType", getFieldsByType)
	apiRouter.HandleFunc("/api/getLayoutEditInfo", getLayoutEditInfo)

	apiRouter.HandleFunc("/api/newRecord", newRecord)
	apiRouter.HandleFunc("/api/setTextFieldValue", setTextFieldValue)
	apiRouter.HandleFunc("/api/setNumberFieldValue", setNumberFieldValue)
	apiRouter.HandleFunc("/api/getRecord", getRecord)
	apiRouter.HandleFunc("/api/getRecords", getRecords)
	apiRouter.HandleFunc("/api/getFilteredRecords", getFilteredRecords)

	apiRouter.HandleFunc("/api/newRecordFilterRule", newRecordFilterRule)
	apiRouter.HandleFunc("/api/getRecordFilterRules", getRecordFilterRules)

	apiRouter.HandleFunc("/api/validateCalcFieldEqn", validateCalcFieldEqn)

	http.Handle("/api/", apiRouter)
}
