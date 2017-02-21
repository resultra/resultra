


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
		
		var fieldVal = recordRef.fieldValues[imageFieldID]
		
		var fileVals = fieldVal.files
		
		// Populate the image component container with thumbnail images of the images.
		// TODO - Transition to use a gallery or slideshow instead.
		for (var currFileIndex = 0; currFileIndex < fileVals.length; currFileIndex++) {
			var currFileVal = fileVals[currFileIndex]
			var getUrlParams = { 
				parentDatabaseID:recordRef.parentDatabaseID,
				recordID: recordRef.recordID, 
				fieldID: imageFieldID,
				cloudFileName: currFileVal.cloudName }
			jsonAPIRequest("record/getFieldValUrl", getUrlParams, function(urlResp) {
				var $imageContainer = $(imageLinkHTML(imageContainerID,urlResp.url))
				$imageContainer.magnificPopup({type:'image'})
				$imageInnerContainer.append($imageContainer)
			})
		}
	
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
	
	var uploadImageURL = "/api/attachment/upload"
	
	var uploadedFileResults = []
		
	var $imageUploadInput = imageUploadInputFromImageComponentContainer($imageContainer)
	$imageUploadInput.fileupload({
	        dataType: 'json',
			// paramName corresponds to the name given to the file when it is sent to the server. 
			// This name needs to match the name given to the FormFile() function on the server.
		paramName: "uploadFile",
			url:uploadImageURL,
		
		start: function(e) {
			uploadedFileResults = []
			console.log("Attachment: Starting file upload operation")
		},
		
			// The done callback is invoked every time an individual file complete's its upload, not when the 
			// entire upload operation is done. This is the default behavior, unless the 'singleFileUploads' option
			// is set to false, in which case multiple files may be uploaded.
	        done: function (e, data) {
				console.log("Attachment:  upload request done.")
	            $.each(data.result.files, function (index, file) {
					
					console.log("Attachment: Done uploading file: " + file.name + " url = " + file.url)
					
					var $fileNameLabel = fileNameLabelFromImageComponentContainer($imageContainer)		
					$fileNameLabel.text(file.name)
					
					var $imageInnerContainer = imageInnerContainerFromImageComponentContainer($imageContainer)
										
					$imageInnerContainer.html(imageLinkHTML(imageContainerID,file.url));
					
					var linkID = imageLinkIDFromContainerElemID(imageContainerID)
					$('#'+linkID).magnificPopup({type:'image'})
					
					
					uploadedFileResults.push(file.attachmentInfo)
										
	            })
	        },
			
			progress: function(e,data) {
				 var uploadProgress = parseInt(data.loaded / data.total * 100, 10);
				 console.log("Attachment: File upload progress: " + uploadProgress)
			},

			progressall: function(e,data) {
				 var uploadProgress = parseInt(data.loaded / data.total * 100, 10);
				 console.log("Attachment: All file upload progress: " + uploadProgress)
			},

			
			stop: function(e) {
				console.log("Attachment: All file uploads complete: " + JSON.stringify(uploadedFileResults))
				
				var currRecordRef = recordProxy.getRecordFunc()
				
				// Start with the current file list, then append the newly uploaded attachments.
				var fileValList = []
				if(currRecordRef.fieldValues.hasOwnProperty(imageFieldID)) {
					fileValList = currRecordRef.fieldValues[imageFieldID].files
				}
				for(var currFileIndex = 0; currFileIndex < uploadedFileResults.length; currFileIndex++) {
					var currFileInfo = uploadedFileResults[currFileIndex]
					var currFileVal = {
						cloudName: currFileInfo.cloudFileName,
						origName: currFileInfo.origFileName}
					fileValList.push(currFileVal)
				}
				
				
				var recordUpdateParams = {
					parentDatabaseID:currRecordRef.parentDatabaseID,
					fieldID: imageFieldID, 
					recordID: currRecordRef.recordID,
					changeSetID: recordProxy.changeSetID,
					valueFormatContext: "image",
					valueFormatFormat: "general",
					files: fileValList }
				console.log("Attachment: Setting file field value: " + JSON.stringify(recordUpdateParams))
				jsonAPIRequest("recordUpdate/setFileFieldValue", recordUpdateParams, function(updatedRecord) {
					console.log("Attachment: Done uploading file: updated record ref = " + JSON.stringify(updatedRecord))
					recordProxy.updateRecordFunc(updatedRecord)
				})
					
				uploadedFileResults = []
			}
	    });
		
		// Wait until actually starting the upload to initialize the form data with the record ID and
		// field ID. The reason this can't happen at the same time as the initial upload button initialization is
		// that the records haven't been loaded when the initial initialization takes place, so the current
		// record is unknown.
		$imageUploadInput.bind('fileuploadsubmit', function (e, data) {
		    // The example input, doesn't have to be part of the upload form:
			
			var currRecordRef = recordProxy.getRecordFunc()
			
								
			var fileUploadParams = { parentDatabaseID:currRecordRef.parentDatabaseID }
				
			console.log("Attachment: Starting file upload: params = " + JSON.stringify(fileUploadParams))
			
		    data.formData = fileUploadParams
		});
		
}