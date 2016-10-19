function initRatingViewProperties(ratingRef) {
	console.log("Init checkbox properties panel")
	
	var elemPrefix = "rating_"
	var componentLink = ratingRef.properties.componentLink
	var currRecordRef = currRecordSet.currRecordRef()	

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)

		
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#ratingViewProps')
	
	
}