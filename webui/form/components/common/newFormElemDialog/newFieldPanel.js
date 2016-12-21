
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
	
	
	
	var validationRules = {}
	validationRules[fieldNameInput.id] = { required: true } 
	if (multipleFieldTypes) {
		validationRules[fieldTypeSelection.id] = { required: true } 		
	}
	validationRules[fieldRefNameInput.id] = { required: true } 
	var validator = $panelForm.validate({rules: validationRules })
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
		initButtonClickHandler(prevButtonSelector, function() {
			transitionToPrevWizardDlgPanelByPanelID($parentDialog,createNewOrExistingFieldDialogPanelID)	
		})
		var doneButtonSelector = createPrefixedSelector(elemPrefix,'NewFormComponentCreateNewFieldDoneButton')
		initButtonClickHandler(doneButtonSelector, function() {
			console.log("New field done")
			if($panelForm.valid()) {
				console.log("New field panel form validated")
				var newFieldParams = getPanelValues($parentDialog)
				console.log("creating new field: params= " + JSON.stringify(newFieldParams))
				jsonAPIRequest("field/new",newFieldParams,function(newField) {
					console.log("new field created: " + JSON.stringify(newField))
									
					// Re-initialize the field information used by different elements.
					initFieldInfo(panelParams.databaseID,function() {
						var newComponentAPIParams = {
							parentFormID: panelParams.formID,
							geometry: panelParams.containerParams.geometry,
							componentLink: {
								linkedValType: "field",
								fieldID: newField.fieldID
							}
						}
						panelParams.createNewFormComponent($parentDialog,newComponentAPIParams)
						
						panelParams.doneFunc($parentDialog)
					})
									
				})

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

