function loadUserSelectionProperties(userSelectionRef) {
	console.log("Loading user selection properties")
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#userSelectionProps')
		
	toggleFormulaEditorForComponent(userSelectionRef.properties.componentLink)
	
}