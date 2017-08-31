function initSocialButtonViewProperties(componentSelectionParams) {
	console.log("Init checkbox properties panel")
	
	var socialButtonRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()		

	var elemPrefix = "socialButton_"
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#socialButtonViewProps')
	
	
}