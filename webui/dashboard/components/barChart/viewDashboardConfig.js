function barChartViewDashboardConfig(barChartRef) {
	
	var barChartElemPrefix = "barChart_"
	
	function reloadBarChart() {
	
		var currentFilterIDs = getCurrentFilterPanelFilterIDsWithDefaults(barChartElemPrefix, 
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
				defaultFilterIDs: selectedBarChartRef.properties.defaultFilterIDs,
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