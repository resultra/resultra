function createNewFieldDialogPanelConfig(elemPrefix) {
	
	function createParamFormElemInfo(elemPrefix,elemSuffix) {
		return {
			id: elemPrefix + elemSuffix,
			selector: '#' + elemPrefix + elemSuffix
		}
	}
	
	var panelSelector = "#" + elemPrefix + "NewFieldPanel"
	var fieldRefNameInput = createParamFormElemInfo(elemPrefix,"NewFieldRefName")
	var fieldNameInput = createParamFormElemInfo(elemPrefix,"NewFieldName")
	var refNameHelpPopup = createParamFormElemInfo(elemPrefix,"RefNameHelp")
	var isCalcFieldField = createParamFormElemInfo(elemPrefix,"NewFieldCalcFieldCheckbox")
	var isCalcFieldInput = createParamFormElemInfo(elemPrefix,"NewFieldIsCalcFieldInput")
	var fieldTypeSelection = createParamFormElemInfo(elemPrefix,"NewFieldValTypeSelection")
	
	var newFieldPanelConfig = {
		divID: panelSelector,
		progressPerc:40,
		dlgButtons: { 
			"Previous": function() { 
				/*
				transitionToPrevWizardDlgPanel(this,textBoxProgressDivID,
							newFieldPanelConfig,newOrExistingFieldPanelConfig)	*/
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

