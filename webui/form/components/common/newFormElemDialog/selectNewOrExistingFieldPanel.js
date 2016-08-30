
var createNewOrExistingFieldDialogPanelID = "newOrExistingField"


function createNewOrExistingFieldPanelContextBootstrap(panelConfig) {


	// Build up a set of selectors based upon the prefix. The suffixes must match
	// those given in the template newFormElemDialogCommon.html
	var panelSelector = "#" + panelConfig.elemPrefix + "SelectExistingOrNewFieldPanel"
	var selectExistingField = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"SelectExistingFieldField")
	
	var formSelector = "#" + panelConfig.elemPrefix + "SelectExistingOrNewFieldForm"
	var $panelForm = $(formSelector)
	
	// Selectors & ids for individual radio buttons
	var createNewFieldRadio = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"CreateNewFieldRadio")
	var useExistingFieldRadio = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"UseExistingFieldRadio")
	var useExistingGlobalRadio = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"UseExistingGlobalRadio")

	// Selectors for field and global selection
	var selectField = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"FieldSelection")
	var selectGlobal = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"GlobalSelection")
	
	// Selector for the radio group as a whole
	var newOrExistingRadioInputSelector = "input[name='" + panelConfig.elemPrefix + "NewOrExistingRadio']"
	
	
	var newOrExistingRadioInputCheckedSelector = newOrExistingRadioInputSelector + ":checked"
	
	
	

	var panelID = "newOrExistingField"
		
				
	function doneButtonClicked() {
		if($panelForm.valid()) {
			panelConfig.doneFunc(this)	
		}
	}
	
	function cancelButtonClicked() {
		$(this).dialog('close');	
	}
	
	function initSelectNewOrExistingFieldPanel($parentDialog) {
		
		var validationRules = {}
			
		validationRules[selectField.id] = { 
			required: {
				depends: function(element) {
					return radioButtonIsChecked(useExistingFieldRadio.selector)
				}
			} 
		}
		validationRules[selectGlobal.id] = { 
			required: {
				depends: function(element) {
					return radioButtonIsChecked(useExistingGlobalRadio.selector)
				}
			} 
		}
		var validator = $panelForm.validate({rules: validationRules })
		
		validator.resetForm()	
		
		
		// Populate the select field dialog box with a list of possible fields to
		// connect the new form element to.
		$(selectField.selector).dropdown()
		loadFieldInfo(panelConfig.parentTableID,panelConfig.fieldTypes,function(fieldsByID) {
			populateFieldSelectionMenu(fieldsByID,selectField.selector)
		})
		
		populateGlobalSelectionMenu(selectGlobal.selector,panelConfig.databaseID)

		// By default, the radio button to create a new field is selected.
		setWizardDialogButtonSet("newFormComponentDlgNewFieldButtons")
		disableFormControl(selectGlobal.selector)
		disableFormControl(selectField.selector)			
		$(createNewFieldRadio.selector).prop("checked", true);
		
		// Listen to the radio buttons for changes. Depending upon 
		// which radio button is selected, the additional selection for existing
		// fields or globals is also enabled or disabled.
		$(newOrExistingRadioInputSelector).change(function() {
			console.log("new or existing radio value:", this.value);
			if (this.value == "new") {
				setWizardDialogButtonSet("newFormComponentDlgNewFieldButtons")
				disableFormControl(selectGlobal.selector)
				disableFormControl(selectField.selector)
			} else if (this.value == "existing") {
				setWizardDialogButtonSet("newFormComponentDlgExistingFieldButtons")
				disableFormControl(selectGlobal.selector)
				enableFormControl(selectField.selector)
			} else if (this.value == "newGlobal") {
				setWizardDialogButtonSet("newFormComponentDlgNewFieldButtons")
				disableFormControl(selectGlobal.selector)
				disableFormControl(selectField.selector)			
			} else {
				setWizardDialogButtonSet("newFormComponentDlgExistingFieldButtons")
				enableFormControl(selectGlobal.selector)
				disableFormControl(selectField.selector)
			}
		});
					
		var nextButtonSelector = '#' + panelConfig.elemPrefix + 'NewFormComponentNewFieldNextButton'
		initButtonClickHandler(nextButtonSelector,function() {
			if (validateForm()) {
				if (radioButtonIsChecked(createNewFieldRadio.selector)) {
					transitionToNextWizardDlgPanelByID($parentDialog,newFieldDialogPanelID)
				} else {
					//transitionToNextWizardDlgPanel(this, dialogProgressDivID,
					//	newOrExistingFieldPanelConfig, newTextBoxValidateFormatEntriesPanel)
				}
			} // if validate form
		})
		
		var doneButtonSelector = '#' + panelConfig.elemPrefix + 'NewFormComponentNewFieldDoneButton'
		initButtonClickHandler(doneButtonSelector,function() {
			if($panelForm.valid()) {
				panelConfig.doneFunc($parentDialog)	
			}
		})
		
	}
	
	function getPanelValues() {
		var panelVals = {}
		if (radioButtonIsChecked(createNewFieldRadio.selector)) {
			panelVals.newField = true
		} else {
			panelVals.newField = false
			panelVals.selectedFieldID = $(fieldSelectionSelector).val()
		}
		return panelVals
	}
		
	var newOrExistingFieldPanelConfig = {
		panelID: createNewOrExistingFieldDialogPanelID,
		divID: panelSelector,
		progressPerc: 20,
		dlgButtons: null, // dialog buttons - TODO - reimplement with Bootstrap buttons
		initPanel: initSelectNewOrExistingFieldPanel, // init panel
		getPanelVals: getPanelValues,
		transitionIntoPanel: function ($dialog) {
			if (radioButtonIsChecked(createNewFieldRadio.selector)) {
				setWizardDialogButtonSet("newFormComponentDlgNewFieldButtons")				
			} else {
				setWizardDialogButtonSet("newFormComponentDlgExistingFieldButtons")				
			}
		}
	} // wizard dialog configuration for panel to create new field

	return newOrExistingFieldPanelConfig;

} // createNewOrExistingFieldPanelConfig
