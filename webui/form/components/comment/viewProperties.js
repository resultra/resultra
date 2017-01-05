function initCommentViewProperties(commentRef) {
	console.log("Init comment component properties panel")
	
	var elemPrefix = "comment_"
//	var componentLink = datePickerRef.properties.componentLink
	var currRecordRef = currRecordSet.currRecordRef()	

/* TODO - A comment field doesn't have a component link
	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)
	*/
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#commentViewProps')
	
	
}