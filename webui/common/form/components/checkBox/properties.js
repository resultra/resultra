
function loadCheckboxProperties($container, checkBoxRef) {
	console.log("Loading checkbox properties")
	
	
	
	var formatParams = {
		initialColorScheme: checkBoxRef.properties.colorScheme,
		setColorScheme: function(newColorScheme) {
			var colorSchemeParams = {
				parentFormID: checkBoxRef.parentFormID,
				checkBoxID: checkBoxRef.checkBoxID,
				colorScheme: newColorScheme
			}
			console.log("Setting new color scheme: " + JSON.stringify(colorSchemeParams))
	
			jsonAPIRequest("frm/checkBox/setColorScheme",colorSchemeParams,function(updatedCheckboxRef) {
				setContainerComponentInfo($container,updatedCheckboxRef,updatedCheckboxRef.checkBoxID)		
			})
		},
		initialStrikethrough: checkBoxRef.properties.strikethroughCompleted,
		setStrikethrough: function(strikethroughCompleted) {
			var strikethroughParams = {
				parentFormID: checkBoxRef.parentFormID,
				checkBoxID: checkBoxRef.checkBoxID,
				strikethroughCompleted: newVal
			}
			console.log("Setting new strikethrough settings: " + JSON.stringify(strikethroughParams))

			jsonAPIRequest("frm/checkBox/setStrikethrough",strikethroughParams,function(updatedCheckboxRef) {
				setContainerComponentInfo($container,updatedCheckboxRef,updatedCheckboxRef.checkBoxID)		
			})
		}
	}
	initCheckBoxFormatProps(formatParams)


	var validationParams = {
		initialValidationConfig: checkBoxRef.properties.validation,
		setValidation: function(validationConfig) {
			var validationParams = {
				parentFormID: checkBoxRef.parentFormID,
				checkBoxID: checkBoxRef.checkBoxID,
				validation: validationConfig
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("frm/checkBox/setValidation",validationParams,function(updatedCheckboxRef) {
				setContainerComponentInfo($container,updatedCheckboxRef,updatedCheckboxRef.checkBoxID)		
			})
		
		}
	}
	initCheckBoxValidationProps(validationParams)	
	
	var elemPrefix = "checkbox_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: checkBoxRef.parentFormID,
			checkBoxID: checkBoxRef.checkBoxID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/checkBox/setLabelFormat", formatParams, function(updatedCheckboxRef) {
			setCheckBoxComponentLabel($container,updatedCheckboxRef)
			setContainerComponentInfo($container,updatedCheckboxRef,updatedCheckboxRef.checkBoxID)		
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: checkBoxRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: checkBoxRef.parentFormID,
			checkBoxID: checkBoxRef.checkBoxID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/checkBox/setVisibility",params,function(updatedCheckboxRef) {
			setContainerComponentInfo($container,updatedCheckboxRef,updatedCheckboxRef.checkBoxID)		
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: checkBoxRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)

	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: checkBoxRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: checkBoxRef.parentFormID,
				checkBoxID: checkBoxRef.checkBoxID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/checkBox/setPermissions",params,function(updatedCheckboxRef) {
				setContainerComponentInfo($container,updatedCheckboxRef,updatedCheckboxRef.checkBoxID)		
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: checkBoxRef.parentFormID,
		componentID: checkBoxRef.checkBoxID,
		componentLabel: 'check box',
		$componentContainer: $container
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#checkBoxProps')
		
	toggleFormulaEditorForField(checkBoxRef.properties.fieldID)
	
	
}