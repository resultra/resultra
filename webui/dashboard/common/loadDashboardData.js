function loadDashboardData(dashboardID)
{
	
	var dashboardLayoutSelector = '#dashboardCanvas'
	
	function initBarChartLayout($componentRow,barChartData) {
		
		var barChartHTML = barChartContainerHTML(barChartData.barChartID);
		var barChartElem = $(barChartHTML)
		
		$componentRow.append(barChartElem)
		setElemDimensions(barChartElem,barChartData.barChart.properties.geometry)
		
		initBarChartData(dashboardID,barChartData);			
	}
	
	// Load the dashboard data
	var loadBarChartDataParams = { dashboardID: dashboardID }
	jsonAPIRequest("dashboard/getData",loadBarChartDataParams,function(dashboardData) {
		
		var compenentIDComponentMap = {}
			
		for (var barChartDataIndex in dashboardData.barChartsData) {
			var barChartData = dashboardData.barChartsData[barChartDataIndex]
			console.log ("Loading bar chart: id = " + barChartData.barChartID)
			compenentIDComponentMap[barChartData.barChartID] = {
				componentInfo: barChartData,
				initFunc: initBarChartLayout
			}			
			
		}
		
		function saveUpdatedDashboardComponentLayout(updatedLayout) {
			console.log("saveUpdatedDashboardComponentLayout: component layout = " + JSON.stringify(updatedLayout))		
			var setLayoutParams = {
				dashboardID: dashboardID,
				layout: updatedLayout
			}
			jsonAPIRequest("dashboard/setLayout", setLayoutParams, function(dashboardInfo) {})
			
		}
		
		var dashboardLayout = dashboardData.dashboard.properties.layout
		populateComponentLayout(dashboardLayout,dashboardLayoutSelector,
				compenentIDComponentMap,saveUpdatedDashboardComponentLayout)
		
		
		initObjectCanvasSelectionBehavior(dashboardLayoutSelector, function() {
			initDesignDashboardProperties(designDashboardContext.dashboardID)
			hideSiblingsShowOne('#dashboardProps')
		})
		
		
						
	})
	
}
