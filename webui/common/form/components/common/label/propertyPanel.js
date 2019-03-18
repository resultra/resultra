// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initComponentLabelPropertyPanel(params) {
	console.log("Initializing component label property panel")
	
	var $form = createPrefixedContainerObj(params.elemPrefix,'ComponentLabelPropsForm')
	var $customLabelInput = createPrefixedContainerObj(params.elemPrefix,'ComponentLabelPropsCustomLabelInput')
	var customLabelInputSelector = createPrefixedSelector(params.elemPrefix,'ComponentLabelPropsCustomLabelInput')
	var $labelTypeSelection = createPrefixedContainerObj(params.elemPrefix,'ComponentLabelPropsLabelType')
	var $customLabelFormGroup = createPrefixedContainerObj(params.elemPrefix,'ComponentLabelCustomLabelFormGroup')
	
	$customLabelInput.val(params.initialVal.customLabel)
	$labelTypeSelection.val(params.initialVal.labelType)
	
	
	function updateCustomLabelVisibility() {
		var labelType = $labelTypeSelection.val()
		if (labelType === "custom") {
			$customLabelFormGroup.show()
		} else {
			$customLabelFormGroup.hide()
		}
	}
	updateCustomLabelVisibility()
		
	var remoteValidationParams = {
		url: '/api/generic/stringValidation/validateOptionalItemLabel',
		data: {
			label: function() { return $customLabelInput.val(); }
		}
	}
	
	var customLabelInputName = params.elemPrefix + 'ComponentLabelPropsCustomLabelInput'
	var validationRules = {}
	validationRules[customLabelInputName] = { remote: remoteValidationParams }

	var validationSettings = createInlineFormValidationSettings({ rules: validationRules })	
	
	
	var validator = $form.validate(validationSettings)
	
	function saveValidatedLabel(validatedLabel) {
		var updatedLabelProps = {
			labelType: $labelTypeSelection.val(),
			customLabel:validatedLabel,
		}
		console.log("updating component label properties: " + JSON.stringify(updatedLabelProps))
		params.saveLabelPropsCallback(updatedLabelProps)
		
	}
	
	initSelectControlChangeHandler($labelTypeSelection,function(newLabelType) {
		console.log("Selection changed for label type: " + newLabelType)
		updateCustomLabelVisibility()
		if(newLabelType == "custom") {
			// If the label type is custom, the label itself must also be validated
			// when the custom label type is selected.
			validateRemoteInlineInput(validator,customLabelInputSelector,
				remoteValidationParams, saveValidatedLabel)			
		} else {
			var updatedLabelProps = {
				labelType: $labelTypeSelection.val(),
				customLabel:"",
			}
			console.log("updating component label properties: " + JSON.stringify(updatedLabelProps))
			params.saveLabelPropsCallback(updatedLabelProps)
		}
	})

	initInlineInputValidationOnBlur(validator,customLabelInputSelector,
		remoteValidationParams, saveValidatedLabel)	
	
}