
function selectDashboardBarChart ($container,barChartRef) {
	console.log("Select bar chart: " + barChartRef.barChartID)
	
	var barChartPropsArgs = {
		databaseID: designDashboardContext.databaseID,
		dashboardID: barChartRef.parentDashboardID,
		barChartID: barChartRef.barChartID,
		$barChart:$container,
		
		propertyUpdateComplete: function (updatedBarChartRef) {
			
			var updateContainer = $('#'+updatedBarChartRef.barChartID)
			updateContainer.data("barChartRef",updatedBarChartRef)
			
			var getDataParams = {
				parentDashboardID:updatedBarChartRef.parentDashboardID,
				barChartID:updatedBarChartRef.barChartID,
				$barChart:$container,
				filterRules: updatedBarChartRef.properties.defaultFilterRules
			}
			jsonAPIRequest("dashboardController/getBarChartData",getDataParams,function(updatedBarChartData) {
				console.log("Redrawing barchart after properties update")
				drawBarChart(updatedBarChartData) // redraw the chart
			})
		}
	}
	
	loadBarChartProperties(barChartPropsArgs)
	
}

function resizeDashboardBarChart($container,geometry) {
	
	var barChartRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentDashboardID: barChartRef.parentDashboardID,
		barChartID: barChartRef.barChartID,
		geometry: geometry
	}	
			
 	jsonAPIRequest("dashboard/barChart/setDimensions",resizeParams,function(updatedBarChartRef) {
 		console.log("bar chart dimensions updated")
 	})	

	console.log("Resize bar chart: " +  JSON.stringify(resizeParams))
	
}

function initDesignDashboardBarChart() {
	
}

function populatePlaceholderBarchartData($barChart) {
	drawDesignModeDummyBarChart($barChart);
}

var barChartDashboardDesignConfig = {
	draggableHTMLFunc:	barChartContainerHTML,
	populatePlaceholderData: populatePlaceholderBarchartData,
	createNewComponentAfterDropFunc: openNewBarChartDialog,
	resizeConstraints: elemResizeConstraints(300,1200,300,300),
	resizeFunc: resizeDashboardBarChart,
	initFunc: initDesignDashboardBarChart,
	selectionFunc: selectDashboardBarChart
	
}