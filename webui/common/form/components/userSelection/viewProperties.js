function initUserSelectionViewProperties(componentSelectionParams) {
	
	console.log("Init user selection properties panel")
	
	var userSelectionRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()

	var elemPrefix = "userSelection_"
	var componentLink = userSelectionRef.properties.componentLink

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)

	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#userSelectionViewProps')
	
	
}