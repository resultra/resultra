
function selectDashboardBarChart (barChartRef) {
	console.log("Select bar chart: " + barChartRef.barChartID)
	
	
	var barChartPropertiesArgs = {
		
	}
	loadBarChartProperties(barChartPropertiesArgs)
}

function resizeDashboardBarChart(barChartID,geometry) {
	var resizeParams = {
//		parentFormID: designFormContext.formID,
		barChartID: barChartID,
		geometry: geometry
	}

	console.log("Resize bar chart: " +  JSON.stringify(resizeParams))

	
//	jsonAPIRequest("frm/textBox/resize", resizeParams, function(updatedObjRef) {
//		setElemObjectRef(textBoxID,updatedObjRef)
//	})	
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