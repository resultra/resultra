


function loadFormButtonProperties($button,buttonRef) {
	
	console.log("Loading button properties")
	
	function saveBehaviorProperties(popupBehavior) {
		var setPopupBehaviorParams = { 
			parentFormID: buttonRef.parentFormID,
			buttonID: buttonRef.buttonID,
			popupBehavior: popupBehavior
		}
		jsonAPIRequest("frm/formButton/setPopupBehavior",setPopupBehaviorParams,function(updatedButtonRef) {
		})
	}
	initFormButtonBehaviorProperties(buttonRef,saveBehaviorProperties)
	
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
	
	var elemPrefix = "button_"
	
	var defaultValPropParams = {
		databaseID: designFormContext.databaseID,
		elemPrefix: elemPrefix,
		defaultDefaultValues: buttonRef.properties.defaultValues,
		updateDefaultValues: function(updatedDefaultVals) {
			console.log("Updateing default values for form button: " + JSON.stringify(updatedDefaultVals))
			
			var setDefaultValsParams = {
				parentFormID: buttonRef.parentFormID,
				buttonID: buttonRef.buttonID,
				defaultValues: updatedDefaultVals }
			
			jsonAPIRequest("frm/formButton/setDefaultVals",setDefaultValsParams,function(updatedButtonRef) {
					setContainerComponentInfo($button,updatedButtonRef,updatedButtonRef.buttonID)	
			})
		}
	}
	initDefaultValuesPropertyPanel(defaultValPropParams)
	
	var visibilityElemPrefix = "buttonVisibility_"
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
				parentFormID: buttonRef.parentFormID,
				buttonID: buttonRef.buttonID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/formButton/setVisibility",params,function(updatedButtonRef) {
			setContainerComponentInfo($button,updatedButtonRef,updatedButtonRef.buttonID)	
		})
	}
	var visibilityParams = {
		elemPrefix: visibilityElemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: buttonRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: buttonRef.parentFormID,
		componentID: buttonRef.buttonID,
		componentLabel: 'form button',
		$componentContainer: $button
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	
	function saveButtonLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for form button")
		var formatParams = {
			parentFormID: buttonRef.parentFormID,
			buttonID: buttonRef.buttonID,
			buttonLabelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/formButton/setButtonLabelFormat", formatParams, function(updatedButton) {
				setContainerComponentInfo($button,updatedButton,updatedButton.buttonID)	
				setFormButtonLabel($button,updatedButton)
		})	
	}
	var buttonLabelParams = {
		elemPrefix: elemPrefix,
		initialVal: buttonRef.properties.buttonLabelFormat,
		saveLabelPropsCallback: saveButtonLabelProps
	}
	initFormButtonLabelPropertyPanel(buttonLabelParams)
	
		
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#formButtonProps')
		
	closeFormulaEditor()
	
}