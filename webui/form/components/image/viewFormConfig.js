


function loadRecordIntoImage(imageElem, recordRef) {
	
	console.log("loadRecordIntoImage: loading record into html editor: " + JSON.stringify(recordRef))
	
	var imageObjectRef = imageElem.data("objectRef")
	var imageFieldID = imageObjectRef.fieldID
	
	console.log("loadRecordIntoImage: Field ID to load data:" + imageFieldID)

	var imageContainerID = imageObjectRef.imageID
	var imageDivID = imageIDFromContainerElemID(imageContainerID)
	var imageDivIDSelector = '#' + imageDivID


	// Populate the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(imageFieldID)) {
		
		// If record has a value for the current container's associated field ID,
		// retrieve an URL for the image and add it to the container.
		var getUrlParams = { recordID: recordRef.recordID, fieldID: imageFieldID }
		jsonAPIRequest("record/getFieldValUrl", getUrlParams, function(urlResp) {
					
			$(imageDivIDSelector).html(imageLinkHTML(imageContainerID,urlResp.url));
			var linkID = imageLinkIDFromContainerElemID(imageContainerID)
			$('#'+linkID).magnificPopup({type:'image'})
		})
		
	} else {
		// There's no value in the current record for this field, so clear the value in the container
		$(imageDivIDSelector).html('')
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
	var imageUploadID = imageUploadInputIDFromContainerElemID(imageContainerID)
		
	$('#'+imageUploadID).fileupload({
	        dataType: 'json',
			autoUpload:true,
			maxNumberOfFiles:1,	
			paramName: "uploadFile",
			url:"/api/recordUpdate/uploadFileToFieldValue",
			// paramName corresponds to the name given to the file when it is sent to the server. 
			// This name needs to match the name given to the FormFile() function on the server.
	        done: function (e, data) {
				console.log("File upload request done.")
	            $.each(data.result.files, function (index, file) {
					
					console.log("Done uploading file: " + file.name + " url = " + file.url)
					
					var fileNameLabelID = fileNameLabelFromContainerElemID(imageObjectRef.imageID)		
					$('#'+fileNameLabelID).text(file.name)
					var imageDivID = imageIDFromContainerElemID(imageObjectRef.imageID)
					var imageDivIDSelector = '#' + imageDivID
					
					$(imageDivIDSelector).html(imageLinkHTML(imageContainerID,file.url));
					var linkID = imageLinkIDFromContainerElemID(imageContainerID)
					$('#'+linkID).magnificPopup({type:'image'})
					
					// After uploading the file, the local cache of records in currentRecordSet will
					// be out of date. So after updating the record on the server, the locally cached
					// version of the record also needs to be updated. However, unlike other field
					// types, there is no need to refresh all the fields in the record, since 
					// calculated fields' calculations don't (yet) occur based upon files.
					currRecordSet.updateRecordRef(file.updatedRecordRef)
					
	            })
	        },
			
	    });
		
		// Wait until actually starting the upload to initialize the form data with the record ID and
		// field ID. The reason this can't happen at the same time as the initial upload button initialization is
		// that the records haven't been loaded when the initial initialization takes place, so the current
		// record is unknown.
		$('#'+imageUploadID).bind('fileuploadsubmit', function (e, data) {
		    // The example input, doesn't have to be part of the upload form:
			currRecordRef = currRecordSet.currRecordRef()
		    data.formData = { fieldID: imageObjectRef.fieldID, recordID: currRecordRef.recordID }
		});
		
}