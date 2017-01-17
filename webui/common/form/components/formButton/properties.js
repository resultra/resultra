

function loadFormButtonProperties(buttonRef) {
	
	console.log("Loading button properties")
	
	initHeaderTextProperties(headerRef)
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#formButtonProps')
		
	closeFormulaEditor()
	
}