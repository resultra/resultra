// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

var newFieldDialogPanelID = "newField"


function createNewFieldDialogPanelContextBootstrap(panelParams) {
	
	
	var elemPrefix = panelParams.elemPrefix

	var panelSelector = "#" + elemPrefix + "NewFieldPanel"
	
	var $panelForm = $("#" + elemPrefix + "NewFieldForm")
	
	
	var fieldRefNameInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldRefName")
	var fieldNameInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldName")
	var refNameLabel = createPrefixedTemplElemInfo(elemPrefix,"RefNameLabel")
	var isCalcFieldField = createPrefixedTemplElemInfo(elemPrefix,"NewFieldCalcFieldCheckbox")
	var isCalcFieldInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldIsCalcFieldInput")
	
	var fieldTypeSelection = createPrefixedTemplElemInfo(elemPrefix,"NewFieldValTypeSelection")
	var fieldTypeSelectionFormGroup = createPrefixedTemplElemInfo(elemPrefix,"NewFieldValTypeSelectionFormGroup")
	
	var dialogProgressSelector = "#" + elemPrefix + "NewFormElemDialogProgress"
	
	// Only show the checkbox to create a calculated field when the form component links to a type of field
	// which can be a calculated field.
	if(panelParams.hideCreateCalcFieldCheckbox) {
		$(isCalcFieldField.selector).hide()
	} else {
		$(isCalcFieldField.selector).show()
	}
	
	// Initialize the field type selection: If there is only 1 field type to choose from,
	// this field type is set in singleFieldType and the form's selection for field type is
	// disabled. 
	var multipleFieldTypes = false
	var singleFieldType = null
	if (panelParams.fieldTypes.length == 1) {
		singleFieldType = panelParams.fieldTypes[0]
		$(fieldTypeSelectionFormGroup.selector).hide()
	} else {
		$(fieldTypeSelectionFormGroup.selector).show()
		multipleFieldTypes = true
		$(fieldTypeSelection.selector).empty()
		$(fieldTypeSelection.selector).append(defaultSelectOptionPromptHTML("Select a Field Type"))
		for(var fieldTypeIndex = 0; fieldTypeIndex < panelParams.fieldTypes.length; fieldTypeIndex++) {
			var fieldType = panelParams.fieldTypes[fieldTypeIndex]
			$(fieldTypeSelection.selector).append(selectOptionHTML(fieldType,fieldTypeLabel(fieldType)))
		}
		
	}
	function getFieldTypeSelection() {
		if (multipleFieldTypes) {
			return $(fieldTypeSelection.selector).val()
		} else {
			return singleFieldType
		}
	}
	
	// Set a default reference name based upon the field name the user inputs.
	// However, after the field reference name has been manually edited, don't
	// change it to the default reference name anymore.
	var fieldRefNameManuallyEdited = false
	$(fieldNameInput.selector).on('input',function() {
		if (!fieldRefNameManuallyEdited) {
			var fieldName = $(fieldNameInput.selector).val()
			var defaultFieldRefName = fieldName.replace(/[^0-9a-zA-Z]/g,"")
			$(fieldRefNameInput.selector).val(defaultFieldRefName)
			// Immediately trigger validation of the field reference name, based
			// upon the default value.
			$(fieldRefNameInput.selector).valid()
		}
	})
	$(fieldRefNameInput.selector).change(function() {
		fieldRefNameManuallyEdited = true
	})
	
	
	
	var validationRules = {}
	validationRules[fieldNameInput.id] = { required: true } 
	if (multipleFieldTypes) {
		validationRules[fieldTypeSelection.id] = { required: true } 		
	}
	validationRules[fieldRefNameInput.id] = { required: true } 
	var validator = $panelForm.validate({rules: validationRules })
	
	resetFormValidationFeedback($panelForm)
	resetAllFormInputs($panelForm)
	validator.resetForm()	
	
	$(refNameLabel.selector).tooltip()
	
	
	function getPanelValues($parentDialog) {
		var newFieldParams = {
			parentDatabaseID: panelParams.databaseID,
			name: $(fieldNameInput.selector).val(),
			refName: $(fieldRefNameInput.selector).val(),
			isCalcField: $(isCalcFieldInput.selector).prop("checked"),
			type: getFieldTypeSelection()
		}
		return newFieldParams
	}
	
	function initNewFieldPanel($parentDialog) {
		
		
		var prevButtonSelector = createPrefixedSelector(elemPrefix,'NewFormComponentSelectFieldPrevButton')
		if(panelParams.existingFieldsToChooseFrom) {
			$(prevButtonSelector).show()
			initButtonClickHandler(prevButtonSelector, function() {
				transitionToPrevWizardDlgPanelByPanelID($parentDialog,createNewOrExistingFieldDialogPanelID)	
			})
		} else {
			$(prevButtonSelector).hide()
		}
		
		
		var doneButtonSelector = createPrefixedSelector(elemPrefix,'NewFormComponentCreateNewFieldDoneButton')
		initButtonClickHandler(doneButtonSelector, function() {
			console.log("New field done")
			if($panelForm.valid()) {
				console.log("New field panel form validated")
				var newFieldParams = getPanelValues($parentDialog)
				
				if (newFieldParams.isCalcField) {
					newFieldParams.formulaText = "" // initially empty formula
					console.log("creating new calculated field: params= " + JSON.stringify(newFieldParams))
					jsonAPIRequest("calcField/new",newFieldParams,function(newField) {
						console.log("new field created: " + JSON.stringify(newField))
									
						// Re-initialize the field information used by different elements.
						initFieldInfo(panelParams.databaseID,function() {
							var newComponentAPIParams = {
								parentFormID: panelParams.formID,
								geometry: panelParams.containerParams.geometry,
								fieldID: newField.fieldID
							}
							panelParams.createNewFormComponent($parentDialog,newComponentAPIParams)
						
							panelParams.doneFunc($parentDialog)
						})
									
					})
					
				} else {
					console.log("creating new field: params= " + JSON.stringify(newFieldParams))
					jsonAPIRequest("field/new",newFieldParams,function(newField) {
						console.log("new field created: " + JSON.stringify(newField))
									
						// Re-initialize the field information used by different elements.
						initFieldInfo(panelParams.databaseID,function() {
							var newComponentAPIParams = {
								parentFormID: panelParams.formID,
								geometry: panelParams.containerParams.geometry,
								fieldID: newField.fieldID
							}
							panelParams.createNewFormComponent($parentDialog,newComponentAPIParams)
						
							panelParams.doneFunc($parentDialog)
						})
									
					})
					
				}
				

			}	
		})
		
	}
	
	var newFieldPanelConfig = {
		panelID: newFieldDialogPanelID,
		divID: panelSelector,
		progressPerc:60,
		dlgButtons: null, // todo - initialize buttons with Bootstrap based button handling
		initPanel: initNewFieldPanel, // init panel
		getPanelVals: getPanelValues,
		transitionIntoPanel: function ($dialog) { 
			setWizardDialogButtonSet('newFormComponentNewFieldButtons')
		}
	} // wizard dialog configuration for panel to create new field
	
	return newFieldPanelConfig;
	
}

