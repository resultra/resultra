

function loadBarChartProperties(barChartPropsArgs) {

	var barChartContainer = $('#'+barChartPropsArgs.barChartID)
	var barChartRef = getContainerObjectRef(barChartPropsArgs.$barChart)
	var barChartElemPrefix = "barChart_"

	var filterPropertyPanelParams = {
		elemPrefix: barChartElemPrefix,
		databaseID: barChartPropsArgs.databaseID,
		defaultFilterRules: barChartRef.properties.defaultFilterRules,
		initDone: function () {},
		updateFilterRules: function (updatedFilterRules) {
			var setDefaultFiltersParams = {
				parentDashboardID:barChartPropsArgs.dashboardID,
				barChartID: barChartRef.barChartID,
				defaultFilterRules: updatedFilterRules
			}
			jsonAPIRequest("dashboard/barChart/setDefaultFilterRules",setDefaultFiltersParams,function(updatedBarChart) {
				console.log(" Default filters updated")
				setContainerComponentInfo(barChartPropsArgs.$barChart,updateBarChart,updatedBarChart.barChartID)
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
		databaseID: barChartPropsArgs.databaseID,
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
		databaseID: barChartPropsArgs.databaseID,
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
