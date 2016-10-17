function initCheckBoxViewProperties(checkboxRef) {
	console.log("Init checkbox properties panel")
	
	var elemPrefix = "checkbox_"
	
	initFormComponentTimelinePane(elemPrefix)
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#checkBoxViewProps')
	
	
}