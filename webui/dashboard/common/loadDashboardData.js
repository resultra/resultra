function loadDashboardData(loadDashboardConfig)
{
		
	var dashboardID = loadDashboardConfig.dashboardContext.dashboardID
	
	var dashboardLayoutSelector = '#dashboardCanvas'
	
	function initBarChartLayout($componentRow,barChartData) {
		
		
		var barChartHTML = barChartContainerHTML();
		var $barChartElem = $(barChartHTML)

		setContainerComponentInfo($barChartElem,barChartData.barChart,barChartData.barChartID)
		
		$componentRow.append($barChartElem)
		setElemDimensions($barChartElem,barChartData.barChart.properties.geometry)
		
		initBarChartData(dashboardID,$barChartElem, barChartData);	
		
		loadDashboardConfig.initBarChartComponent($barChartElem,barChartData.barChart)	
	}

	function initSummaryTableLayout($componentRow,summaryTableData) {
		
		var summaryTableHTML = summaryTableComponentHTML(summaryTableData.summaryTableID);
		var $summaryTableElem = $(summaryTableHTML)
	
		setContainerComponentInfo($summaryTableElem,summaryTableData.summaryTable,summaryTableData.summaryTableID)
		
		$componentRow.append($summaryTableElem)
		setElemDimensions($summaryTableElem,summaryTableData.summaryTable.properties.geometry)
		
		initSummaryTableData(dashboardID,$summaryTableElem,summaryTableData)
		
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
		var $parentLayout = $(dashboardLayoutSelector)
		
		populateComponentLayout(dashboardLayout,$parentLayout,compenentIDComponentMap)
		
		loadDashboardConfig.doneLoadingDashboardDataFunc()
						
	}) // getData
	
}
