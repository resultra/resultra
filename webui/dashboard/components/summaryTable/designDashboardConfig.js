
var summaryTableDashboardDesignConfig = {
	
}

function selectDashboardSummaryTable(summaryTableRef) {
	console.log("Select summary table: " + summaryTableRef.summaryTableID)
	
	
	var summaryTablePropertiesArgs = {
		summaryTableID: summaryTableRef.summaryTableID
	}
	loadSummaryTableProperties(summaryTablePropertiesArgs)
}

function resizeDashboardSummaryTable(summaryTableID,geometry) {
	var resizeParams = {
//		parentFormID: designFormContext.formID,
		barChartID: barChartID,
		geometry: geometry
	}

	console.log("Resize summary table: " +  JSON.stringify(resizeParams))

	
//	jsonAPIRequest("frm/textBox/resize", resizeParams, function(updatedObjRef) {
//		setElemObjectRef(textBoxID,updatedObjRef)
//	})	
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