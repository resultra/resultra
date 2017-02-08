function initCheckBoxViewProperties(componentSelectionParams) {
	console.log("Init checkbox properties panel")
		
	var checboxRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()	

	var elemPrefix = "checkbox_"
	var componentLink = checkboxRef.properties.componentLink

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)

		
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#checkBoxViewProps')
	
	
}