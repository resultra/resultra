function initLabelViewProperties(componentSelectionParams) {
	
	console.log("Init user selection properties panel")
	
	var labelRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()

	var elemPrefix = "label_"

	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#labelViewProps')
	
	
}