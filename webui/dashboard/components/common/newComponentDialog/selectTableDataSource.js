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
	
	var dashboardSelectTablePanelConfig = {
		divID: panelSelector,
		panelID: dashboardComponentSelectTablePanelID,
		progressPerc:20,
		dlgButtons: { 
			"Next" : function() { 
				if(validateTableSelectionForm()) {
				
					// Since fiels have tables as their parent, the field selection can only be 
					// initialized after table selection has been made.
					var formID = '#newBarchartDialog'
					var selectedTableID = $(tableSelection.selector).val()
					console.log("Selected table: " + selectedTableID)
					/*
					loadFieldInfo(selectedTableID,[fieldTypeAll],function(fieldsByID) {
						populateFieldSelectionMenu(fieldsByID,'#xAxisFieldSelection')
						populateFieldSelectionMenu(fieldsByID,'#yAxisFieldSelection')
						newBarChartParams.fieldsByID = fieldsByID
					})
				
					transitionToNextWizardDlgPanel(this,newBarChartParams.progressDivID,
							barChartTablePanelConfig,barChartXAxisPanelConfig)
					*/
				} // if validate panel's form
			},
			"Cancel" : function() { $(this).dialog('close'); },
	 	}, // dialog buttons
		initPanel: function() {
			var tableListParams =  { databaseID: databaseID }
			jsonAPIRequest("table/getList",tableListParams,function(tableRefs) {
				populateTableSelectionMenu(tableRefs,tableSelection.selector)
			})
			
			revalidateNonEmptyFormFieldOnChange(tableSelection.selector)
					
			return {}
		}, // init panel
	}
	
	return dashboardSelectTablePanelConfig

}


