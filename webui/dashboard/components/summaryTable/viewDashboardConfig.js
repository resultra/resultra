function summaryTableViewDashboardConfig(summaryTableRef) {
	
	var summaryTableElemPrefix =  "summaryTable_"
	
	function reloadSummaryTable() {
		var currentFilterIDs = getCurrentFilterPanelFilterIDsWithDefaults(summaryTableElemPrefix, 
			summaryTableRef.properties.defaultFilterIDs,
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
				defaultFilterIDs: updatedSummaryTableRef.properties.defaultFilterIDs,
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