
var createNewOrExistingFieldDialogPanelID = "newOrExistingField"

function createNewOrExistingFieldPanelConfig(panelConfig) {

	// Build up a set of selectors based upon the prefix. The suffixes must match
	// those given in the template newFormElemDialogCommon.html
	var panelSelector = "#" + panelConfig.elemPrefix + "SelectExistingOrNewFieldPanel"
	var selectExistingField = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"SelectExistingFieldField")
	
	var createNewFieldRadio = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"CreateNewFieldRadio")
	var newOrExistingRadioInputSelector = "input[name='" + panelConfig.elemPrefix + "NewOrExistingRadio']"
	var dialogProgressDivID = panelConfig.elemPrefix + "NewFormElemDialogProgress"
	var selectField = createPrefixedTemplElemInfo(panelConfig.elemPrefix,"FieldSelection")
	
	var fieldSelectionPropertyName = panelConfig.elemPrefix + "FieldSelection"
	
	var panelID = "newOrExistingField"
	
	function doCreateNewFieldWithTextBox() {
	
		return $(panelSelector).form('get field',createNewFieldRadio.id).prop('checked')
	}
	
	function nextButtonClicked() {
		if ($(panelSelector).form('validate form')) {
			console.log("New Field checked: " + doCreateNewFieldWithTextBox())

			if (doCreateNewFieldWithTextBox()) {
				transitionToNextWizardDlgPanelByID(this, dialogProgressDivID,
						createNewOrExistingFieldDialogPanelID, newFieldDialogPanelID)
			} else {
				//transitionToNextWizardDlgPanel(this, dialogProgressDivID,
				//	newOrExistingFieldPanelConfig, newTextBoxValidateFormatEntriesPanel)
			}
			
		} // if validate form
	}
	
	function doneButtonClicked() {
		if($(panelSelector).form('validate form')) {
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
		progressPerc: 0,
		dlgButtons: selectNewButtons, // dialog buttons

		initPanel: function(parentDialog) {

			function enableSelectExistingField() {
				setWizardDialogButtons(parentDialog,selectExistingButtons)
				$(selectExistingField.selector).removeClass("disabled")
				
				var fieldValidation = {}
				fieldValidation[selectField.id] = {
					rules: [{
						type: 'empty',
						prompt: 'Please select a field'
					}]
				}
				
				$(panelSelector).form({
					fields: fieldValidation,
					inline: true,
				})
			}

			function disableSelectExistingField() {
				setWizardDialogButtons(parentDialog,selectNewButtons)
				$(selectExistingField.selector).addClass("disabled")

				$(panelSelector).form({
					fields: {},
					inline: true,
				})
				// After changing the validation rules, re-validate the form.
				// This will remove any outstanding errors, which no longer apply
				// since selection of an existing field is no longer required.
				$(panelSelector).form('validate form')
			}

			// Populate the select field dialog box with a list of possible fields to
			// connect the new form element to.
			$(selectField.selector).dropdown()
			loadFieldInfo(function(fieldsByID) {
				populateFieldSelectionMenu(fieldsByID,selectField.selector)
			},panelConfig.fieldTypes)

			disableSelectExistingField();
			$(createNewFieldRadio.selector).prop("checked", true);
			$(newOrExistingRadioInputSelector).change(function() {
				console.log("new or existing radio value:", this.value);
				if (this.value == "new") {
					disableSelectExistingField()
				} else {
					enableSelectExistingField()
				}
			});
			
			var panelFormInfo = {
				panelSelector: panelSelector,
				existingFieldSelection: selectField.id,
				newFieldRadio: createNewFieldRadio.id
			}
			
			return panelFormInfo;

		} // init panel
	} // wizard dialog configuration for panel to create new field

	return newOrExistingFieldPanelConfig;

} // createNewOrExistingFieldPanelConfig

