function initImageViewProperties(imageRef) {
	console.log("Image properties panel")
	
	var elemPrefix = "image_"
	var componentLink = imageRef.properties.componentLink
	var currRecordRef = currRecordSet.currRecordRef()	

	var timelineParams = {
		elemPrefix: elemPrefix,
		tableID: viewFormContext.tableID,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#imageViewProps')
	
	
}