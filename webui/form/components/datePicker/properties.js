
function loadDatePickerProperties(datePickerRef) {
	console.log("Loading date picker properties")
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#datePickerProps')
	
	toggleFormulaEditorForField(datePickerRef.fieldRef)
	
}