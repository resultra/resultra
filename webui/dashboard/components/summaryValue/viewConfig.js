// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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