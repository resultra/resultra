
function loadToggleProperties($container, toggleRef) {
	console.log("Loading toggle properties")
		
	var $colorSchemeSelection = $('#adminToggleComponentColorSchemeSelection')
	$colorSchemeSelection.val(toggleRef.properties.colorScheme)
	initSelectControlChangeHandler($colorSchemeSelection,function(newColorScheme) {
		
		var colorSchemeParams = {
			parentFormID: toggleRef.parentFormID,
			toggleID: toggleRef.toggleID,
			colorScheme: newColorScheme
		}
		console.log("Setting new color scheme: " + JSON.stringify(colorSchemeParams))
		
		jsonAPIRequest("frm/toggle/setColorScheme",colorSchemeParams,function(updatedToggleRef) {
			setContainerComponentInfo($container,updatedToggleRef,updatedToggleRef.toggleID)		
		})
		
	})
	

	initToggleChangeHandler('#adminToggleComponentValidationRequired', 
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
			setCheckBoxComponentLabel($container,updatedToggleRef)
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