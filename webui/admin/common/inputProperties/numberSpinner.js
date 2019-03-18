// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initSpinnerButtonProps(params) {
	
	var $showSpinner = $('#numberInputShowValueSpinnerButtons')
	initCheckboxControlChangeHandler($showSpinner, params.initialShowSpinner,function(showSpinner) {
		console.log("Update spinner buttons show/hide:" + showSpinner)
		params.setShowSpinner(showSpinner)
	})
	
	var validationSettings = createInlineFormValidationSettings({
		rules: {
			numberInputSpinnerButtonStep: {
				required: true,
				positiveNumber: true
			}
		},
		messages: {
			numberInputSpinnerButtonStep: {
				positiveNumber: "Step value must be a positive number.",
				required: "Step value must be a positive number."
			}
		}
	})	
	var $form = $('#numberSpinnerPropsForm')
	var validator = $form.validate(validationSettings)
	
	var $stepSizeInput = $('#numberInputSpinnerButtonStep')
	$stepSizeInput.val(params.initialStepSize)
	function setStepSizeIfValid() {
		if(validator.valid()) {
			var stepSize = Number($stepSizeInput.val())
			console.log("Setting step size:" + stepSize)
			params.setStepSize(stepSize)
			
		}
	}
	
	$stepSizeInput.unbind("blur")
	$stepSizeInput.blur(function() { setStepSizeIfValid() })
	
	
}
