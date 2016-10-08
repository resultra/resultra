function summaryTableViewDashboardConfig() {
	
	var summaryTableElemPrefix =  "summaryTable_"
	
	function reloadSummaryTable() {
		
	}
	
	function selectSummaryTable(summaryTableRef) {
			console.log("Select summary table: " + summaryTableRef.summaryTableID)
	
	
			var filterPaneParams = {
				elemPrefix: summaryTableElemPrefix,
				tableID: summaryTableRef.properties.dataSrcTableID,
				defaultFilterIDs: summaryTableRef.properties.defaultFilterIDs,
				availableFilterIDs: summaryTableRef.properties.availableFilterIDs,
				refilterCallbackFunc: reloadSummaryTable
			}

			initRecordFilterPanel(filterPaneParams)
	
			// Toggle to the summary properties, hiding the other property panels
			hideSiblingsShowOne('#summaryTableViewProps')
	}
	
	var viewConfig = {
		selectionFunc: selectSummaryTable
	}
	
	return viewConfig
}