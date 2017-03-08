function loadRecordIntoHtmlEditor($htmlEditor, recordRef) {
	
	console.log("loadRecordIntoHtmlEditor: loading record into html editor: " + JSON.stringify(recordRef))
	
	var htmlEditorObjectRef = $htmlEditor.data("objectRef")
			
	// The editor is stored in alongside the editable DIV, see 
	// initHtmlEditorRecordEditBehavior below.
	var editor = $htmlEditor.data("htmlEditor")

	var htmlEditorFieldID = htmlEditorObjectRef.properties.fieldID
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

}


function initHtmlEditorRecordEditBehavior($htmlEditor,componentContext,recordProxy,htmlEditorObjectRef) {
	
		
	$htmlEditor.data("viewFormConfig", {
		loadRecord: loadRecordIntoHtmlEditor
	})
	
	$htmlEditor.find(".htmlEditorContent").click(function(e) {
		
		// This is important - if a click hits an object, then stop the propagation of the click
		// to the parent div(s), including the canvas itself. If the parent canvas
		// gets a click, it will deselect all the items (see initObjectCanvasSelectionBehavior)
		e.stopPropagation();
	})
	

    CKEDITOR.disableAutoInline = true;
	
	var $htmlEditorInput = htmlInputFromHTMLEditorContainer($htmlEditor)
	
	var editor = enableInlineCKEditor($htmlEditorInput)
	
	$htmlEditor.data("htmlEditor",editor)

	editor.on('blur', function(event) {
		
	    console.log("html editor blur")
		
		// Get the most recent copy of the object reference. It could have changed between
		// initialization time and the time the checkbox was changed.
		var containerID = htmlEditorObjectRef.htmlEditorID
		var objectRef = getContainerObjectRef($htmlEditor)
		
		var editor = $htmlEditor.data("htmlEditor")
		var inputVal = editor.getData();
		
		var currRecordRef = recordProxy.getRecordFunc()

		
		var htmlEditorFieldID = objectRef.properties.fieldID
		
		var textBoxTextValueFormat = {
			context:"htmlEditor",
			format:"general"
		}
		
		
		var setRecordValParams = { 
			parentDatabaseID:currRecordRef.parentDatabaseID,
			recordID:currRecordRef.recordID,
			changeSetID: recordProxy.changeSetID,
			fieldID:htmlEditorFieldID, 
			value:inputVal,
		valueFormat:textBoxTextValueFormat }
	
		console.log("Setting date value: " + JSON.stringify(setRecordValParams))
	
		jsonAPIRequest("recordUpdate/setLongTextFieldValue",setRecordValParams,function(updatedRecordRef) {
		
			// After updating the record, the local cache of records in currentRecordSet will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			recordProxy.updateRecordFunc(updatedRecordRef)
		}) // set record's text field value
				
	});	
	
}