
function loadImageProperties(imageRef) {
	console.log("Loading html editor properties")
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#imageProps')

	toggleFormulaEditorForField(imageRef.fieldRef)
	
}