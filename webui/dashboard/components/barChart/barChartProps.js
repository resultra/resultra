

function loadBarChartProperties(barChartPropsArgs) {
	
	var barChartContainer = $('#'+barChartPropsArgs.barChartID)
	var barChartRef = barChartContainer.data("barChartRef")
	
	
	var filterPropertyPanelParams = {
		elemPrefix: "barChart_",
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
		
	// Toggle to the bar chart properties, hiding the other property panels
	hideSiblingsShowOne('#barChartProps')
			
}
