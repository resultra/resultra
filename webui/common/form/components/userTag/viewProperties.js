function initUserTagViewProperties(componentSelectionParams) {
	
	console.log("Init user selection properties panel")
	
	var userTagRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()

	var elemPrefix = "userTag_"

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		fieldID: userTagRef.properties.fieldID
	}
	initFormComponentTimelinePane(timelineParams)

	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#userTagViewProps')
	
	
}