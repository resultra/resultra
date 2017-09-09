function initFileViewProperties(componentSelectionParams) {
	console.log("Init text box properties panel")

	var fileRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()		

	var elemPrefix = "file_"
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#fileViewProps')
	
	
}