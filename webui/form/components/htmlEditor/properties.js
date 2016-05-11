
function loadHtmlEditorProperties(htmlEditorRef) {
	console.log("Loading html editor properties")
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#htmlEditorProps')
	
	
	$( "#htmlEditorProps" ).accordion();
	
	toggleFormulaEditorForField(htmlEditorRef.fieldRef)
	
}