function initRatingViewProperties(componentSelectionParams) {
	console.log("Init checkbox properties panel")
	
	var ratingRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()		

	var elemPrefix = "rating_"

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		fieldID: ratingRef.properties.fieldID
	}
	initFormComponentTimelinePane(timelineParams)

	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#ratingViewProps')
	
	
}