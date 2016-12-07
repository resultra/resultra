function initSelectionViewProperties(selectionRef) {
	console.log("Init selection properties panel")

	var elemPrefix = "selection_"
	var componentLink = selectionRef.properties.componentLink
	var currRecordRef = currRecordSet.currRecordRef()	

	var timelineParams = {
		elemPrefix: elemPrefix,
		tableID: viewListContext.tableID,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)

	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#selectionViewProps')
	
	
}