// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


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
	
	function populateNewComponentSummarizeBySelection(fieldType) {
		populateSummarizeBySelection(summarizeBySelection.selector,fieldType)
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
				
						populateNewComponentSummarizeBySelection(fieldInfo.type)
						$(summarizeBySelection.selector).attr("disabled",false)
					}
			    }); // change
				
			}) // loadFieldInfo
		},
	} // panel config
	
	return valueSummaryPanelConfig
}

