// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

var dashboardComponentValueGroupingPanelID = "dashboardComponentValueGrouping"

function createNewDashboardComponentValueGroupingPanelConfig(elemPrefix,databaseID) {
	
	var panelSelector = "#" + elemPrefix + "DashboardComponentValueGroupingPanel"
	
		// Setup the same property panel used for editing the grouping val from the side-bar
		// of the dashboard property page.
	var groupingPropertyPanelParams = {
		elemPrefix: elemPrefix,
		databaseID: databaseID,
		valGroupingProps: null, // used to initialize the properties -- null since these are new properties
		saveValueGroupingFunc: function(newValueGroupingParams) { 
			// No-op since this is for a new grouping 
		}
	}
	var panelFormInputs = new initDashboardValueGroupingPropertyPanel(groupingPropertyPanelParams)
		
	function validateValueGroupingForm() {
		
		var rowGrouping = panelFormInputs.getRowGrouping()
		if (rowGrouping !== null) {
			return true
		} else {
			return false
		}
		
	}
		
	function getPanelValues() {
		
		var rowGrouping = panelFormInputs.getRowGrouping()		
		return rowGrouping
	}
	
	var dashboardComponentValueGroupingPanelConfig = {
		divID: panelSelector,
		panelID: dashboardComponentValueGroupingPanelID,
		progressPerc:40,
		getPanelVals:getPanelValues,
		initPanel: function($dialog) {
					
			var nextButtonSelector = '#' + elemPrefix + 'NewDashboardComponentValueGroupingNextButton'
			initButtonClickHandler(nextButtonSelector,function() {
				if(validateValueGroupingForm()) {				
					transitionToNextWizardDlgPanelByID($dialog,dashboardComponentValueSummaryPanelID)
				} // if validate panel's form
			})

		}, // init panel
		transitionIntoPanel: function ($dialog) { 
			
			setWizardDialogButtonSet("newDashboardComponentValueGroupingButtons")									
				
		} // transitionIntoPanel
		
	}
	
	return dashboardComponentValueGroupingPanelConfig
	
} // createNewDashboardComponentValueGroupingPanelConfig

