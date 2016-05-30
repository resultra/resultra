
var newFieldDialogPanelID = "newField"

function createNewFieldDialogPanelConfig(elemPrefix) {
	
	var panelSelector = "#" + elemPrefix + "NewFieldPanel"
	var fieldRefNameInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldRefName")
	var fieldNameInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldName")
	var refNameHelpPopup = createPrefixedTemplElemInfo(elemPrefix,"RefNameHelp")
	var isCalcFieldField = createPrefixedTemplElemInfo(elemPrefix,"NewFieldCalcFieldCheckbox")
	var isCalcFieldInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldIsCalcFieldInput")
	var fieldTypeSelection = createPrefixedTemplElemInfo(elemPrefix,"NewFieldValTypeSelection")
	
	var dialogProgressSelector = "#" + elemPrefix + "NewFormElemDialogProgress"
	
	function validateForm() {
		
	}
	
	var newFieldPanelConfig = {
		panelID: newFieldDialogPanelID,
		divID: panelSelector,
		progressPerc:60,
		dlgButtons: { 
			"Previous": function() { 
				transitionToPrevWizardDlgPanelByPanelID(this,createNewOrExistingFieldDialogPanelID)	
			 },
			"Next" : function() { 
				if(validateForm()) {
			/*
					if(newFieldIsCalcField()){
						transitionToNextWizardDlgPanel(this,textBoxProgressDivID,
							newFieldPanelConfig,calcFieldFormulaPanelConfig)
					}
					else {
						transitionToNextWizardDlgPanel(this,textBoxProgressDivID,
							newFieldPanelConfig,newTextBoxValidateFormatEntriesPanel)
					}	
					*/
				} // if validate panel's form	
			},
			"Cancel" : function() { $(this).dialog('close'); },
	 	}, // dialog buttons
		initPanel: function (dialog) {
			
			var fieldValidation = {}
			fieldValidation[fieldNameInput.id] = {
	          rules: [
	            {
	              type   : 'empty',
	              prompt : 'Please enter a field name'
	            }]
			} // field name input validation
			fieldValidation[fieldRefNameInput.id] = {
	          rules: [
	            {
	              type   : 'empty',
	              prompt : 'Please enter a reference name'
	            }]
			} // field reference name validation
			
	// TODO - Re-integrate with Bootstrap
/*			$(panelSelector ).form({
		    	fields: {
			        newFieldValTypeSelection: {
			          rules: [
			            {
			              type   : 'empty',
			              prompt : 'Please select a type'
			            }
			          ]
			        }, // newFieldValTypeSelection validation
		     	},
		  	}) */
	// TODO - Re-integrate with Bootstrap
	//		$(refNameHelpPopup.selector).popup({on: 'hover'});
			
			var panelFormInfo = {
				panelSelector: panelSelector
			}
			return panelFormInfo
		
		} // init panel
	} // wizard dialog configuration for panel to create new field
	
	return newFieldPanelConfig;
	
}

