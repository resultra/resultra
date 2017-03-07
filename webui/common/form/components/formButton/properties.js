
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

function loadFormButtonProperties($button,buttonRef) {
	
	console.log("Loading button properties")
	
	initFormButtonPopupBehaviorProperties(buttonRef)
	
	
	function initButtonSizeProperties() {
		var $sizeSelection = $('#adminButtonComponentSizeSelection')
		$sizeSelection.val(buttonRef.properties.size)
		initSelectControlChangeHandler($sizeSelection,function(newSize) {
		
			var sizeParams = {
				parentFormID: buttonRef.parentFormID,
				buttonID: buttonRef.buttonID,
				size: newSize
			}
			jsonAPIRequest("frm/formButton/setSize",sizeParams,function(updatedButton) {
				setContainerComponentInfo($button,updatedButton,updatedButton.buttonID)	
				setFormButtonSize($button,updatedButton.properties.size)
			})
		
		})
		
	}
	initButtonSizeProperties()


	function initColorSchemeProperties() {
		var $schemeSelection = $('#adminButtonComponentColorSchemeSelection')
		$schemeSelection.val(buttonRef.properties.colorScheme)
		initSelectControlChangeHandler($schemeSelection,function(newScheme) {
		
			var sizeParams = {
				parentFormID: buttonRef.parentFormID,
				buttonID: buttonRef.buttonID,
				colorScheme: newScheme
			}
			jsonAPIRequest("frm/formButton/setColorScheme",sizeParams,function(updatedButton) {
				setContainerComponentInfo($button,updatedButton,updatedButton.buttonID)	
				setFormButtonColorScheme($button,updatedButton.properties.colorScheme)
			})
		
		})
		
	}
	initColorSchemeProperties()
	
	
	function initIconProperties() {
		var $iconSelection = $('#adminButtonComponentIconSelection')
		$iconSelection.val(buttonRef.properties.icon)
		initSelectControlChangeHandler($iconSelection,function(newIcon) {
		
			var iconParams = {
				parentFormID: buttonRef.parentFormID,
				buttonID: buttonRef.buttonID,
				icon: newIcon
			}
			jsonAPIRequest("frm/formButton/setIcon",iconParams,function(updatedButton) {
				setContainerComponentInfo($button,updatedButton,updatedButton.buttonID)	
				setFormButtonLabel($button,updatedButton)
			})
		
		})
		
	}
	initIconProperties()
	

	
	var defaultValPropParams = {
		databaseID: designFormContext.databaseID,
		elemPrefix: "button_",
		defaultDefaultValues: buttonRef.properties.popupBehavior.defaultValues,
		updateDefaultValues: function(updatedDefaultVals) {
			console.log("Updateing default values for form button: " + JSON.stringify(updatedDefaultVals))
			
			var setDefaultValsParams = {
				parentFormID: buttonRef.parentFormID,
				buttonID: buttonRef.buttonID,
				defaultValues: updatedDefaultVals }
			
			jsonAPIRequest("frm/formButton/setDefaultVals",setDefaultValsParams,function(updatedButtonRef) {
			})
		}
	}
	initDefaultValuesPropertyPanel(defaultValPropParams)
		
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#formButtonProps')
		
	closeFormulaEditor()
	
}