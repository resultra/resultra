function loadDashboardData(dashboardID)
{
	// Load the dashboard data
	var loadBarChartDataParams = { dashboardID: dashboardID }
	jsonAPIRequest("dashboard/getData",loadBarChartDataParams,function(dashboardData) {
		
		/* TODO - Re-enable dashboard loading once the new grid system is in place 
		for (var barChartDataIndex in dashboardData.barChartsData) {
			var barChartData = dashboardData.barChartsData[barChartDataIndex]
			console.log ("Loading bar chart: id = " + barChartData.barChartID)
			
			var barChartHTML = barChartContainerHTML(barChartData.barChartID);
			var barChartElem = $(barChartHTML)
			
			$("#dashboardCanvas").append(barChartElem)
			setElemDimensions(barChartElem,barChartData.barChart.properties.geometry)
			
			initBarChartData(dashboardID,barChartData);			
			
		}
		
		initObjectCanvasSelectionBehavior('#dashboardCanvas', function() {
			initDesignDashboardProperties(designDashboardContext.dashboardID)
			hideSiblingsShowOne('#dashboardProps')
		})
		
		*/
		
						
	})
	
}
