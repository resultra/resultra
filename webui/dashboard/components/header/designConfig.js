// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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