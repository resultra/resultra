function initDashboardValueSummaryPropertyPanel(panelParams) {
	
	var summaryFieldSelectionElemInfo = createPrefixedTemplElemInfo(panelParams.elemPrefix,"SummaryFieldSelection")
	var summarizeBySelectionElemInfo = createPrefixedTemplElemInfo(panelParams.elemPrefix,"SummarizeBySelection")


	var saveChangesButtonElemInfo = createPrefixedTemplElemInfo(panelParams.elemPrefix,"ValSummaryPropertiesSaveChangesButton")
	
	var $propertyForm = $(createPrefixedSelector(panelParams.elemPrefix,"ValueSummaryPropertyPanelForm"))
	
	
	var validationRules = {}	
	validationRules[summaryFieldSelectionElemInfo.id] = { required: true }
	validationRules[summarizeBySelectionElemInfo.id] = { required: true }
	
	var validationSettings = createInlineFormValidationSettings({rules: validationRules })	
	var validator = $propertyForm.validate(validationSettings)
		
	validator.resetForm()	
	disableButton(saveChangesButtonElemInfo.selector)
	
	loadFieldInfo(panelParams.tableID,[fieldTypeAll],function(valueSummaryFieldsByID) {
		
		// Initialize the field selection and "summarize with" selections with the existing values.
		var existingFieldInfo = valueSummaryFieldsByID[panelParams.valSummaryProps.summarizeByFieldID]
		populateFieldSelectionMenu(valueSummaryFieldsByID,summaryFieldSelectionElemInfo.selector)
		$(summaryFieldSelectionElemInfo.selector).val(panelParams.valSummaryProps.summarizeByFieldID)
		
		populateSummarizeBySelection(summarizeBySelectionElemInfo.selector,existingFieldInfo.type)
		$(summarizeBySelectionElemInfo.selector).val(panelParams.valSummaryProps.summarizeValsWith)
		
		// Initially disable the save changes button, until the user makes a change
		disableButton(saveChangesButtonElemInfo.selector)
		
		
		initSelectionChangedHandler(summaryFieldSelectionElemInfo.selector, function(fieldID) {
			if(fieldID in valueSummaryFieldsByID) {
				fieldInfo = valueSummaryFieldsByID[fieldID]			
				populateSummarizeBySelection(summarizeBySelectionElemInfo.selector,fieldInfo.type)
				disableButton(saveChangesButtonElemInfo.selector)
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
				}
				console.log("Saving new value summary: " + JSON.stringify(newValSummaryParams))
				panelParams.saveValueSummaryFunc(newValSummaryParams)
				disableButton(saveChangesButtonElemInfo.selector)
			}
		})
		
				
	}) // loadFieldInfo
	


}