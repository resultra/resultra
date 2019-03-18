// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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