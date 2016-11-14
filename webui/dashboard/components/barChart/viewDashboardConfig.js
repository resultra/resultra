function barChartViewDashboardConfig(barChartRef) {
	
	var barChartElemPrefix = "barChart_"
	
	// Start with the defaultFilterIDs for the given bar chart. Then, if the selection changes,
	// the currentFilterIDs will also change. Then, if the bar chart is selected again, the 
	// current IDs can be used instead of the default IDs.
	var currentFilterIDs = barChartRef.properties.defaultFilterIDs
	
	function reloadBarChart() {
	
		// TODO - Include filtering parameters when getting data
		var getDataParams = {
			parentDashboardID:barChartRef.parentDashboardID,
			barChartID:barChartRef.barChartID
		}
		jsonAPIRequest("dashboardController/getBarChartData",getDataParams,function(updatedBarChartData) {
			console.log("Redrawing barchart after changing filter selection")
			drawBarChart(updatedBarChartData) // redraw the chart
		})
		
	}
	
	function selectBarChart(selectedBarChartRef) {
			console.log("Select bar chart: " + selectedBarChartRef.barChartID)
			// Toggle to the summary properties, hiding the other property panels
			
			var filterPaneParams = {
				elemPrefix: barChartElemPrefix,
				tableID: selectedBarChartRef.properties.dataSrcTableID,
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