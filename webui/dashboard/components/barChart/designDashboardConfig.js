
function selectDashboardBarChart (barChartRef) {
	console.log("Select bar chart: " + barChartRef.barChartID)
	
	var barChartPropsArgs = {
		dashboardID: barChartRef.parentDashboardID,
		barChartID: barChartRef.barChartID,
		
		propertyUpdateComplete: function (updatedBarChartRef) {
			
			var updateContainer = $('#'+updatedBarChartRef.barChartID)
			updateContainer.data("barChartRef",updatedBarChartRef)
			
			var getDataParams = {
				parentDashboardID:updatedBarChartRef.parentDashboardID,
				barChartID:updatedBarChartRef.barChartID
			}
			jsonAPIRequest("dashboardController/getBarChartData",getDataParams,function(updatedBarChartData) {
				console.log("Redrawing barchart after properties update")
				drawBarChart(updatedBarChartData) // redraw the chart
			})
		}
	}
	
	loadBarChartProperties(barChartPropsArgs)
	
}

function resizeDashboardBarChart(barChartID,geometry) {
	
	var barChartRef = getElemObjectRef(barChartID)
	
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

function populatePlaceholderBarchartData(placeholderID) {
	drawDesignModeDummyBarChart(placeholderID);
}

var barChartDashboardDesignConfig = {
	draggableHTMLFunc:	barChartContainerHTML,
	populatePlaceholderData: populatePlaceholderBarchartData,
	createNewComponentAfterDropFunc: openNewBarChartDialog,
	resizeConstraints: elemResizeConstraints(100,600,400,400),
	resizeFunc: resizeDashboardBarChart,
	initFunc: initDesignDashboardBarChart,
	selectionFunc: selectDashboardBarChart
	
}