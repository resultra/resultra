var dashboardComponentSelectTablePanelID = "dashboardComponentSelectTable"

function createNewDashboardComponentSelectTablePanelConfig(elemPrefix) {
	
	var panelSelector = "#" + elemPrefix + "DashboardSelectTablePanel"
	var tableSelection = createPrefixedTemplElemInfo(elemPrefix,"DashboardTableSelection")
	var tableSelectionGroup = createPrefixedTemplElemInfo(elemPrefix,"DashboardTableSelectionGroup")
	
	function validateTableSelectionForm() {
		if(!validateNonEmptyFormField(tableSelection.selector)) {
			return false
		}
		return true
	}
	
	function populateTableSelectionMenu(tableRefList, menuSelector) {
		$(menuSelector).empty()		
		$(menuSelector).append(defaultSelectOptionPromptHTML("Select a Table"))
		$.each(tableRefList, function(index, tableRef) {
			$(menuSelector).append(selectFieldHTML(tableRef.tableID, tableRef.name))		
		})
	}
	
	function getPanelValues() {
		var selectedTableID = $(tableSelection.selector).val()
		console.log("Selected table: " + selectedTableID)
		return selectedTableID
	}
	
	
	var dashboardSelectTablePanelConfig = {
		divID: panelSelector,
		panelID: dashboardComponentSelectTablePanelID,
		progressPerc:20,
		getPanelVals:getPanelValues,
		initPanel: function($dialog) {
			var tableListParams =  { databaseID: databaseID }
			jsonAPIRequest("table/getList",tableListParams,function(tableRefs) {
				populateTableSelectionMenu(tableRefs,tableSelection.selector)
			})
			
			revalidateNonEmptyFormFieldOnChange(tableSelection.selector)

			var nextButtonSelector = '#' + elemPrefix + 'NewDashboardComponentSelectTableNextButton'
			initButtonClickHandler(nextButtonSelector,function() {
				if(validateTableSelectionForm()) {
					transitionToNextWizardDlgPanelByID($dialog,dashboardComponentValueGroupingPanelID)
				} // if validate panel's form
			})


		}, // init panel
		transitionIntoPanel: function ($dialog) {
			setWizardDialogButtonSet("newDashboardComponentSelectTableNextButton")
			
		 } // no-op
	}
	
	return dashboardSelectTablePanelConfig

}


