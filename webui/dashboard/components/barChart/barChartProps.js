

function loadBarChartProperties(barChartPropsArgs) {
	
	var barChartContainer = $('#'+barChartPropsArgs.barChartID)
	var barChartRef = barChartContainer.data("barChartRef")
	var barChartElemPrefix = "barChart_"
	
	var filterPropertyPanelParams = {
		elemPrefix: barChartElemPrefix,
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
		
	// Toggle to the bar chart properties, hiding the other property panels
	hideSiblingsShowOne('#barChartProps')
			
}
