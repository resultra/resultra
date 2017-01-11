function loadRecordIntoHtmlEditor(htmlEditorElem, recordRef) {
	
	console.log("loadRecordIntoHtmlEditor: loading record into html editor: " + JSON.stringify(recordRef))
	
	var htmlEditorObjectRef = htmlEditorElem.data("objectRef")
	
	
	console.log("loadRecordIntoHtmlEditor: Field ID to load data:" + htmlEditorFieldID)


	var htmlEditorContainerID = htmlEditorObjectRef.htmlEditorID
	var htmlEditorInputID = htmlInputIDFromContainerElemID(htmlEditorContainerID)
	var htmlEditorInputSelector = '#'+htmlEditorInputID
	
	// The editor is stored in alongsied the editable DIV, see 
	// initHtmlEditorRecordEditBehavior below.
	var editor = $('#'+htmlEditorContainerID).data("htmlEditor")


	var componentLink = htmlEditorObjectRef.properties.componentLink
	
	if(componentLink.linkedValType == linkedComponentValTypeField) {
		var htmlEditorFieldID = componentLink.fieldID
		// Populate the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(htmlEditorFieldID)) {

			// If record has a value for the current container's associated field ID.
			var fieldVal = recordRef.fieldValues[htmlEditorFieldID]		
			editor.setData(fieldVal)
		
		} else {
			// There's no value in the current record for this field, so clear the value in the container
			editor.setData("")
		}	
	} else {
		console.log("Globals not yet supported for HTML editor")
	}
	


}


function initHtmlEditorRecordEditBehavior(componentContext,getRecordFunc, updateRecordFunc,htmlEditorObjectRef) {
	
	var htmlEditorContainerID = htmlEditorObjectRef.htmlEditorID
	var htmlEditorInputID = htmlInputIDFromContainerElemID(htmlEditorContainerID)
	
	console.log("initHtmlEditorRecordEditBehavior: container ID =  " +htmlEditorContainerID)
	
	var htmlEditorContainer = $('#'+htmlEditorContainerID)
	htmlEditorContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoHtmlEditor
	})


	var htmlEditorInputID = htmlInputIDFromContainerElemID(htmlEditorContainerID)

    CKEDITOR.disableAutoInline = true;
    var editor = CKEDITOR.inline( htmlEditorInputID );
	htmlEditorContainer.data("htmlEditor",editor)

	editor.on('blur', function(event) {
	    console.log("html editor blur")
		
		// Get the most recent copy of the object reference. It could have changed between
		// initialization time and the time the checkbox was changed.
		var containerID = htmlEditorObjectRef.htmlEditorID
		var objectRef = getElemObjectRef(containerID)
		
		var htmlEditorInputID = htmlInputIDFromContainerElemID(containerID)
		
		var editor = $('#'+containerID).data("htmlEditor")
		var inputVal = editor.getData();
		
		var currRecordRef = getRecordFunc()
		
		var componentLink = objectRef.properties.componentLink
	
		if(componentLink.linkedValType == linkedComponentValTypeField) {
			
			var htmlEditorFieldID = componentLink.fieldID
			
			var textBoxTextValueFormat = {
				context:"htmlEditor",
				format:"general"
			}
			
			
			var setRecordValParams = { 
				parentDatabaseID:viewListContext.databaseID,
				recordID:currRecordRef.recordID, 
				fieldID:htmlEditorFieldID, 
				value:inputVal,
			valueFormat:textBoxTextValueFormat }
		
			console.log("Setting date value: " + JSON.stringify(setRecordValParams))
		
			jsonAPIRequest("recordUpdate/setLongTextFieldValue",setRecordValParams,function(updatedRecordRef) {
			
				// After updating the record, the local cache of records in currentRecordSet will
				// be out of date. So after updating the record on the server, the locally cached
				// version of the record also needs to be updated.
				updateRecordFunc(updatedRecordRef)
			}) // set record's text field value
		} else {
			console.log("HTML editor global values not yet supported")
		}
		
		
	});	
	
}