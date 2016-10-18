function initCheckBoxViewProperties(checkboxRef) {
	console.log("Init checkbox properties panel")
	
	var elemPrefix = "checkbox_"
	var componentLink = checkboxRef.properties.componentLink
	var currRecordRef = currRecordSet.currRecordRef()	
		
	if(componentLink.linkedValType == linkedComponentValTypeField) {
		var timelineParams = {
			elemPrefix: elemPrefix,
			recordID: currRecordRef.recordID,
			fieldID: componentLink.fieldID
		}
	
		initFormComponentTimelinePane(timelineParams)
	}
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#checkBoxViewProps')
	
	
}