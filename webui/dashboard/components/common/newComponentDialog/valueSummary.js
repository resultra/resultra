

var dashboardComponentValueSummaryPanelID = "dashboardComponentValueSummary"

function createNewDashboardComponentValueSummaryPanelConfig(elemPrefix,doneCallbackFunc,databaseID) {
	
	var panelSelector = "#" + elemPrefix + "ValueSummaryPanel"
	var summaryFieldSelection = createPrefixedTemplElemInfo(elemPrefix,"NewComponentSummaryFieldSelection")
	var summarizeBySelection = createPrefixedTemplElemInfo(elemPrefix,"NewComponentSummarizeBySelection")
	
	function validateValueSummaryPanel() {
		var validationResults = true
		
		// Any one of the fields not passing validation makes the whole validation fail
		if(!validateNonEmptyFormField(summaryFieldSelection.selector)) { validationResults = false }
		if(!validateNonEmptyFormField(summarizeBySelection.selector)) { validationResults = false }
		
		return validationResults
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

	function getPanelValues() {
		var valSummary = {
			fieldID: summaryFieldSelection.val(),
			summarizeValsWith: summarizeBySelection.val()}
		return valSummary
	}

	
	var valueSummaryPanelConfig = {
		divID: panelSelector,
		panelID: dashboardComponentValueSummaryPanelID,
		progressPerc:80,
		getPanelVals:getPanelValues,
		initPanel: function($dialog) {

			$(summarizeBySelection.selector).empty()
			$(summarizeBySelection.selector).attr("disabled",true)

			revalidateNonEmptyFormFieldOnChange(summaryFieldSelection.selector)
			revalidateNonEmptyFormFieldOnChange(summarizeBySelection.selector)

			var prevButtonSelector = '#' + elemPrefix + 'NewDashboardComponentValueSummaryPrevButton'
			initButtonClickHandler(prevButtonSelector,function() {
				transitionToPrevWizardDlgPanelByPanelID($dialog,dashboardComponentValueGroupingPanelID)	
			})
			
			var doneButtonSelector = '#' + elemPrefix + 'NewDashboardComponentValueSummaryDoneButton'
			initButtonClickHandler(doneButtonSelector,function() {
				if(validateValueSummaryPanel()) {
					doneCallbackFunc($dialog)
				} // if validate panel's form
			})
		},	// init panel
		
		transitionIntoPanel: function ($dialog) { 

			setWizardDialogButtonSet("newDashboardComponentValueSummaryButtons")
			
			loadSortedFieldInfo(databaseID,[fieldTypeAll],function(sortedValueSummaryFields) {
				
				var valueSummaryFieldsByID = createFieldsByIDMap(sortedValueSummaryFields)
				
				var $summaryFieldSelection = $(summaryFieldSelection.selector)
				
				populateSortedFieldSelectionMenu($summaryFieldSelection,sortedValueSummaryFields)
				
				$summaryFieldSelection.unbind("change")		
				$summaryFieldSelection.change(function(){
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

