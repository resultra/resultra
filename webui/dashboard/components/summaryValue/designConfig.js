function selectDashboardSummaryVal($container,summaryValRef) {
	
	var propsArgs = {		
		databaseID: designDashboardContext.databaseID,
		dashboardID: summaryValRef.parentDashboardID,
		summaryValID: summaryValRef.summaryValID,
		$summaryVal:$container
	}
	
	loadDashboardSummaryValProperties(propsArgs)
	
}

function resizeDashboardSummaryVal($container,geometry) {
	
	var summaryValRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentDashboardID: summaryValRef.parentDashboardID,
		summaryValID: summaryValRef.summaryValID,
		geometry: geometry
	}	
			
 	jsonAPIRequest("dashboard/summaryVal/setDimensions",resizeParams,function(updatedSummaryVal) {
 		console.log("summary value dimensions updated")
 	})	
	
}

function initDesignDashboardSummaryVal() {
	
}

var summaryValDashboardDesignConfig = {
	draggableHTMLFunc:	dashboardSummaryValContainerHTML,
	populatePlaceholderData: function($summaryVal) {},
	createNewComponentAfterDropFunc: openNewDashboardSummaryValDialog,
	resizeConstraints: elemResizeConstraints(80,720,50,50),
	resizeFunc: resizeDashboardSummaryVal,
	initFunc: initDesignDashboardSummaryVal,
	selectionFunc: selectDashboardSummaryVal
	
}