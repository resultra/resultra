
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
	
	var dialogProgressSelector = "#" + elemPrefix + "NewFormElemDialogProgress"
	
	var validationRules = {}
	validationRules[fieldNameInput.id] = { required: true } 
	validationRules[fieldTypeSelection.id] = { required: true } 
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
			type: $(fieldTypeSelection.selector).val()
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

