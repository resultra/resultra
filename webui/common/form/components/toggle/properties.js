
function loadToggleProperties($container, toggleRef) {
	console.log("Loading toggle properties")
	
	function initFormatProperties() {
		
		var $form = $('#adminToggleComponentFormatForm')
		
		var $offLabel = $form.find("input[name=adminToggleOffComponentLabel]")
		$offLabel.val(toggleRef.properties.offLabel)
		var validationOffLabelParams = {
			url: '/api/generic/stringValidation/validateItemLabel',
			data: { label: function() { return $offLabel.val(); } }
		}
		function saveValidatedOffLabel(newLabel) {
			console.log("Toggle format properties: new off label: " + newLabel)
			
			var labelParams = {
				parentFormID: toggleRef.parentFormID,
				toggleID: toggleRef.toggleID,
				label: newLabel
			}
			jsonAPIRequest("frm/toggle/setOffLabel",labelParams,function(updatedToggleRef) {
				reInitToggleComponentControl($container,updatedToggleRef)
				setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)		
			})
		}

		var $onLabel = $form.find("input[name=adminToggleOnComponentLabel]")
		$onLabel.val(toggleRef.properties.onLabel)
		var validationOnLabelParams = {
			url: '/api/generic/stringValidation/validateItemLabel',
			data: { label: function() { return $onLabel.val(); } }
		}
		function saveValidatedOnLabel(newLabel) {
			console.log("Toggle format properties: new on label: " + newLabel)
			var labelParams = {
				parentFormID: toggleRef.parentFormID,
				toggleID: toggleRef.toggleID,
				label: newLabel
			}
			jsonAPIRequest("frm/toggle/setOnLabel",labelParams,function(updatedToggleRef) {
				reInitToggleComponentControl($container,updatedToggleRef)
				setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)		
			})
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
		$toggleOffColorSchemeSelection.val(toggleRef.properties.offColorScheme)
		initSelectControlChangeHandler($toggleOffColorSchemeSelection,function(newColorScheme) {
		
			var colorSchemeParams = {
				parentFormID: toggleRef.parentFormID,
				toggleID: toggleRef.toggleID,
				colorScheme: newColorScheme
			}
			console.log("Setting new color scheme: " + JSON.stringify(colorSchemeParams))
		
			jsonAPIRequest("frm/toggle/setOffColorScheme",colorSchemeParams,function(updatedToggleRef) {
				reInitToggleComponentControl($container,updatedToggleRef)
				setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)		
			})
		
		})
		
		var $toggleOnColorSchemeSelection = $('#adminToggleOnComponentColorSchemeSelection')
		$toggleOnColorSchemeSelection.val(toggleRef.properties.onColorScheme)
		initSelectControlChangeHandler($toggleOnColorSchemeSelection,function(newColorScheme) {
		
			var colorSchemeParams = {
				parentFormID: toggleRef.parentFormID,
				toggleID: toggleRef.toggleID,
				colorScheme: newColorScheme
			}
			console.log("Setting new color scheme: " + JSON.stringify(colorSchemeParams))
		
			jsonAPIRequest("frm/toggle/setOnColorScheme",colorSchemeParams,function(updatedToggleRef) {
				reInitToggleComponentControl($container,updatedToggleRef)
				setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)	
			})
		
		})
		
	}
	initFormatProperties()
	
	
	initCheckboxChangeHandler('#adminToggleComponentValidationRequired', 
				toggleRef.properties.validation.valueRequired, function (newVal) {
				
		var validationProps = {
			valueRequired: newVal
		}		
				
		var validationParams = {
			parentFormID: toggleRef.parentFormID,
			toggleID: toggleRef.toggleID,
			validation: validationProps
		}
		console.log("Setting new validation settings: " + JSON.stringify(validationParams))

		jsonAPIRequest("frm/toggle/setValidation",validationParams,function(updatedToggleRef) {
			setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)		
		})
	})
	
	
	var elemPrefix = "toggle_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: toggleRef.parentFormID,
			toggleID: toggleRef.toggleID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/toggle/setLabelFormat", formatParams, function(updatedToggleRef) {
			setToggleComponentLabel($container,updatedToggleRef)
			setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)		
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: toggleRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: toggleRef.parentFormID,
			toggleID: toggleRef.toggleID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/toggle/setVisibility",params,function(updatedToggleRef) {
			setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)		
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: toggleRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)

	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: toggleRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: toggleRef.parentFormID,
				toggleID: toggleRef.toggleID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/toggle/setPermissions",params,function(updatedToggleRef) {
				setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)		
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: toggleRef.parentFormID,
		componentID: toggleRef.toggleID,
		componentLabel: 'toggle',
		$componentContainer: $container
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#toggleProps')
		
	toggleFormulaEditorForField(toggleRef.properties.fieldID)
	
	
}