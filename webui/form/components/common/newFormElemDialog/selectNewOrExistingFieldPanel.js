
var createNewOrExistingFieldDialogPanelID = "newOrExistingField"

function createNewOrExistingFieldPanelConfig(panelConfig) {

	// Build up a set of selectors based upon the prefix. The suffixes must match
	// those given in the template newFormElemDialogCommon.html
	var panelSelector = "#" + panelConfig.elemPrefix + "SelectExistingOrNewFieldPanel"
	var selectExistingField = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"SelectExistingFieldField")
	
	var formSelector = "#" + panelConfig.elemPrefix + "SelectExistingOrNewFieldForm"
	
	var createNewFieldRadio = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"CreateNewFieldRadio")
	var newOrExistingRadioInputSelector = "input[name='" + panelConfig.elemPrefix + "NewOrExistingRadio']"
	var newOrExistingRadioInputCheckedSelector = newOrExistingRadioInputSelector + ":checked"
	var dialogProgressDivID = panelConfig.elemPrefix + "NewFormElemDialogProgress"
	var dialogProgressSelector = '#' + dialogProgressDivID
	var selectField = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"FieldSelection")
	
	var fieldSelectionPropertyName = panelConfig.elemPrefix + "FieldSelection"
	var fieldSelectionSelector = '#' + fieldSelectionPropertyName
	
	var panelID = "newOrExistingField"
		
	function validateForm() {
		if(radioButtonIsChecked(createNewFieldRadio.selector)) {
			return true			
		} else {
			if(formFieldValueIsEmpty(fieldSelectionSelector)) {
				addFormControlError(selectExistingField.selector)
				return false
			}
			removeFormControlError(selectExistingField.selector)				
			return true;
		}
		console.log("createNewOrExistingFieldPanelConfig: radio selection: " + newOrExistingSelection)
	}
	
	// Remove any errors on the selection if a non-empty value is selected.
	$(fieldSelectionSelector).change(function() {
		if(formFieldValueIsNonEmpty(fieldSelectionSelector)) {
			removeFormControlError(selectExistingField.selector)			
		}
	})
	
	
	function nextButtonClicked() {
		if (validateForm()) {
			if (radioButtonIsChecked(createNewFieldRadio.selector)) {
				transitionToNextWizardDlgPanelByID(this,newFieldDialogPanelID)
			} else {
				//transitionToNextWizardDlgPanel(this, dialogProgressDivID,
				//	newOrExistingFieldPanelConfig, newTextBoxValidateFormatEntriesPanel)
			}
			
		} // if validate form
	}
	
	function doneButtonClicked() {
		if(validateForm()) {
			panelConfig.doneFunc(this)	
		}
	}
	
	function cancelButtonClicked() {
		$(this).dialog('close');	
	}
	
	var selectExistingButtons = {
		"Done": doneButtonClicked,
		"Cancel": cancelButtonClicked
	}
	
	var selectNewButtons = {
		"Next": nextButtonClicked,
		"Cancel": cancelButtonClicked,		
	}

	var newOrExistingFieldPanelConfig = {
		panelID: createNewOrExistingFieldDialogPanelID,
		divID: panelSelector,
		progressPerc: 20,
		dlgButtons: selectNewButtons, // dialog buttons

		initPanel: function(parentDialog) {

			function enableSelectExistingField() {
				setWizardDialogButtons(parentDialog,selectExistingButtons)
				console.log("Enabling field selection")
				enableFormControl(fieldSelectionSelector)				
			}

			function disableSelectExistingField() {
				setWizardDialogButtons(parentDialog,selectNewButtons)
				disableFormControl(fieldSelectionSelector)
							
				validateForm()
			}

			// Populate the select field dialog box with a list of possible fields to
			// connect the new form element to.
			$(selectField.selector).dropdown()
			loadFieldInfo(panelConfig.parentTableID,panelConfig.fieldTypes,function(fieldsByID) {
				populateFieldSelectionMenu(fieldsByID,selectField.selector)
			})

			disableSelectExistingField();
			$(createNewFieldRadio.selector).prop("checked", true);
			$(newOrExistingRadioInputSelector).change(function() {
				console.log("new or existing radio value:", this.value);
				if (this.value == "new") {
					disableSelectExistingField()
					removeFormControlError(selectExistingField.selector)		
				} else {
					enableSelectExistingField()
				}
			});
			
			
			var panelFormInfo = {
				panelSelector: panelSelector,
				existingFieldSelection: selectField.id,
				existingFieldSelectionSelector: selectField.selector,
				newFieldRadio: createNewFieldRadio.id,
				newFieldRadioSelector: createNewFieldRadio.selector
			}
			
			return panelFormInfo;

		} // init panel
	} // wizard dialog configuration for panel to create new field

	return newOrExistingFieldPanelConfig;

} // createNewOrExistingFieldPanelConfig

