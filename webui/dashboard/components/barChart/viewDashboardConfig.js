function barChartViewDashboardConfig() {
	
	function selectBarChart(barChartRef) {
			console.log("Select bar chart: " + barChartRef.barChartID)
			// Toggle to the summary properties, hiding the other property panels
			hideSiblingsShowOne('#barChartViewProps')
	}
	
	var viewConfig = {
		selectionFunc: selectBarChart
	}
	
	return viewConfig
}