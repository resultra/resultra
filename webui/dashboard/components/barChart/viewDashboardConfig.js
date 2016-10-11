function barChartViewDashboardConfig(barChartRef) {
	
	var barChartElemPrefix = "barChart_"
	
	// Start with the defaultFilterIDs for the given bar chart. Then, if the selection changes,
	// the currentFilterIDs will also change. Then, if the bar chart is selected again, the 
	// current IDs can be used instead of the default IDs.
	var currentFilterIDs = barChartRef.properties.defaultFilterIDs
	
	function reloadBarChart() {
	
		currentFilterIDs = getCurrentFilterPanelFilterIDsWithDefaults(barChartElemPrefix, 
			barChartRef.properties.defaultFilterIDs,
			barChartRef.properties.availableFilterIDs)
	
		var getDataParams = {
			parentDashboardID:barChartRef.parentDashboardID,
			barChartID:barChartRef.barChartID,
			filterIDs: currentFilterIDs
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
				defaultFilterIDs: currentFilterIDs,
				availableFilterIDs: selectedBarChartRef.properties.availableFilterIDs,
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