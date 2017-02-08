function initTextBoxViewProperties(componentSelectionParams) {
	console.log("Init text box properties panel")

	var textBoxRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()		

	var elemPrefix = "textBox_"
	var componentLink = textBoxRef.properties.componentLink

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)

	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#textBoxViewProps')
	
	
}