function loadTextBoxProperties(textBoxRef) {
	console.log("loading text box properties")
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#textBoxProps')
		
	toggleFormulaEditorForField(textBoxRef.fieldRef)
		
}