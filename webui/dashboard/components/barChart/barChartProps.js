

function loadBarChartProperties(barChartPropsArgs) {
	
	var barChartContainer = $('#'+barChartPropsArgs.barChartID)
	var barChartRef = barChartContainer.data("barChartRef")
	var barChartElemPrefix = "barChart_"
	
	var filterPropertyPanelParams = {
		elemPrefix: barChartElemPrefix,
		tableID: barChartRef.properties.dataSrcTableID,
		defaultFilterIDs: barChartRef.properties.defaultFilterIDs,
		setDefaultFilterFunc: function(defaultFilterIDs) {
			var params = {
				barChartID: barChartRef.barChartID,
				parentDashboardID: barChartPropsArgs.dashboardID,
				defaultFilterIDs: defaultFilterIDs }
			jsonAPIRequest("dashboard/barChart/setDefaultFilters",params,function(updatedBarChart) {
				barChartContainer.data("barChartRef",updatedBarChart)
				console.log("Default filters updated")
			}) // set record's number field value
			
		},
		availableFilterIDs: barChartRef.properties.availableFilterIDs,
		setAvailableFilterFunc: function(availFilterIDs) {
			var params = {
				barChartID: barChartRef.barChartID,
				parentDashboardID: barChartPropsArgs.dashboardID,
				availableFilterIDs: availFilterIDs }
			jsonAPIRequest("dashboard/barChart/setAvailableFilters",params,function(updatedBarChart) {
				barChartContainer.data("barChartRef",updatedBarChart)
				console.log("Available filters updated")
			}) // set record's number field value
			
		}
	}
	initFilterPropertyPanel(filterPropertyPanelParams)
	
	var titlePropertyPanelParams = {
		dashboardID: barChartPropsArgs.dashboardID,
		title: barChartRef.properties.title,
		setTitleFunc: function(newTitle) {
			
			var setTitleParams = {
				parentDashboardID:barChartPropsArgs.dashboardID,
				barChartID: barChartRef.barChartID,
				newTitle:newTitle
			}
			jsonAPIRequest("dashboard/barChart/setTitle",setTitleParams,function(updatedBarChart) {
					barChartContainer.data("barChartRef",updatedBarChart)
			})
			
		}
	}
	initDashboardComponentTitlePropertyPanel(barChartElemPrefix,titlePropertyPanelParams)
	
	
	var xAxisPropertyPanelParams = {
		elemPrefix: barChartElemPrefix,
		tableID: barChartRef.properties.dataSrcTableID,
		valGroupingProps: barChartRef.properties.xAxisVals,
		saveValueGroupingFunc: function(newValueGroupingParams) {
			var setXAxisValGroupingParams = {
				parentDashboardID:barChartPropsArgs.dashboardID,
				barChartID: barChartRef.barChartID,
				xAxisValueGrouping:newValueGroupingParams
			}
			jsonAPIRequest("dashboard/barChart/setXAxisValueGrouping",setXAxisValGroupingParams,function(updatedBarChart) {
					barChartContainer.data("barChartRef",updatedBarChart)
			})
		}
		
	}
	initDashboardValueGroupingPropertyPanel(xAxisPropertyPanelParams)
	
	var yAxisPropertyPanelParams = {
		elemPrefix: barChartElemPrefix,
		tableID: barChartRef.properties.dataSrcTableID,
		valSummaryProps: barChartRef.properties.yAxisValSummary,
		saveValueSummaryFunc: function(newValSummaryParams) {
			var setYAxisSummaryParams = {
				parentDashboardID:barChartPropsArgs.dashboardID,
				barChartID: barChartRef.barChartID,
				yAxisValSummary:newValSummaryParams
			}
			jsonAPIRequest("dashboard/barChart/setYAxisSummaryVals",
								setYAxisSummaryParams,function(updatedBarChart) {
				barChartContainer.data("barChartRef",updatedBarChart)
			})
			
		}
	}
	initDashboardValueSummaryPropertyPanel(yAxisPropertyPanelParams)
	
	// Toggle to the bar chart properties, hiding the other property panels
	hideSiblingsShowOne('#barChartProps')
			
}
