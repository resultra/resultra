function initCommentViewProperties(commentRef) {
	console.log("Init comment component properties panel")
	
	var elemPrefix = "comment_"
	var currRecordRef = currRecordSet.currRecordRef()
	
	// Comment boxes are not linked to the timeline with a ComponentLink,
	// However, a ComponentLink can be synthesized with just a field ID.
	var componentLink = fieldComponentValType(commentRef.properties.fieldID)

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#commentViewProps')
	
	
}