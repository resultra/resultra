function initUserSelectionViewProperties(userSelectionRef) {
	console.log("Init user selection properties panel")
	
	var elemPrefix = "userSelection_"
	var componentLink = userSelectionRef.properties.componentLink
	var currRecordRef = currRecordSet.currRecordRef()	

	var timelineParams = {
		elemPrefix: elemPrefix,
		tableID: viewFormContext.tableID,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)

		
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#userSelectionViewProps')
	
	
}