function initTextBoxViewProperties(textBoxRef) {
	console.log("Init text box properties panel")

	var elemPrefix = "textBox_"
	var componentLink = textBoxRef.properties.componentLink
	var currRecordRef = currRecordSet.currRecordRef()	

	var timelineParams = {
		elemPrefix: elemPrefix,
		tableID: viewListContext.tableID,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)

	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#textBoxViewProps')
	
	
}