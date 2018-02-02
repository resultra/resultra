
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

