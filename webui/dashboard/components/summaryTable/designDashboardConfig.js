
var summaryTableDashboardDesignConfig = {
	
}

function selectDashboardSummaryTable(summaryTableRef) {
	console.log("Select summary table: " + summaryTableRef.summaryTableID)
	
	var summaryTablePropertiesArgs = {
		databaseID: designDashboardContext.databaseID,
		dashboardID: summaryTableRef.parentDashboardID,
		summaryTableID: summaryTableRef.summaryTableID
	}
	loadSummaryTableProperties(summaryTablePropertiesArgs)
}

function resizeDashboardSummaryTable(summaryTableID,geometry) {
	
	var summaryTableRef = getElemObjectRef(summaryTableID)
	
	var resizeParams = {
		parentDashboardID: summaryTableRef.parentDashboardID,
		summaryTableID: summaryTableID,
		geometry: geometry
	}

	console.log("Resize summary table: " +  JSON.stringify(resizeParams))

	jsonAPIRequest("dashboard/summaryTable/setDimensions", resizeParams, function(updatedObjRef) {
		setElemObjectRef(updatedObjRef.summaryTableID,updatedObjRef)
	})	
}

function initDesignDashboardSummaryTable() {
	
}

function populatePlaceholderSummaryTableData(placeholderID) {
}

var summaryTableDashboardDesignConfig = {
	draggableHTMLFunc:	summaryTableComponentHTML,
	populatePlaceholderData: populatePlaceholderSummaryTableData,
	createNewComponentAfterDropFunc: openNewSummaryTableDialog,
	resizeConstraints: elemResizeConstraints(100,600,400,400),
	resizeFunc: resizeDashboardSummaryTable,
	initFunc: initDesignDashboardSummaryTable,
	selectionFunc: selectDashboardSummaryTable
	
}