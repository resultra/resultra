var dashboardSelectTablePanelConfig = {
	divID: "#newBarChartSelectTablePanel",
	panelID: "barChartSelectTable",
	progressPerc:20,
	dlgButtons: { 
		"Next" : function() { 
			if($( "#newBarChartSelectTablePanel" ).form('validate form')) {
				
				// Since fiels have tables as their parent, the field selection can only be 
				// initialized after table selection has been made.
				var formID = '#newBarchartDialog'
				var selectedTableID = getFormStringValue(formID,'barChartTableSelection')
				loadFieldInfo(selectedTableID,[fieldTypeAll],function(fieldsByID) {
					populateFieldSelectionMenu(fieldsByID,'#xAxisFieldSelection')
					populateFieldSelectionMenu(fieldsByID,'#yAxisFieldSelection')
					newBarChartParams.fieldsByID = fieldsByID
				})
				
				transitionToNextWizardDlgPanel(this,newBarChartParams.progressDivID,
						barChartTablePanelConfig,barChartXAxisPanelConfig)
			} // if validate panel's form
		},
		"Cancel" : function() { $(this).dialog('close'); },
 	}, // dialog buttons
	initPanel: function() {
		
		
		// TODO - Add Bootstrap validation to ensure a table selection takes place.		
		
		var tableListParams =  { databaseID: databaseID }
		jsonAPIRequest("table/getList",tableListParams,function(tableRefs) {
			populateTableSelectionMenu(tableRefs,"#barChartTableSelection")
			$('#barChartTableSelection').dropdown()
		})
		
		return {}
	}, // init panel
}
