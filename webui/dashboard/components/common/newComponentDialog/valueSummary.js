

var dashboardComponentValueSummaryPanelID = "dashboardComponentValueSummary"


function createNewDashboardComponentValueSummaryPanelConfig(elemPrefix) {
	
	var panelSelector = "#" + elemPrefix + "ValueSummaryPanel"
	var summaryFieldSelection = createPrefixedTemplElemInfo(elemPrefix,"SummaryFieldSelection")
	var summarizeBySelection = createPrefixedTemplElemInfo(elemPrefix,"SummarizeBySelection")
	
	function validateValueSummaryPanel() {
		return true
	}
	
	function populateSummarizeBySelection(fieldType) {
		$(summarizeBySelection.selector).empty()
		$(summarizeBySelection.selector).append(defaultSelectOptionPromptHTML("Choose how to summarize values"))
		if(fieldType == fieldTypeNumber) {
			$(summarizeBySelection.selector).append(selectOptionHTML("count","Count of values"))
			$(summarizeBySelection.selector).append(selectOptionHTML("sum","Sum of values"))
			$(summarizeBySelection.selector).append(selectOptionHTML("average","Average of values"))
		}
		else if (fieldType == fieldTypeText) {
			$(summarizeBySelection.selector).append(selectOptionHTML("count","Count of values"))
		}
		else {
			console.log("unrecocognized field type: " + fieldType)
		}
	}
	
	var valueSummaryPanelConfig = {
		divID: panelSelector,
		panelID: dashboardComponentValueSummaryPanelID,
		progressPerc:80,
		dlgButtons: { 
			"Previous": function() {
				transitionToPrevWizardDlgPanelByPanelID(this,dashboardComponentValueGroupingPanelID)	
			 },
			"Done" : function() { 
				if(validateValueSummaryPanel()) {
				//	saveNewBarChart()
				} // if validate panel's form
			},
			"Cancel" : function() { $(this).dialog('close'); },
	 	}, // dialog buttons
	
	
		initPanel: function() {

			$(summarizeBySelection.selector).empty()
			$(summarizeBySelection.selector).attr("disabled",true)
		
			return {}
		},	// init panel
		
		transitionIntoPanel: function ($dialog) { 
			
			var selectedTableID = getWizardDialogPanelData($dialog,
					elemPrefix,dashboardComponentSelectTablePanelID)
			
			loadFieldInfo(selectedTableID,[fieldTypeAll],function(valueSummaryFieldsByID) {
				populateFieldSelectionMenu(valueSummaryFieldsByID,summaryFieldSelection.selector)
				
				$(summaryFieldSelection.selector).unbind("change")		
				$(summaryFieldSelection.selector).change(function(){
					var fieldID = summaryFieldSelection.val()
			        console.log("select field: " + fieldID )
					if(fieldID in valueSummaryFieldsByID) {
						fieldInfo = valueSummaryFieldsByID[fieldID]			
			        	console.log("select field: field ID = " + fieldID  + " name = " + fieldInfo.name + " type = " + fieldInfo.type)
				
						populateSummarizeBySelection(fieldInfo.type)
						$(summarizeBySelection.selector).attr("disabled",false)
					}
			    }); // change
				
			}) // loadFieldInfo
		},
	} // panel config
	
	return valueSummaryPanelConfig
}

