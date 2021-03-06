// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function initFormButtonBehaviorProperties(buttonRef,saveBehaviorCallback) {
	
	var $behaviorSelection = $('#formButtonPopupBehaviorSelection')
	var $customModalSaveLabelInput = $('#formButtonCustomModalSaveLabelInput')
	var $whereShowFormSelection = $('#formButtonWhereShowFormSelection')
			
	$behaviorSelection.val(buttonRef.properties.popupBehavior.popupMode)
	$customModalSaveLabelInput.val(buttonRef.properties.popupBehavior.customLabelModalSave)
	$whereShowFormSelection.val(buttonRef.properties.popupBehavior.whereShowForm)
	
	function updateModelLabelVisibility(popupMode) {
		var $modalLabelFormGroup = $('#formButtonCustomModalSaveLabelFormGroup')
		if(popupMode === FormButtonPopupBehaviorModal) {
			$modalLabelFormGroup.show()
		} else {
			$modalLabelFormGroup.hide()
		}	
	}
	
	updateModelLabelVisibility(buttonRef.properties.popupBehavior.popupMode)
	
	function savePropertiesFromControls() {
		var popupBehavior =  {
			popupMode: $behaviorSelection.val(),
			customLabelModalSave: $customModalSaveLabelInput.val(),
			whereShowForm: $whereShowFormSelection.val()
		}
		updateModelLabelVisibility(popupBehavior.popupMode)
		
		saveBehaviorCallback(popupBehavior)		
	}
	
	initSelectControlChangeHandler($behaviorSelection, function() {
		console.log("Popup behavior changed: " + $behaviorSelection.val())
		savePropertiesFromControls()
	})
	initSelectControlChangeHandler($whereShowFormSelection, function() {
		console.log("Popup behavior changed: " + $behaviorSelection.val())
		savePropertiesFromControls()
	})
	
	
	var $popupBehaviorForm = $('#formButtonPopupBehaviorForm')
		
	var remoteValidationParams = {
		url: '/api/generic/stringValidation/validateOptionalItemLabel',
		data: {
			label: function() { return  $customModalSaveLabelInput.val(); }
		}
	}	
			
	var validationSettings = createInlineFormValidationSettings({
		rules: {
			formButtonCustomModalSaveLabelInput: {
				remote: remoteValidationParams
			} // newRoleNameInput
		}
	})	
	
	
	var validator = $popupBehaviorForm.validate(validationSettings)
	
	initInlineInputValidationOnBlur(validator,'#formButtonCustomModalSaveLabelInput',
		remoteValidationParams, function(validatedName) {
			savePropertiesFromControls()	
	})	
	
}