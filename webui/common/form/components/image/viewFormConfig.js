


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
		
		var getRefParams = { attachmentIDs: fieldVal.attachments }
		jsonAPIRequest("attachment/getReferences", getRefParams, function(attachRefs) {
			for(var refIndex = 0; refIndex < attachRefs.length; refIndex++) {
				
				var attachRef = attachRefs[refIndex]
				
				var $imageContainer = $(imageLinkHTML(imageContainerID,attachRef.url))
				$imageContainer.magnificPopup({type:'image'})
				$imageInnerContainer.append($imageContainer)
				
			}
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
	
	var uploadImageURL = "/api/attachment/upload"
	
	var uploadedFileResults = []
		
	var $imageUploadInput = imageUploadInputFromImageComponentContainer($imageContainer)
	
	function saveRecordUpdateWithNewAttachments(attachments) {
		var currRecordRef = recordProxy.getRecordFunc()
		
		// Start with the current file list, then append the newly uploaded attachments.
		var attachmentList = []
		if(currRecordRef.fieldValues.hasOwnProperty(imageFieldID)) {
			attachmentList = currRecordRef.fieldValues[imageFieldID].attachments
		}
		for(var currFileIndex = 0; currFileIndex < attachments.length; currFileIndex++) {
			var newAttachment = attachments[currFileIndex]
			attachmentList.push(newAttachment.attachmentID)
		}
		
		
		var recordUpdateParams = {
			parentDatabaseID:currRecordRef.parentDatabaseID,
			fieldID: imageFieldID, 
			recordID: currRecordRef.recordID,
			changeSetID: recordProxy.changeSetID,
			valueFormatContext: "image",
			valueFormatFormat: "general",
			attachments: attachmentList }
		console.log("Attachment: Setting file field value: " + JSON.stringify(recordUpdateParams))
		jsonAPIRequest("recordUpdate/setFileFieldValue", recordUpdateParams, function(updatedRecord) {
			console.log("Attachment: Done uploading file: updated record ref = " + JSON.stringify(updatedRecord))
			recordProxy.updateRecordFunc(updatedRecord)
		})
		
	}
	
	var addAttachmentParams = {
		parentDatabaseID: componentContext.databaseID,
		$addAttachmentInput: $imageUploadInput,
		attachDoneCallback: saveRecordUpdateWithNewAttachments }
	initAddAttachmentControl(addAttachmentParams)
		
}