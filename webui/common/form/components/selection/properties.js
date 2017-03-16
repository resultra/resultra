function loadSelectionProperties($selection,selectionRef) {
	console.log("loading selection properties")
	
	
	initSelectableValuesProperties($selection,selectionRef)
	
	
	var elemPrefix = "selection_"
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: selectionRef.parentFormID,
			selectionID: selectionRef.selectionID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/selection/setLabelFormat", formatParams, function(updatedSelection) {
			setSelectionComponentLabel($selection,updatedSelection)
			setContainerComponentInfo($selection,updatedSelection,updatedSelection.selectionID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: selectionRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#selectionProps')
		
	toggleFormulaEditorForField(selectionRef.properties.fieldID)
		
}