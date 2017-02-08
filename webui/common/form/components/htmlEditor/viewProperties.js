function initHTMLEditorViewProperties(componentSelectionParams) {
	console.log("Init html editor properties panel")

	var htmlEditorRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()	

	var elemPrefix = "htmlEditor_"
	var componentLink = htmlEditorRef.properties.componentLink	

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)

	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#htmlEditorViewProps')
	
	
}