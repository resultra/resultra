function initFormButtonLabelPropertyPanel(params) {
	console.log("Initializing form button label property panel")
	
	var $form = createPrefixedContainerObj(params.elemPrefix,'FormButtonLabelPropsForm')
	var $customLabelInput = createPrefixedContainerObj(params.elemPrefix,'FormButtonLabelPropsCustomLabelInput')
	var customLabelInputSelector = createPrefixedSelector(params.elemPrefix,'FormButtonLabelPropsCustomLabelInput')
	var $labelTypeSelection = createPrefixedContainerObj(params.elemPrefix,'FormButtonLabelPropsLabelType')
	var $customLabelFormGroup = createPrefixedContainerObj(params.elemPrefix,'FormButtonLabelCustomLabelFormGroup')
	
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
	
	var customLabelInputName = params.elemPrefix + 'FormButtonLabelPropsCustomLabelInput'
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