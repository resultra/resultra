function initTextBoxViewProperties(componentSelectionParams) {
	console.log("Init text box properties panel")

	var textBoxRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()		

	var elemPrefix = "textBox_"

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		fieldID: textBoxRef.properties.fieldID
	}
	initFormComponentTimelinePane(timelineParams)

	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#textBoxViewProps')
	
	
}