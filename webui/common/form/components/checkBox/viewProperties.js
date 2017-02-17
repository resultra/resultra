function initCheckBoxViewProperties(componentSelectionParams) {
	console.log("Init checkbox properties panel")
		
	var checkboxRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()	

	var elemPrefix = "checkbox_"

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		fieldID: checkboxRef.properties.fieldID
	}
	initFormComponentTimelinePane(timelineParams)

		
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#checkBoxViewProps')
	
	
}