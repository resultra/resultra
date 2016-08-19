function loadDashboardData(dashboardID)
{
	// Load the dashboard data
	var loadBarChartDataParams = { dashboardID: dashboardID }
	jsonAPIRequest("dashboard/getData",loadBarChartDataParams,function(dashboardData) {
		
		for (var barChartDataIndex in dashboardData.barChartsData) {
			var barChartData = dashboardData.barChartsData[barChartDataIndex]
			console.log ("Loading bar chart: id = " + barChartData.barChartID)
			
			var barChartHTML = barChartContainerHTML(barChartData.barChartID);
			var barChartElem = $(barChartHTML)
			
			$("#dashboardCanvas").append(barChartElem)
			setElemGeometry(barChartElem,barChartData.barChart.properties.geometry)
			
			initBarChartData(dashboardID,barChartData);			
			
		}
		
		initObjectCanvasSelectionBehavior('#dashboardCanvas', function() {
			initDesignDashboardProperties(designDashboardContext.dashboardID)
			hideSiblingsShowOne('#dashboardProps')
		})
		
						
	})
	
}
