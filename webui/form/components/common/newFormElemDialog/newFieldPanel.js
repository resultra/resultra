
var newFieldDialogPanelID = "newField"

function createNewFieldDialogPanelConfig(elemPrefix) {
	
	var panelSelector = "#" + elemPrefix + "NewFieldPanel"
	var fieldRefNameInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldRefName")
	var fieldNameInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldName")
	var refNameHelpPopup = createPrefixedTemplElemInfo(elemPrefix,"RefNameHelp")
	var isCalcFieldField = createPrefixedTemplElemInfo(elemPrefix,"NewFieldCalcFieldCheckbox")
	var isCalcFieldInput = createPrefixedTemplElemInfo(elemPrefix,"NewFieldIsCalcFieldInput")
	var fieldTypeSelection = createPrefixedTemplElemInfo(elemPrefix,"NewFieldValTypeSelection")
	
	var newFieldPanelConfig = {
		panelID: newFieldDialogPanelID,
		divID: panelSelector,
		progressPerc:40,
		dlgButtons: { 
			"Previous": function() { 
				transitionToPrevWizardDlgPanelByPanelID(this,textBoxProgressDivID,
							newFieldDialogPanelID,createNewOrExistingFieldDialogPanelID)	
			 },
			"Next" : function() { 
				if($(panelSelector).form('validate form')) {
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
		initPanel: function () {
			
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
			
			$(panelSelector ).form({
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
		  	})
			$(refNameHelpPopup.selector).popup({on: 'hover'});
		
		}
	} // wizard dialog configuration for panel to create new field
	
	return newFieldPanelConfig;
	
}

