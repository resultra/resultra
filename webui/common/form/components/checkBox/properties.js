
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
	
	
}