
function loadCheckboxProperties() {
	console.log("Loading checkbox properties")
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#checkBoxProps')
	
	
	$( "#checkBoxProps" ).accordion();
	
}