function loadRecordIntoImage(imageElem, recordRef) {
	
	console.log("loadRecordIntoImage: loading record into html editor: " + JSON.stringify(recordRef))
	
	var imageObjectRef = imageElem.data("objectRef")
	var imageFieldID = imageObjectRef.fieldRef.fieldID
	
	console.log("loadRecordIntoImage: Field ID to load data:" + imageFieldID)

	var imageContainerID = imageObjectRef.imageID


	// Populate the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(imageFieldID)) {

		// If record has a value for the current container's associated field ID.
			// TODO - Set the image
	} else {
		// There's no value in the current record for this field, so clear the value in the container
		// TODO - Clear the image
	}	
}


function initImageRecordEditBehavior(imageObjectRef) {
	
	var imageContainerID = imageObjectRef.imageID

	console.log("initImageRecordEditBehavior: container ID =  " +imageContainerID)
	
	var imageContainer = $('#'+imageContainerID)
	imageContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoImage
	})

	// Initialize image uploader plugin
	var imageDropZoneContainerID = imageIDFromContainerElemID(imageContainerID)
	var imageDropZoneSelector = '#' + imageDropZoneContainerID
	console.log("Initializing image drop zone: " + imageDropZoneSelector)
	$(imageDropZoneSelector).dropzone({
		url: "/api/record/uploadFile",
		maxFiles: 1
	})
		

	// TODO - Handle notificactions of new image upload - need to refresh display.
	// e.g. imageContanier.onChange ...
	/*
	imageContainer.on('blur', function(event) {

	    console.log("html editor blur")
		
		// Get the most recent copy of the object reference. It could have changed between
		// initialization time and the time the checkbox was changed.
		var containerID = imageObjectRef.imageID
		var objectRef = getElemObjectRef(containerID)
		
		var imageInputID = htmlInputIDFromContainerElemID(containerID)
		
		var editor = $('#'+containerID).data("image")
		var inputVal = editor.getData();
		
		currRecordRef = currRecordSet.currRecordRef()
		
		var setRecordValParams = { recordID:currRecordRef.recordID, fieldID:objectRef.fieldRef.fieldID, value:inputVal }
		
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
	*/
}