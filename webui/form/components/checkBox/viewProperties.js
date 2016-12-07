function initCheckBoxViewProperties(checkboxRef) {
	console.log("Init checkbox properties panel")
	
	var elemPrefix = "checkbox_"
	var componentLink = checkboxRef.properties.componentLink
	var currRecordRef = currRecordSet.currRecordRef()	

	var timelineParams = {
		elemPrefix: elemPrefix,
		tableID: viewListContext.tableID,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)

		
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#checkBoxViewProps')
	
	
}