function loadCommentComponentProperties(commentRef) {
	console.log("Loading comment component properties")
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#commentComponentProps')
	
	toggleFormulaEditorForField(commentRef.properties.fieldID)
	
}