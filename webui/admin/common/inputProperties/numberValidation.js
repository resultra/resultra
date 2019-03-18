// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.



function initNumberInputValidationProperties(params) {
	
	var valProps = params.initialValidationConfig
	
	var $validationForm = $('#numberInputValidationPropsForm')
	var $validationTypeSelection = $('#adminNumberInputValidationSelection')
	var $valInputParams = $validationForm.find(".numberInputParam")
	var $rangeParams = $validationForm.find(".numberInputRangeVal")
	var $compareValInputParam = $validationForm.find(".numberInputCompareVal")
	var $minParam = $validationForm.find(".numberInputMinVal")
	var $maxParam = $validationForm.find(".numberInputMaxVal")
	
	
	function getValidationConfig() {
		var validationType = $validationTypeSelection.val()
		var validationConfig = null
		switch(validationType) {
		case "none":
		case "required":
			validationConfig = { rule: validationType }
			break;
		case "greater":
			var compareVal = Number($compareValInputParam.val())
			if (!isNaN(compareVal)) {
				validationConfig = { rule: validationType, compareVal:compareVal }
			}
			break;
		case "between":
			var minVal = Number($minParam.val())
			var maxVal = Number($maxParam.val())
			if((!isNaN(minVal)) && (!isNaN(maxVal)) && maxVal > minVal) {
				validationConfig = { rule: validationType, minVal: minVal, maxVal: maxVal }
			}
			break;
		}
		
		console.log("Validation config: " + JSON.stringify(validationConfig))
		
		return validationConfig		
		
	}
	
	function updateValidationSettingsIfValid() {
		var validationConfig = getValidationConfig()
		if (validationConfig != null) {
			params.setValidation(validationConfig)
			
		}
	}
	
	function configureControlsForValidationType(validationType) {
		switch(validationType) {
		case "none":
			$valInputParams.hide()
			break;
		case "required":
			$valInputParams.hide()
			break;
		case "greater":
			$valInputParams.hide()
			$compareValInputParam.val("")
			$compareValInputParam.show()
			break;
		case "between":
			$valInputParams.hide()
			$rangeParams.val("")
			$rangeParams.show()
			break;
		}			
	}
	
	var defaultValidationType =  valProps.rule
	$validationTypeSelection.val(defaultValidationType) // Set to the default
	configureControlsForValidationType(defaultValidationType)
	$minParam.val(valProps.minVal)
	$maxParam.val(valProps.maxVal)
	$compareValInputParam.val(valProps.compareVal)
	
	
	initSelectControlChangeHandler($validationTypeSelection,function(newValidationType) {
		configureControlsForValidationType(newValidationType)
		updateValidationSettingsIfValid()
	})
	
	$valInputParams.unbind("blur")
	$valInputParams.blur(function() {
		updateValidationSettingsIfValid()
	})
}
