function initGaugeViewProperties(componentSelectionParams) {
	console.log("Init gauge properties panel")
		
	var gaugeRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()	

	var elemPrefix = "gauge_"		
	
	// Toggle to the gauges properties, hiding the other property panels
	hideSiblingsShowOne('#gaugeViewProps')
	
	
}