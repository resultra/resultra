
var FormButtonPopupBehaviorModal = "modal"

function initFormButtonPopupBehaviorProperties(buttonRef) {
	
	var $behaviorSelection = $('#formButtonPopupBehaviorSelection')
	var $customModalSaveLabelInput = $('#formButtonCustomModalSaveLabelInput')
			
	$behaviorSelection.val(buttonRef.properties.popupBehavior.popupMode)
	$customModalSaveLabelInput.val(buttonRef.properties.popupBehavior.customLabelModalSave)
	
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
		var setPopupBehaviorParams = { 
			parentFormID: buttonRef.parentFormID,
			buttonID: buttonRef.buttonID,
			popupBehavior: {
				popupMode: $behaviorSelection.val(),
				customLabelModalSave: $customModalSaveLabelInput.val()
			}
		}
		jsonAPIRequest("frm/formButton/setPopupBehavior",setPopupBehaviorParams,function(updatedButtonRef) {
			updateModelLabelVisibility(updatedButtonRef.properties.popupBehavior.popupMode)
		})
		
	}
	
	initSelectControlChangeHandler($behaviorSelection, function() {
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

function loadFormButtonProperties(buttonRef) {
	
	console.log("Loading button properties")
	
	initFormButtonPopupBehaviorProperties(buttonRef)
		
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#formButtonProps')
		
	closeFormulaEditor()
	
}