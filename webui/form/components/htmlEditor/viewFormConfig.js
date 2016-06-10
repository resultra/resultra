function loadRecordIntoHtmlEditor(htmlEditorElem, recordRef) {
	
	console.log("loadRecordIntoHtmlEditor: loading record into html editor: " + JSON.stringify(recordRef))
	
	var htmlEditorObjectRef = htmlEditorElem.data("objectRef")
	var htmlEditorFieldID = htmlEditorObjectRef.properties.fieldID
	
	console.log("loadRecordIntoHtmlEditor: Field ID to load data:" + htmlEditorFieldID)


	var htmlEditorContainerID = htmlEditorObjectRef.htmlEditorID
	var htmlEditorInputID = htmlInputIDFromContainerElemID(htmlEditorContainerID)
	var htmlEditorInputSelector = '#'+htmlEditorInputID
	
	// The editor is stored in alongsied the editable DIV, see 
	// initHtmlEditorRecordEditBehavior below.
	var editor = $('#'+htmlEditorContainerID).data("htmlEditor")


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


function initHtmlEditorRecordEditBehavior(htmlEditorObjectRef) {
	
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
		
		currRecordRef = currRecordSet.currRecordRef()
		
		var setRecordValParams = { 
			parentTableID:viewFormContext.tableID,
			recordID:currRecordRef.recordID, 
			fieldID:objectRef.properties.fieldID, 
			value:inputVal }
		
		console.log("Setting date value: " + JSON.stringify(setRecordValParams))
		
		jsonAPIRequest("recordUpdate/setLongTextFieldValue",setRecordValParams,function(updatedRecordRef) {
			
			// After updating the record, the local cache of records in currentRecordSet will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			currRecordSet.updateRecordRef(updatedRecordRef)
			// After changing the value, some of the calculated fields may have changed. For this
			// reason, it is necessary to reload the record into the layout/form, so the most
			// up to date values will be displayed.
			loadCurrRecordIntoLayout()
		}) // set record's text field value
		
	});	
	
}