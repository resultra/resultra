function loadNumberInputProperties($numberInput,numberInputRef) {
	console.log("loading text box properties")
	
	var elemPrefix = "numberInput_"
	
	var validationParams = {
		initialValidationConfig: numberInputRef.properties.validation,
		setValidation: function(validationConfig) {
			var validationParams = {
				parentFormID: numberInputRef.parentFormID,
				numberInputID: numberInputRef.numberInputID,
				validation: validationConfig
			}
			jsonAPIRequest("frm/numberInput/setValidation", validationParams, function(updatedNumberInput) {
				setContainerComponentInfo($numberInput,updatedNumberInput,updatedNumberInput.numberInputID)
			})	
		}
	}
	initNumberInputValidationProperties(validationParams)
	
	var formatSelectionParams = {
		elemPrefix: elemPrefix,
		initialFormat: numberInputRef.properties.valueFormat.format,
		formatChangedCallback: function (newFormat) {
			console.log("Format changed for text box: " + newFormat)
			
			var newValueFormat = {
				format: newFormat
			}
			var formatParams = {
				parentFormID: numberInputRef.parentFormID,
				numberInputID: numberInputRef.numberInputID,
				valueFormat: newValueFormat
			}
			jsonAPIRequest("frm/numberInput/setValueFormat", formatParams, function(updatedNumberInput) {
				setContainerComponentInfo($numberInput,updatedNumberInput,updatedNumberInput.numberInputID)
			})	
			
		}
	}
	initNumberFormatSelection(formatSelectionParams)
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: numberInputRef.parentFormID,
			numberInputID: numberInputRef.numberInputID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/numberInput/setLabelFormat", formatParams, function(updatedNumberInput) {
			setNumberInputComponentLabel($numberInput,updatedNumberInput)
			setContainerComponentInfo($numberInput,updatedNumberInput,updatedNumberInput.numberInputID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: numberInputRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: numberInputRef.parentFormID,
			numberInputID: numberInputRef.numberInputID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/numberInput/setVisibility",params,function(updatedNumberInput) {
			setContainerComponentInfo($numberInput,updatedNumberInput,updatedNumberInput.numberInputID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: numberInputRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: numberInputRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: numberInputRef.parentFormID,
				numberInputID: numberInputRef.numberInputID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/numberInput/setPermissions",params,function(updatedNumberInput) {
				configureNumberInputButtonSpinner($numberInput,updatedNumberInput)
				configureNumberInputClearValueButton($numberInput, updatedNumberInput)			
				setContainerComponentInfo($numberInput,updatedNumberInput,updatedNumberInput.numberInputID)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: numberInputRef.parentFormID,
		componentID: numberInputRef.numberInputID,
		componentLabel: 'number input',
		$componentContainer: $numberInput
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	var spinnerParams = {
		initialShowSpinner: numberInputRef.properties.showValueSpinner,
		initialStepSize: numberInputRef.properties.valueSpinnerStepSize,
		setStepSize: function(stepSize) {
			var params = {
				parentFormID: numberInputRef.parentFormID,
				numberInputID: numberInputRef.numberInputID,
				valueSpinnerStepSize: stepSize
			}
			jsonAPIRequest("frm/numberInput/setSpinnerStepSize",params,function(updatedNumberInput) {
				setContainerComponentInfo($numberInput,updatedNumberInput,updatedNumberInput.numberInputID)
			})
		},
		setShowSpinner: function(showSpinner) {
			var params = {
				parentFormID: numberInputRef.parentFormID,
				numberInputID: numberInputRef.numberInputID,
				showValueSpinner: showSpinner
			}
			jsonAPIRequest("frm/numberInput/setShowSpinner",params,function(updatedNumberInput) {
				configureNumberInputButtonSpinner($numberInput,updatedNumberInput)
				setContainerComponentInfo($numberInput,updatedNumberInput,updatedNumberInput.numberInputID)
			})
		}
	}
	initSpinnerButtonProps(spinnerParams)


	var clearValueParams = {
		initialVal: numberInputRef.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentFormID: numberInputRef.parentFormID,
				numberInputID: numberInputRef.numberInputID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("frm/numberInput/setClearValueSupported",formatParams,function(updatedNumberInput) {
				setContainerComponentInfo($numberInput,updatedNumberInput,updatedNumberInput.numberInputID)
				configureNumberInputClearValueButton($numberInput, updatedNumberInput)			
			})
		}
	}
	initClearValueProps(clearValueParams)

	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#numberInputProps')
		
	toggleFormulaEditorForField(numberInputRef.properties.fieldID)
		
}