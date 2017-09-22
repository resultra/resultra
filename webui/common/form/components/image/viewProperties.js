function initImageViewProperties(componentSelectionParams) {
	console.log("Init text box properties panel")

	var imageRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()		

	var elemPrefix = "image_"
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#imageViewProps')
	
	
}