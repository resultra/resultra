


function loadRecordIntoImage(imageElem, recordRef) {
	
	console.log("loadRecordIntoImage: loading record into html editor: " + JSON.stringify(recordRef))
	
	var imageObjectRef = imageElem.data("objectRef")
	var imageContainerID = imageObjectRef.imageID
	
	var $imageInnerContainer = imageInnerContainerFromImageComponentContainer(imageElem)
	
	function initImageContainer(imageURL) {
		
	}
	
	var imageFieldID = imageObjectRef.properties.fieldID

	console.log("loadRecordIntoImage: Field ID to load data:" + imageFieldID)

	// Populate the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(imageFieldID)) {
	
		var cloudFileName = recordRef.fieldValues[imageFieldID].cloudName
		// If record has a value for the current container's associated field ID,
		// retrieve an URL for the image and add it to the container.
		var getUrlParams = { 
			parentDatabaseID:recordRef.parentDatabaseID,
			recordID: recordRef.recordID, 
			fieldID: imageFieldID,
			cloudFileName: cloudFileName }
		jsonAPIRequest("record/getFieldValUrl", getUrlParams, function(urlResp) {
				
			$imageInnerContainer.html(imageLinkHTML(imageContainerID,urlResp.url));
			var linkID = imageLinkIDFromContainerElemID(imageContainerID)
			$('#'+linkID).magnificPopup({type:'image'})
		})
	
	} else {
		// There's no value in the current record for this field, so clear the value in the container
		$imageInnerContainer.html('')
	}	
		
}


function initImageRecordEditBehavior($imageContainer, componentContext,recordProxy,imageObjectRef) {
	
	var imageContainerID = imageObjectRef.imageID

	console.log("initImageRecordEditBehavior: container ID =  " +imageContainerID)
	
	$imageContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoImage
	})		
	
	var imageFieldID = imageObjectRef.properties.fieldID
	
	var uploadImageURL = "/api/recordUpdate/uploadFileToFieldValue"
		
	var $imageUploadInput = imageUploadInputFromImageComponentContainer($imageContainer)
	$imageUploadInput.fileupload({
	        dataType: 'json',
			autoUpload:true,
			maxNumberOfFiles:1,	
			paramName: "uploadFile",
			url:uploadImageURL,
			// paramName corresponds to the name given to the file when it is sent to the server. 
			// This name needs to match the name given to the FormFile() function on the server.
	        done: function (e, data) {
				console.log("File upload request done.")
	            $.each(data.result.files, function (index, file) {
					
					console.log("Done uploading file: " + file.name + " url = " + file.url)
					
					var $fileNameLabel = fileNameLabelFromImageComponentContainer($imageContainer)		
					$fileNameLabel.text(file.name)
					
					var $imageInnerContainer = imageInnerContainerFromImageComponentContainer($imageContainer)
										
					$imageInnerContainer.html(imageLinkHTML(imageContainerID,file.url));
					
					var linkID = imageLinkIDFromContainerElemID(imageContainerID)
					$('#'+linkID).magnificPopup({type:'image'})
					
					console.log("Done uploading file: updated record ref = " + JSON.stringify(file.updatedRecord))
					recordProxy.updateRecordFunc(file.updatedRecord)
										
	            })
	        },
			
	    });
		
		// Wait until actually starting the upload to initialize the form data with the record ID and
		// field ID. The reason this can't happen at the same time as the initial upload button initialization is
		// that the records haven't been loaded when the initial initialization takes place, so the current
		// record is unknown.
		$imageUploadInput.bind('fileuploadsubmit', function (e, data) {
		    // The example input, doesn't have to be part of the upload form:
			
			var fileUploadParams = {}

			var currRecordRef = recordProxy.getRecordFunc()
							
			fileUploadParams = {
				parentDatabaseID:currRecordRef.parentDatabaseID,
				fieldID: imageFieldID, 
				recordID: currRecordRef.recordID,
				changeSetID: recordProxy.changeSetID,
				valueFormatContext: "image",
				valueFormatFormat: "general"}
				
			console.log("Starting file upload: params = " + JSON.stringify(fileUploadParams))
			
		    data.formData = fileUploadParams
		});
		
}