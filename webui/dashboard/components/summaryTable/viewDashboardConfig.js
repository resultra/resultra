function summaryTableViewDashboardConfig() {
	
	function selectSummaryTable(summaryTableRef) {
			console.log("Select summary table: " + summaryTableRef.summaryTableID)
			// Toggle to the summary properties, hiding the other property panels
			hideSiblingsShowOne('#summaryTableViewProps')
	}
	
	var viewConfig = {
		selectionFunc: selectSummaryTable
	}
	
	return viewConfig
}