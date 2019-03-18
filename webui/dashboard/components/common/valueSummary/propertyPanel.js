// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initDashboardValueSummaryPropertyPanel(panelParams) {
	
	var summaryFieldSelectionElemInfo = createPrefixedTemplElemInfo(panelParams.elemPrefix,"SummaryFieldSelection")
	var summarizeBySelectionElemInfo = createPrefixedTemplElemInfo(panelParams.elemPrefix,"SummarizeBySelection")
	var summarizeNumberFormatElemInfo = createPrefixedTemplElemInfo(panelParams.elemPrefix,"SummarizeNumberFormat")


	var saveChangesButtonElemInfo = createPrefixedTemplElemInfo(panelParams.elemPrefix,"ValSummaryPropertiesSaveChangesButton")
	
	var $propertyForm = $(createPrefixedSelector(panelParams.elemPrefix,"ValueSummaryPropertyPanelForm"))
	
	
	var validationRules = {}	
	validationRules[summaryFieldSelectionElemInfo.id] = { required: true }
	validationRules[summarizeBySelectionElemInfo.id] = { required: true }
	validationRules[summarizeNumberFormatElemInfo.id] = { required: true }
	
	var validationSettings = createInlineFormValidationSettings({rules: validationRules })	
	var validator = $propertyForm.validate(validationSettings)
		
	validator.resetForm()	
	disableButton(saveChangesButtonElemInfo.selector)
	
	loadSortedFieldInfo(panelParams.databaseID,[fieldTypeAll],function(sortedFields) {
		
		var valueSummaryFieldsByID = createFieldsByIDMap(sortedFields)
		
		var  $summaryFieldSelection = $(summaryFieldSelectionElemInfo.selector)
		
		// Initialize the field selection and "summarize with" selections with the existing values.
		var existingFieldInfo = valueSummaryFieldsByID[panelParams.valSummaryProps.summarizeByFieldID]
		populateSortedFieldSelectionMenu($summaryFieldSelection,sortedFields)
		$(summaryFieldSelectionElemInfo.selector).val(panelParams.valSummaryProps.summarizeByFieldID)
		
		populateSummarizeBySelection(summarizeBySelectionElemInfo.selector,existingFieldInfo.type)
		$(summarizeBySelectionElemInfo.selector).val(panelParams.valSummaryProps.summarizeValsWith)
		
		populateNumberFormatSelection($(summarizeNumberFormatElemInfo.selector))
		$(summarizeNumberFormatElemInfo.selector).val(panelParams.valSummaryProps.numberFormat)
		
		// Initially disable the save changes button, until the user makes a change
		disableButton(saveChangesButtonElemInfo.selector)
		
		
		initSelectionChangedHandler(summaryFieldSelectionElemInfo.selector, function(fieldID) {
			if(fieldID in valueSummaryFieldsByID) {
				fieldInfo = valueSummaryFieldsByID[fieldID]			
				populateSummarizeBySelection(summarizeBySelectionElemInfo.selector,fieldInfo.type)
				disableButton(saveChangesButtonElemInfo.selector)
			}
		})
		
		initSelectionChangedHandler(summarizeNumberFormatElemInfo.selector,function(summarizeBy) {
			if($propertyForm.valid()) {
				enableButton(saveChangesButtonElemInfo.selector)
			}
		})

		initSelectionChangedHandler(summarizeBySelectionElemInfo.selector, function(summarizeBy) {
			enableButton(saveChangesButtonElemInfo.selector)
		})
		
		initButtonClickHandler(saveChangesButtonElemInfo.selector, function() {
			if($propertyForm.valid()) {
				var newValSummaryParams = {
					summarizeByFieldID: summaryFieldSelectionElemInfo.val(),
					summarizeValsWith: summarizeBySelectionElemInfo.val(),
					numberFormat: summarizeNumberFormatElemInfo.val()
				}
				console.log("Saving new value summary: " + JSON.stringify(newValSummaryParams))
				panelParams.saveValueSummaryFunc(newValSummaryParams)
				disableButton(saveChangesButtonElemInfo.selector)
			}
		})
		
				
	}) // loadFieldInfo
	


}