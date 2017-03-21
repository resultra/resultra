
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
	resizeConstraints: elemResizeConstraints(100,600,400,400),
	resizeFunc: resizeDashboardSummaryTable,
	initFunc: initDesignDashboardSummaryTable,
	selectionFunc: selectDashboardSummaryTable
	
}