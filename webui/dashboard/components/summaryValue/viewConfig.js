function summaryValViewDashboardConfig(summaryValRef) {
	
	var currFilterRules = summaryValRef.properties.defaultFilterRules
	
	
	function reloadSummaryVal($summaryTableContainer) {
			
		var getDataParams = {
			parentDashboardID:summaryValRef.parentDashboardID,
			summaryValID:summaryValRef.summaryValID,
			filterRules: currFilterRules
		}
		jsonAPIRequest("dashboardController/getSummaryValData",getDataParams,function(summaryValData) {
			initSummaryValData(summaryValRef.parentDashboardID,$summaryTableContainer, summaryValData)
		})		
		
	}
	
	
	function selectSummaryVal($container,selectedSummaryValRef) {
		
		var summaryValElemPrefix = "summaryVal_"
		
		console.log("Select summary value: " + selectedSummaryValRef.summaryValID)


		var filterPaneParams = {
			elemPrefix: summaryValElemPrefix,
			databaseID: viewDashboardContext.databaseID,
			defaultFilterRules: currFilterRules,
			initDone: function () {},
			updateFilterRules: function (updatedFilterRules) {
				currFilterRules = updatedFilterRules
				reloadSummaryVal($container)
			},
			refilterWithCurrentFilterRules: function() {
				reloadSummaryVal($container)
			}
		}

		initRecordFilterViewPanel(filterPaneParams)

		// Toggle to the summary properties, hiding the other property panels
		hideSiblingsShowOne('#summaryValueViewProps')
	}
	
	var viewConfig = {
		selectionFunc: selectSummaryVal
	}
	
	return viewConfig
}