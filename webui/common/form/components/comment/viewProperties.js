function initCommentViewProperties(componentSelectionParams) {
	console.log("Init comment component properties panel")
	
	var commentRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()
	
	// Comment boxes are not linked to the timeline with a ComponentLink,
	// However, a ComponentLink can be synthesized with just a field ID.
	var elemPrefix = "comment_"
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