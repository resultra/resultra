function summaryTableViewDashboardConfig(summaryTableRef) {
	
	var summaryTableElemPrefix =  "summaryTable_"

	// Store the default filter rules in a local variable (closure). If the
	// filter rules are modified locally in the dashboard view, the rules can
	// be modified there. 
	var currFilterRules = summaryTableRef.properties.defaultFilterRules
		
	function reloadSummaryTable($summaryTableContainer) {
	
		var getDataParams = {
			parentDashboardID:summaryTableRef.parentDashboardID,
			summaryTableID:summaryTableRef.summaryTableID,
			filterRules: currFilterRules
		}
		jsonAPIRequest("dashboardController/getSummaryTableData",getDataParams,function(updatedSummaryTableData) {
			console.log("Repopulating summary table after changing filter selection")
			initSummaryTableData(summaryTableRef.parentDashboardID,$summaryTableContainer,updatedSummaryTableData)
		})
		
	}
	
	function selectSummaryTable($container,updatedSummaryTableRef) {
			console.log("Select summary table: " + updatedSummaryTableRef.summaryTableID)
	
	
			var filterPaneParams = {
				elemPrefix: summaryTableElemPrefix,
				databaseID: viewDashboardContext.databaseID,
				defaultFilterRules: summaryTableRef.properties.defaultFilterRules,
				initDone: function () {},
				updateFilterRules: function (updatedFilterRules) {
					currFilterRules = updatedFilterRules
					reloadSummaryTable($container)
				},
				refilterWithCurrentFilterRules: function() {
					reloadSummaryTable($container)
				}
			}

			initRecordFilterViewPanel(filterPaneParams)
	
			// Toggle to the summary properties, hiding the other property panels
			hideSiblingsShowOne('#summaryTableViewProps')
	}
	
	var viewConfig = {
		selectionFunc: selectSummaryTable
	}
	
	return viewConfig
}