
function createNewOrExistingFieldPanelConfig(elemPrefix) {

	// Build up a set of selectors based upon the prefix. The suffixes must match
	// those given in the template newFormElemDialogCommon.html
	var panelSelector = "#" + elemPrefix + "SelectExistingOrNewFieldPanel"
	var selectExistingFieldName = elemPrefix + "SelectExistingFieldField"
	var selectExistingFieldSelector = "#" + selectExistingFieldName
	var createNewFieldRadioSelector = "#" + elemPrefix + "CreateNewFieldRadio"
	var newOrExistingRadioInputSelector = "input[name='" + elemPrefix + "NewOrExistingRadio']"
	var dialogProgressDivID = elemPrefix + "NewFormElemDialogProgress"
	var selectFieldDropdown = "#" + elemPrefix + "FieldSelection"
	
	var fieldSelectionPropertyName = elemPrefix + "FieldSelection"

	var newOrExistingFieldPanelConfig = {
		divID: panelSelector,
		progressPerc: 0,
		dlgButtons: {
			"Next": function() {
				if ($(panelSelector).form('validate form')) {
					console.log("New Field checked: " + doCreateNewFieldWithTextBox())
					/* not implemented yet
					if (doCreateNewFieldWithTextBox()) {
						transitionToNextWizardDlgPanel(this, dialogProgressDivID,
							newOrExistingFieldPanelConfig, newFieldPanelConfig)
					} else {
						transitionToNextWizardDlgPanel(this, dialogProgressDivID,
							newOrExistingFieldPanelConfig, newTextBoxValidateFormatEntriesPanel)
					}
					*/
				} // if validate form
			},
			"Cancel": function() {
				$(this).dialog('close');
			},
		}, // dialog buttons

		initPanel: function() {

			function enableSelectExistingField() {
				$(selectExistingFieldSelector).removeClass("disabled")
				
				var fieldValidation = {}
				fieldValidation[selectExistingFieldName] = {
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
				$(selectExistingFieldSelector).addClass("disabled")
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
			$(selectFieldDropdown).dropdown()
			loadFieldInfo(function(fieldsByID) {
				populateFieldSelectionMenu(fieldsByID,selectFieldDropdown)
			})

			disableSelectExistingField();
			$(createNewFieldRadioSelector).prop("checked", true);
			$(newOrExistingRadioInputSelector).change(function() {
				console.log("new or existing radio value:", this.value);
				if (this.value == "new") {
					disableSelectExistingField()
				} else {
					enableSelectExistingField()
				}
			});

		}
	} // wizard dialog configuration for panel to create new field

	return newOrExistingFieldPanelConfig;

} // createNewOrExistingFieldPanelConfig

