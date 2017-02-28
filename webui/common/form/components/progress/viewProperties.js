function initProgressViewProperties(componentSelectionParams) {
	console.log("Init progress indicator properties panel")
		
	var progressRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()	

	var elemPrefix = "progress_"		
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#progressProps')
	
	
}