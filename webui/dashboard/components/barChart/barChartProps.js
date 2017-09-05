

function loadBarChartProperties(barChartPropsArgs) {

	var barChartContainer = $('#'+barChartPropsArgs.barChartID)
	var barChartRef = getContainerObjectRef(barChartPropsArgs.$barChart)
	var barChartElemPrefix = "barChart_"
	
	
	function reloadBarChart(barChartRef) {
		var barChartDataParams = { 
			parentDashboardID: barChartRef.parentDashboardID,
			barChartID: barChartRef.barChartID,
			filterRules: barChartRef.properties.defaultFilterRules
		}
		jsonAPIRequest("dashboardController/getBarChartData",barChartDataParams,function(barChartData) {
			initBarChartData(barChartRef.parentDashboardID,barChartPropsArgs.$barChart, barChartData)
		})		
	}
	
	var preFilterBarCharElemPrefix = "barChartPreFilter_"
	var preFilterPropertyPanelParams = {
		elemPrefix: preFilterBarCharElemPrefix,
		databaseID: barChartPropsArgs.databaseID,
		defaultFilterRules: barChartRef.properties.preFilterRules,
		initDone: function () {},
		updateFilterRules: function (updatedFilterRules) {
			var setPreFiltersParams = {
				parentDashboardID:barChartPropsArgs.dashboardID,
				barChartID: barChartRef.barChartID,
				preFilterRules: updatedFilterRules
			}
			jsonAPIRequest("dashboard/barChart/setPreFilterRules",setPreFiltersParams,function(updatedBarChart) {
				console.log(" Default filters updated")
				reloadBarChart(updatedBarChart)
				setContainerComponentInfo(barChartPropsArgs.$barChart,updatedBarChart,updatedBarChart.barChartID)
			}) 
		}
	}
	initFilterPropertyPanel(preFilterPropertyPanelParams)
	
	

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
				reloadBarChart(updatedBarChart)
				setContainerComponentInfo(barChartPropsArgs.$barChart,updatedBarChart,updatedBarChart.barChartID)
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
				reloadBarChart(updatedBarChart)
				setContainerComponentInfo(barChartPropsArgs.$barChart,updatedBarChart,updatedBarChart.barChartID)
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
				reloadBarChart(updatedBarChart)
				setContainerComponentInfo(barChartPropsArgs.$barChart,updatedBarChart,updatedBarChart.barChartID)
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
				reloadBarChart(updatedBarChart)
				setContainerComponentInfo(barChartPropsArgs.$barChart,updatedBarChart,updatedBarChart.barChartID)
			})

		}
	}
	initDashboardValueSummaryPropertyPanel(yAxisPropertyPanelParams)

	var helpPopupParams = {
		initialMsg: barChartRef.properties.helpPopupMsg,
		elemPrefix: barChartElemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentDashboardID:barChartPropsArgs.dashboardID,
				barChartID: barChartRef.barChartID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("dashboard/barChart/setHelpPopupMsg",params,function(updatedBarChart) {
				setContainerComponentInfo(barChartPropsArgs.$barChart,updatedBarChart,updatedBarChart.barChartID)
				updateComponentHelpPopupMsg(barChartPropsArgs.$barChart, updatedBarChart)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)


	// Toggle to the bar chart properties, hiding the other property panels
	hideSiblingsShowOne('#barChartProps')

}
