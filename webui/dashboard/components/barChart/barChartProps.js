

function loadBarChartProperties(barChartPropsArgs) {
	
	var barChartContainer = $('#'+barChartPropsArgs.barChartID)
	var barChartRef = barChartContainer.data("barChartRef")
		
	// Toggle to the bar chart properties, hiding the other property panels
	hideSiblingsShowOne('#barChartProps')
			
}
