function selectDashboardGauge ($container,gaugeRef) {
	
	var gaugePropsArgs = {		
		databaseID: designDashboardContext.databaseID,
		dashboardID: gaugeRef.parentDashboardID,
		gaugeID: gaugeRef.gaugeID,
		$gauge:$container
	}
	
	loadDashboardGaugeProperties(gaugePropsArgs)
	
}

function resizeDashboardGauge($container,geometry) {
	
	var gaugeRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentDashboardID: gaugeRef.parentDashboardID,
		gaugeID: gaugeRef.gaugeID,
		geometry: geometry
	}	
			
 	jsonAPIRequest("dashboard/gauge/setDimensions",resizeParams,function(updatedGauge) {
 		console.log("gauge dimensions updated")
 	})	
	
}

function initDesignDashboardGauge() {
	
}

var gaugeDashboardDesignConfig = {
	draggableHTMLFunc:	dashboardGaugeContainerHTML,
	populatePlaceholderData: function($gauge) {},
	createNewComponentAfterDropFunc: openNewDashboardGaugeDialog,
	resizeConstraints: elemResizeConstraints(80,720,50,50),
	resizeFunc: resizeDashboardGauge,
	initFunc: initDesignDashboardGauge,
	selectionFunc: selectDashboardGauge
	
}