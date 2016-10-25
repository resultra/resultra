function initDatePickerViewProperties(datePickerRef) {
	console.log("Init date picker properties panel")
	
	var elemPrefix = "datePicker_"
	var componentLink = datePickerRef.properties.componentLink
	var currRecordRef = currRecordSet.currRecordRef()	

	var timelineParams = {
		elemPrefix: elemPrefix,
		tableID: viewFormContext.tableID,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)
	
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#datePickerViewProps')
	
	
}