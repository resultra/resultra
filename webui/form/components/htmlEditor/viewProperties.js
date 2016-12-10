function initHTMLEditorViewProperties(htmlEditorRef) {
	console.log("Init html editor properties panel")


	var elemPrefix = "htmlEditor_"
	var componentLink = htmlEditorRef.properties.componentLink
	var currRecordRef = currRecordSet.currRecordRef()	

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		componentLink: componentLink
	}
	initFormComponentTimelinePane(timelineParams)

	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#htmlEditorViewProps')
	
	
}