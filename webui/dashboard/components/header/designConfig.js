function selectDashboardHeader ($container,headerRef) {
	
	loadDashboardHeaderProperties($container,headerRef)
	
}

function resizeDashboardHeader($container,geometry) {
	
	var headerRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentDashboardID: headerRef.parentDashboardID,
		headerID: headerRef.headerID,
		geometry: geometry
	}	
			
 	jsonAPIRequest("dashboard/header/setDimensions",resizeParams,function(updatedHeader) {
 		console.log("header dimensions updated")
 	})	
	
}

function initDesignDashboardHeader() {
	
}

var headerDashboardDesignConfig = {
	draggableHTMLFunc:	dashboardHeaderContainerHTML,
	populatePlaceholderData: function($header) {},
	createNewComponentAfterDropFunc: createNewDashboardHeader,
	resizeConstraints: elemResizeConstraints(80,720,50,50),
	resizeFunc: resizeDashboardHeader,
	initFunc: initDesignDashboardHeader,
	selectionFunc: selectDashboardHeader
	
}