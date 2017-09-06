function initEmailAddrViewProperties(componentSelectionParams) {
	console.log("Init text box properties panel")

	var emailAddrRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()		

	var elemPrefix = "emailAddr_"
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#emailAddrViewProps')
	
	
}