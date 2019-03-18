// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

var summaryTableDashboardDesignConfig = {
	
}

function selectDashboardSummaryTable($container,summaryTableRef) {
	console.log("Select summary table: " + summaryTableRef.summaryTableID)
	
	var summaryTablePropertiesArgs = {
		databaseID: designDashboardContext.databaseID,
		dashboardID: summaryTableRef.parentDashboardID,
		summaryTableID: summaryTableRef.summaryTableID,
		$summaryTable: $container
	}
	loadSummaryTableProperties(summaryTablePropertiesArgs)
}

function resizeDashboardSummaryTable($container,geometry) {
	
	var summaryTableRef =  getContainerObjectRef($container)
	
	var resizeParams = {
		parentDashboardID: summaryTableRef.parentDashboardID,
		summaryTableID: summaryTableRef.summaryTableID,
		geometry: geometry
	}

	console.log("Resize summary table: " +  JSON.stringify(resizeParams))

	jsonAPIRequest("dashboard/summaryTable/setDimensions", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.summaryTableID)
	})	
}

function initDesignDashboardSummaryTable() {
	
}

function populatePlaceholderSummaryTableData($summaryTable) {
}

var summaryTableDashboardDesignConfig = {
	draggableHTMLFunc:	summaryTableComponentHTML,
	populatePlaceholderData: populatePlaceholderSummaryTableData,
	createNewComponentAfterDropFunc: openNewSummaryTableDialog,
	resizeConstraints: elemResizeConstraints(100,1200,300,1200),
	resizeFunc: resizeDashboardSummaryTable,
	resizeHandles: 'e,s,se',
	initFunc: initDesignDashboardSummaryTable,
	selectionFunc: selectDashboardSummaryTable
	
}