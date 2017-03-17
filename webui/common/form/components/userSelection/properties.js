function loadUserSelectionProperties($userSelection,userSelectionRef) {
	console.log("Loading user selection properties")
	
	var elemPrefix = "userSelection_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for user selection")
		var formatParams = {
			parentFormID: userSelectionRef.parentFormID,
			userSelectionID: userSelectionRef.userSelectionID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/userSelection/setLabelFormat", formatParams, function(updatedUserSelection) {
			setUserSelectionComponentLabel($userSelection,updatedUserSelection)
			setContainerComponentInfo($userSelection,updatedUserSelection,updatedUserSelection.userSelectionID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: userSelectionRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#userSelectionProps')
		
	toggleFormulaEditorForField(userSelectionRef.properties.fieldID)
	
}