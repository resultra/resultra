function initImageViewProperties(componentSelectionParams) {
	console.log("Image properties panel")
	
	var imageRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()		

	var elemPrefix = "image_"
	var componentLink = imageRef.properties.componentLink

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#imageViewProps')
	
	
}