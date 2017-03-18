
function loadCheckboxProperties($container, checkBoxRef) {
	console.log("Loading checkbox properties")
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#checkBoxProps')
		
	toggleFormulaEditorForField(checkBoxRef.properties.fieldID)
	
	var $colorSchemeSelection = $('#adminCheckboxComponentColorSchemeSelection')
	$colorSchemeSelection.val(checkBoxRef.properties.colorScheme)
	initSelectControlChangeHandler($colorSchemeSelection,function(newColorScheme) {
		
		var colorSchemeParams = {
			parentFormID: checkBoxRef.parentFormID,
			checkBoxID: checkBoxRef.checkBoxID,
			colorScheme: newColorScheme
		}
		console.log("Setting new color scheme: " + JSON.stringify(colorSchemeParams))
		
		jsonAPIRequest("frm/checkBox/setColorScheme",colorSchemeParams,function(updatedCheckboxRef) {
			setContainerComponentInfo($container,updatedCheckboxRef,updatedCheckboxRef.checkBoxID)		
		})
		
	})
	
	initCheckboxChangeHandler('#adminCheckboxComponentStrikethrough', 
				checkBoxRef.properties.strikethroughCompleted, function (newVal) {
					
		var strikethroughParams = {
			parentFormID: checkBoxRef.parentFormID,
			checkBoxID: checkBoxRef.checkBoxID,
			strikethroughCompleted: newVal
		}
		console.log("Setting new strikethrough settings: " + JSON.stringify(strikethroughParams))

		jsonAPIRequest("frm/checkBox/setStrikethrough",strikethroughParams,function(updatedCheckboxRef) {
			setContainerComponentInfo($container,updatedCheckboxRef,updatedCheckboxRef.checkBoxID)		
		})
	})
	
	
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
	
	
}