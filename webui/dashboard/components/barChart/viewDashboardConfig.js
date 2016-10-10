function barChartViewDashboardConfig() {
	
	var barChartElemPrefix = "barChart_"
	
	function reloadBarChart() {
		
	}
	
	function selectBarChart(barChartRef) {
			console.log("Select bar chart: " + barChartRef.barChartID)
			// Toggle to the summary properties, hiding the other property panels
		
		
			var filterPaneParams = {
				elemPrefix: barChartElemPrefix,
				tableID: barChartRef.properties.dataSrcTableID,
				defaultFilterIDs: barChartRef.properties.defaultFilterIDs,
				availableFilterIDs: barChartRef.properties.availableFilterIDs,
				refilterCallbackFunc: reloadBarChart
			}

			initRecordFilterPanel(filterPaneParams)
		
			hideSiblingsShowOne('#barChartViewProps')
	}
	
	var viewConfig = {
		selectionFunc: selectBarChart
	}
	
	return viewConfig
}