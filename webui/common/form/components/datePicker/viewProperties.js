function initDatePickerViewProperties(componentSelectionParams) {
	console.log("Init date picker properties panel")
	
	var datePickerRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()	

	var elemPrefix = "datePicker_"
	var componentLink = datePickerRef.properties.componentLink
	
	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)
	
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#datePickerViewProps')
	
	
}