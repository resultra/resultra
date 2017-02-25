


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
			$imageInnerContainer.empty()
			for(var refIndex = 0; refIndex < attachRefs.length; refIndex++) {
				
				var attachRef = attachRefs[refIndex]
				
				var $imageContainer = imageGalleryThumbnailContainer(attachRef.url)
				$imageContainer.data("attachRef",attachRef)
				$imageInnerContainer.append($imageContainer)
				
			}
			$imageInnerContainer.magnificPopup({
				delegate: 'div.attachGalleryThumbnailContainer',
				type: 'image',
				gallery: { enabled:true },
				image: {
					tError: '<a href="%url%">The image #%curr%</a> could not be loaded.',
					titleSrc: function(item) {
						var $imageContainer = $(item.el)
						var attachRef = $imageContainer.data("attachRef")
						
						return attachmentTitleAndCaptionHTML(attachRef)
					}
				}
			})
		})
	
	} else {
		// There's no value in the current record for this field, so clear the value in the container
		$imageInnerContainer.empty()
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
	
	function saveRecordUpdateWithAttachmentListChanges(updatedAttachmentList) {
		var currRecordRef = recordProxy.getRecordFunc()
				
		var recordUpdateParams = {
			parentDatabaseID:currRecordRef.parentDatabaseID,
			fieldID: imageFieldID, 
			recordID: currRecordRef.recordID,
			changeSetID: recordProxy.changeSetID,
			valueFormatContext: "image",
			valueFormatFormat: "general",
			attachments: updatedAttachmentList }
		console.log("Attachment: Setting file field value: " + JSON.stringify(recordUpdateParams))
		jsonAPIRequest("recordUpdate/setFileFieldValue", recordUpdateParams, function(updatedRecord) {
			console.log("Attachment: Done uploading file: updated record ref = " + JSON.stringify(updatedRecord))
			recordProxy.updateRecordFunc(updatedRecord)
		})
		
	}
			
	var $manageAttachmentsButton = manageAttachmentsButtonFromImageComponentContainer($imageContainer)
		initButtonControlClickHandler($manageAttachmentsButton,function() {
			
			var currRecordRef = recordProxy.getRecordFunc()
		
			// Start with the current file list, then append the newly uploaded attachments.
			var attachmentList = []
			if(currRecordRef.fieldValues.hasOwnProperty(imageFieldID)) {
				attachmentList = currRecordRef.fieldValues[imageFieldID].attachments
			}
			
			var manageAttachmentParams = {
				parentDatabaseID: componentContext.databaseID,
				attachmentList: attachmentList,
				changeAttachmentsCallback: saveRecordUpdateWithAttachmentListChanges
			}
			openManageAttachmentsDialog(manageAttachmentParams)
		})
		
}