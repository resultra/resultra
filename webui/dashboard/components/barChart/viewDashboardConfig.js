// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function barChartViewDashboardConfig(barChartRef) {
	
	var barChartElemPrefix = "barChart_"
	
	// Start with the default filter rules for the given bar chart. Then, if the selection changes,
	// currFilterRules will also change. Then, if the bar chart is selected again, the 
	// currFilterRules can be used instead of the default filter rules.
	var currFilterRules = barChartRef.properties.defaultFilterRules
	
	function reloadBarChart($container) {
	
		var getDataParams = {
			parentDashboardID:barChartRef.parentDashboardID,
			barChartID:barChartRef.barChartID,
			filterRules: currFilterRules
		}
		jsonAPIRequest("dashboardController/getBarChartData",getDataParams,function(updatedBarChartData) {
			console.log("Redrawing barchart after changing filter selection")
			drawBarChart($container,updatedBarChartData) // redraw the chart
		})
		
	}
	
	function selectBarChart($container,selectedBarChartRef) {
			console.log("Select bar chart: " + selectedBarChartRef.barChartID)
			// Toggle to the summary properties, hiding the other property panels
			
			var filterPaneParams = {
				elemPrefix: barChartElemPrefix,
				databaseID: viewDashboardContext.databaseID,
				defaultFilterRules: currFilterRules,
				initDone: function () {},
				updateFilterRules: function (updatedFilterRules) {
					currFilterRules = updatedFilterRules
					reloadBarChart($container)
				},
				refilterWithCurrentFilterRules: function() {
					reloadBarChart($container)
				}
			}

			initRecordFilterViewPanel(filterPaneParams)
		
			hideSiblingsShowOne('#barChartViewProps')
	}
	
	var viewConfig = {
		selectionFunc: selectBarChart
	}
	
	return viewConfig
}