function loadTextBoxProperties($textBox,textBoxRef) {
	console.log("loading text box properties")
	
	var elemPrefix = "textBox_"
	
	
	var validationParams = {
		initialValidationProps: textBoxRef.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentFormID: textBoxRef.parentFormID,
				textboxID: textBoxRef.textBoxID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("frm/textBox/setValidation",validationParams,function(updatedCheckboxRef) {
				setContainerComponentInfo($textBox,updatedCheckboxRef,updatedCheckboxRef.textBoxID)
			})
		
		}
	}
	initTextInputValidationProperties(validationParams)
	
	
	function dummySetVal(dropdownVal) {}
	
	function saveValueList(valueListID) {
		var setValueListParams = {
			parentFormID: textBoxRef.parentFormID,
			textboxID: textBoxRef.textBoxID,
			valueListID: valueListID
		}
		jsonAPIRequest("frm/textBox/setValueList", setValueListParams, function(updatedTextBox) {
			setContainerComponentInfo($textBox,updatedTextBox,updatedTextBox.textBoxID)
			configureTextBoxComponentValueListDropdown($textBox, updatedTextBox, dummySetVal)
		})			
	}
	var valueListPropertyParams = {
		elemPrefix: elemPrefix,
		databaseID: designFormContext.databaseID,
		saveValueListCallback: saveValueList,
		defaultValueListID: textBoxRef.properties.valueListID
	}
	initValueListSelectionPropertyPanel(valueListPropertyParams)
		
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: textBoxRef.parentFormID,
			textboxID: textBoxRef.textBoxID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/textBox/setLabelFormat", formatParams, function(updatedTextBox) {
			setTextBoxComponentLabel($textBox,updatedTextBox)
			setContainerComponentInfo($textBox,updatedTextBox,updatedTextBox.textBoxID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: textBoxRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: textBoxRef.parentFormID,
			textboxID: textBoxRef.textBoxID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/textBox/setVisibility",params,function(updatedTextBox) {
			setContainerComponentInfo($textBox,updatedTextBox,updatedTextBox.textBoxID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: textBoxRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: textBoxRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: textBoxRef.parentFormID,
				textboxID: textBoxRef.textBoxID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/textBox/setPermissions",params,function(updatedTextBox) {
				setContainerComponentInfo($textBox,updatedTextBox,updatedTextBox.textBoxID)
				configureTextBoxComponentValueListDropdown($textBox, updatedTextBox, dummySetVal)
				initTextBoxClearValueControl($textBox,updatedTextBox)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var clearValueParams = {
		initialVal: textBoxRef.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentFormID: textBoxRef.parentFormID,
				textboxID: textBoxRef.textBoxID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("frm/textBox/setClearValueSupported",formatParams,function(updatedTextBox) {
				setContainerComponentInfo($textBox,updatedTextBox,updatedTextBox.textBoxID)
				initTextBoxClearValueControl($textBox,updatedTextBox)
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: textBoxRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: textBoxRef.parentFormID,
				textboxID: textBoxRef.textBoxID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/textBox/setHelpPopupMsg",params,function(updatedTextBox) {
				setContainerComponentInfo($textBox,updatedTextBox,updatedTextBox.textBoxID)
				updateComponentHelpPopupMsg($textBox, updatedTextBox)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: textBoxRef.parentFormID,
		componentID: textBoxRef.textBoxID,
		componentLabel: 'text box',
		$componentContainer: $textBox
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#textBoxProps')
		
	toggleFormulaEditorForField(textBoxRef.properties.fieldID)
		
}