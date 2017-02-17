function loadSelectionProperties(selectionRef) {
	console.log("loading selection properties")
	
	
	initSelectableValuesProperties(selectionRef)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#selectionProps')
		
	toggleFormulaEditorForField(selectionRef.properties.fieldID)
		
}