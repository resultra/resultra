function initSelectionViewProperties(componentSelectionParams) {
	console.log("Init selection properties panel")

	var selectionRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()	

	var elemPrefix = "selection_"
	var componentLink = selectionRef.properties.componentLink

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)

	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#selectionViewProps')
	
	
}