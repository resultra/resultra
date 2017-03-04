
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

function resizeDashboardSummaryTable($container,geometry) {
	
	var summaryTableRef =  getContainerObjectRef($container)
	
	var resizeParams = {
		parentDashboardID: summaryTableRef.parentDashboardID,
		summaryTableID: summaryTableID,
		geometry: geometry
	}

	console.log("Resize summary table: " +  JSON.stringify(resizeParams))

	jsonAPIRequest("dashboard/summaryTable/setDimensions", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.summaryTableID)
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