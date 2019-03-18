// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function initToggleFormatProperties(params) {
	
	var $form = $('#adminToggleComponentFormatForm')
	
	var $offLabel = $form.find("input[name=adminToggleOffComponentLabel]")
	$offLabel.val(params.initialVals.offLabel)
	var validationOffLabelParams = {
		url: '/api/generic/stringValidation/validateItemLabel',
		data: { label: function() { return $offLabel.val(); } }
	}
	function saveValidatedOffLabel(newLabel) {
		console.log("Toggle format properties: new off label: " + newLabel)
		params.setOffLabel(newLabel)
	}

	var $onLabel = $form.find("input[name=adminToggleOnComponentLabel]")
	$onLabel.val(params.initialVals.onLabel)
	var validationOnLabelParams = {
		url: '/api/generic/stringValidation/validateItemLabel',
		data: { label: function() { return $onLabel.val(); } }
	}
	function saveValidatedOnLabel(newLabel) {
		console.log("Toggle format properties: new on label: " + newLabel)
		params.setOnLabel(newLabel)
	}
	
	
	var validationRules = {
		adminToggleOffComponentLabel: { remote: validationOffLabelParams, required:true },
		adminToggleOnComponentLabel: { remote: validationOnLabelParams, required:true }
	}
	var validationSettings = createInlineFormValidationSettings({ rules: validationRules })	
	var validator = $form.validate(validationSettings)
	
	initInlineInputControlValidationOnBlur(validator,$offLabel,validationOffLabelParams, saveValidatedOffLabel)
	initInlineInputControlValidationOnBlur(validator,$onLabel,validationOnLabelParams, saveValidatedOnLabel)

	
	var $toggleOffColorSchemeSelection = $('#adminToggleOffComponentColorSchemeSelection')
	$toggleOffColorSchemeSelection.val(params.initialVals.offColorScheme)
	initSelectControlChangeHandler($toggleOffColorSchemeSelection,function(newColorScheme) {
		params.setOffColorScheme(newColorScheme)
	})
	
	var $toggleOnColorSchemeSelection = $('#adminToggleOnComponentColorSchemeSelection')
	$toggleOnColorSchemeSelection.val(params.initialVals.onColorScheme)
	initSelectControlChangeHandler($toggleOnColorSchemeSelection,function(newColorScheme) {
		params.setOnColorScheme(newColorScheme)
	})
	
}
