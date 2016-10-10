function loadDashboardData(loadDashboardConfig)
{
		
	var dashboardID = loadDashboardConfig.dashboardContext.dashboardID
	
	var dashboardLayoutSelector = '#dashboardCanvas'
	
	function initBarChartLayout($componentRow,barChartData) {
		
		var barChartHTML = barChartContainerHTML(barChartData.barChartID);
		var $barChartElem = $(barChartHTML)
		
		$componentRow.append($barChartElem)
		setElemDimensions($barChartElem,barChartData.barChart.properties.geometry)
		
		initBarChartData(dashboardID,barChartData);	
		
		loadDashboardConfig.initBarChartComponent($barChartElem,barChartData.barChart)	
	}

	function initSummaryTableLayout($componentRow,summaryTableData) {
		
		var summaryTableHTML = summaryTableComponentHTML(summaryTableData.summaryTableID);
		var $summaryTableElem = $(summaryTableHTML)
		
		$componentRow.append($summaryTableElem)
		setElemDimensions($summaryTableElem,summaryTableData.summaryTable.properties.geometry)
		
		initSummaryTableData(dashboardID,summaryTableData)
		
		loadDashboardConfig.initSummaryTableComponent($summaryTableElem,summaryTableData.summaryTable)
	}

	
	// Load the dashboard data
	var loadBarChartDataParams = { dashboardID: dashboardID }
	jsonAPIRequest("dashboardController/getDefaultData",loadBarChartDataParams,function(dashboardData) {
		
		var compenentIDComponentMap = {}
			
		for (var barChartDataIndex in dashboardData.barChartsData) {
			var barChartData = dashboardData.barChartsData[barChartDataIndex]
			console.log ("Loading bar chart: id = " + barChartData.barChartID)
			compenentIDComponentMap[barChartData.barChartID] = {
				componentInfo: barChartData,
				initFunc: initBarChartLayout
			}		
		}
		for (var summaryTableDataIndex in dashboardData.summaryTablesData) {
			var summaryTableData = dashboardData.summaryTablesData[summaryTableDataIndex]
			console.log ("Loading summary table: id = " + summaryTableData.summaryTableID)
			compenentIDComponentMap[summaryTableData.summaryTableID] = {
				componentInfo: summaryTableData,
				initFunc: initSummaryTableLayout
			}		
		}
		
		
		var dashboardLayout = dashboardData.dashboard.properties.layout
		populateComponentLayout(dashboardLayout,dashboardLayoutSelector,compenentIDComponentMap)
		
		loadDashboardConfig.doneLoadingDashboardDataFunc()
						
	}) // getData
	
}
