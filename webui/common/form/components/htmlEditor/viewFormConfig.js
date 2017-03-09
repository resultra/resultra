function loadRecordIntoHtmlEditor($htmlEditor, recordRef) {
	
	console.log("loadRecordIntoHtmlEditor: loading record into html editor: " + JSON.stringify(recordRef))
	
	var htmlEditorObjectRef = $htmlEditor.data("objectRef")
	
	var $htmlEditorInput = htmlInputFromHTMLEditorContainer($htmlEditor)
	
	var htmlEditorFieldID = htmlEditorObjectRef.properties.fieldID
	// Populate the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(htmlEditorFieldID)) {

		// If record has a value for the current container's associated field ID.
		var $fieldVal = $(recordRef.fieldValues[htmlEditorFieldID])
		$fieldVal.find('a').attr("target","_blank")
		$htmlEditorInput.html($fieldVal)		
	
	} else {
		// There's no value in the current record for this field, so clear the value in the container
		$htmlEditorInput.html("")
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
	
	var $htmlEditorInput = htmlInputFromHTMLEditorContainer($htmlEditor)
	
	var $editButton = $htmlEditor.find(".startEditButton")
	initButtonControlClickHandler($editButton,function() {
		console.log("Starting inline editor")
		
		if (!inlineCKEditorEnabled($htmlEditorInput)) {
		    CKEDITOR.disableAutoInline = true;
	
			var editor = enableInlineCKEditor($htmlEditorInput)
		
			editor.on('blur', function(event) {
				var editorInput = editor.getData();
				var currRecordRef = recordProxy.getRecordFunc()
				var htmlEditorFieldID = htmlEditorObjectRef.properties.fieldID
				var inputVal = editor.getData();
			
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
				
				disableInlineCKEditor($htmlEditorInput,editor)
			
				jsonAPIRequest("recordUpdate/setLongTextFieldValue",setRecordValParams,function(updatedRecordRef) {
		
					// After updating the record, the local cache of records in currentRecordSet will
					// be out of date. So after updating the record on the server, the locally cached
					// version of the record also needs to be updated.
					recordProxy.updateRecordFunc(updatedRecordRef)
				}) // set record's text field value
						
			})
		
			$htmlEditorInput.focus()
		}
		
	})
	
}