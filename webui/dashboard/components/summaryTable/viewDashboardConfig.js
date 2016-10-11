function summaryTableViewDashboardConfig(summaryTableRef) {
	
	var summaryTableElemPrefix =  "summaryTable_"
	
	// Start with the defaultFilterIDs for the given bar chart. Then, if the selection changes,
	// the currentFilterIDs will also change. Then, if the bar chart is selected again, the 
	// current IDs can be used instead of the default IDs.
	var currentFilterIDs = summaryTableRef.properties.defaultFilterIDs
	
	
	function reloadSummaryTable() {
		currentFilterIDs = getCurrentFilterPanelFilterIDsWithDefaults(summaryTableElemPrefix, 
			currentFilterIDs,
			summaryTableRef.properties.availableFilterIDs)
	
		var getDataParams = {
			parentDashboardID:summaryTableRef.parentDashboardID,
			summaryTableID:summaryTableRef.summaryTableID,
			filterIDs: currentFilterIDs
		}
		jsonAPIRequest("dashboardController/getSummaryTableData",getDataParams,function(updatedSummaryTableData) {
			console.log("Repopulating summary table after changing filter selection")
			initSummaryTableData(summaryTableRef.parentDashboardID,updatedSummaryTableData)
		})
		
	}
	
	function selectSummaryTable(updatedSummaryTableRef) {
			console.log("Select summary table: " + updatedSummaryTableRef.summaryTableID)
	
	
			var filterPaneParams = {
				elemPrefix: summaryTableElemPrefix,
				tableID: updatedSummaryTableRef.properties.dataSrcTableID,
				defaultFilterIDs: currentFilterIDs,
				availableFilterIDs: updatedSummaryTableRef.properties.availableFilterIDs,
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