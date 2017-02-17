
var createNewOrExistingFieldDialogPanelID = "newOrExistingField"


function createNewOrExistingFieldPanelContextBootstrap(panelConfig) {


	// Build up a set of selectors based upon the prefix. The suffixes must match
	// those given in the template newFormElemDialogCommon.html
	var panelSelector = "#" + panelConfig.elemPrefix + "SelectExistingOrNewFieldPanel"	
	var $panelForm = $("#" + panelConfig.elemPrefix + "SelectExistingOrNewFieldForm")
	
	// Selectors & ids for individual radio buttons
	var createNewFieldRadio = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"CreateNewFieldRadio")
	var useExistingFieldRadio = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"UseExistingFieldRadio")
	// Selector for the radio group as a whole
	var newOrExistingRadioInputSelector = "input[name='" + panelConfig.elemPrefix + "NewOrExistingRadio']"
	var checkedComponentValRadioSelector = "input[name='" + panelConfig.elemPrefix + "NewOrExistingRadio']:checked"

	// Selectors for field and global selection
	var selectField = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"FieldSelection")
							
	function initSelectNewOrExistingFieldPanel($parentDialog) {
		
		var validationRules = {}
			
		validationRules[selectField.id] = { 
			required: {
				depends: function(element) {
					return radioButtonIsChecked(useExistingFieldRadio.selector)
				}
			} 
		}
		var validator = $panelForm.validate({rules: validationRules })
		
		validator.resetForm()	
		
		
		// Populate the select field dialog box with a list of possible fields to
		// connect the new form element to.
		$(selectField.selector).dropdown()
		loadFieldInfo(panelConfig.databaseID,panelConfig.fieldTypes,function(fieldsByID) {
			populateFieldSelectionMenu(fieldsByID,selectField.selector)
		})
		

		// By default, the radio button to create a new field is selected.
		setWizardDialogButtonSet("newFormComponentDlgNewFieldButtons")
		disableFormControl(selectField.selector)			
		$(createNewFieldRadio.selector).prop("checked", true);
		
		// Listen to the radio buttons for changes. Depending upon 
		// which radio button is selected, the additional selection for existing
		// fields or globals is also enabled or disabled.
		$(newOrExistingRadioInputSelector).change(function() {
			console.log("new or existing radio value:", this.value);
			if (this.value == "newField") {
				setWizardDialogButtonSet("newFormComponentDlgNewFieldButtons")
				disableFormControl(selectField.selector)
			} else if (this.value == "existingField") {
				setWizardDialogButtonSet("newFormComponentDlgExistingFieldButtons")
				enableFormControl(selectField.selector)
			} else {
				setWizardDialogButtonSet("newFormComponentDlgExistingFieldButtons")
				disableFormControl(selectField.selector)
			}
		});
					
		var nextButtonSelector = '#' + panelConfig.elemPrefix + 'NewFormComponentNewFieldNextButton'
		initButtonClickHandler(nextButtonSelector,function() {
			if ($panelForm.valid()) {
				transitionToNextWizardDlgPanelByID($parentDialog,newFieldDialogPanelID)
			} // if validate form
		})
		
		var doneButtonSelector = '#' + panelConfig.elemPrefix + 'NewFormComponentNewFieldDoneButton'
		initButtonClickHandler(doneButtonSelector,function() {
			if($panelForm.valid()) {
				var newComponentAPIParams = {
					parentFormID: panelConfig.formID,
					geometry: panelConfig.containerParams.geometry,
					fieldID: $(selectField.selector).val()
				}
				componentCreated = true
				panelConfig.createNewFormComponent($parentDialog,newComponentAPIParams)					
				panelConfig.doneFunc($parentDialog)	
			}
		})
		
	}
	
	function getPanelValues() {
		var panelVals = {}
		var componentValSelection = $(checkedComponentValRadioSelector).val()
		panelVals.componentValSelection = componentValSelection
		panelVals.linkedValType = linkedComponentValTypeField
		panelVals.selectedFieldID = $(selectField.selector).val()
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
