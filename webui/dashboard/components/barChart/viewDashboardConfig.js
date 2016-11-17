function barChartViewDashboardConfig(barChartRef) {
	
	var barChartElemPrefix = "barChart_"
	
	// Start with the default filter rules for the given bar chart. Then, if the selection changes,
	// currFilterRules will also change. Then, if the bar chart is selected again, the 
	// currFilterRules can be used instead of the default filter rules.
	var currFilterRules = barChartRef.properties.defaultFilterRules
	
	function reloadBarChart() {
	
		// TODO - Include filtering parameters when getting data
		var getDataParams = {
			parentDashboardID:barChartRef.parentDashboardID,
			barChartID:barChartRef.barChartID,
			filterRules: currFilterRules
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
				defaultFilterRules: currFilterRules,
				initDone: function () {},
				updateFilterRules: function (updatedFilterRules) {
					// TODO - Reload table with updated filtering params.
					currFilterRules = updatedFilterRules
					reloadBarChart()
				}
			}

			initDefaultFilterRules(filterPaneParams)
		
			hideSiblingsShowOne('#barChartViewProps')
	}
	
	var viewConfig = {
		selectionFunc: selectBarChart
	}
	
	return viewConfig
}