function summaryTableViewDashboardConfig(summaryTableRef) {
	
	var summaryTableElemPrefix =  "summaryTable_"
		
	function reloadSummaryTable() {

		// TODO - Include filtering parameters when loading table data
	
		var getDataParams = {
			parentDashboardID:summaryTableRef.parentDashboardID,
			summaryTableID:summaryTableRef.summaryTableID,
			filterRules: summaryTableRef.properties.defaultFilterRules
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
				defaultFilterRules: updatedSummaryTableRef.properties.defaultFilterRules,
				initDone: function () {},
				updateFilterRules: function (updatedFilterRules) {
					// TODO - Reload table with updated filtering params.
					reloadSummaryTable()
				}
			}

			initDefaultFilterRules(filterPaneParams)
	
			// Toggle to the summary properties, hiding the other property panels
			hideSiblingsShowOne('#summaryTableViewProps')
	}
	
	var viewConfig = {
		selectionFunc: selectSummaryTable
	}
	
	return viewConfig
}