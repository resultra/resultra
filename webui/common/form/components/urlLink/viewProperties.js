function initUrlLinkViewProperties(componentSelectionParams) {
	console.log("Init text box properties panel")

	var urlLinkRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()		

	var elemPrefix = "urlLink_"
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#urlLinkViewProps')
	
	
}