




function initImageRecordEditBehavior($imageContainer, componentContext,recordProxy,imageObjectRef) {
	
	var imageContainerID = imageObjectRef.imageID
	var $imageInnerContainer = imageInnerContainerFromImageComponentContainer($imageContainer)

	console.log("initImageRecordEditBehavior: container ID =  " +imageContainerID)
	
	
	var validateInput = function(validationCompleteCallback) {
		
		if(formComponentIsReadOnly(imageObjectRef.properties.permissions)) {
			validationCompleteCallback(true)
			return
		}
		var currentAttachmentIDs = getCurrentlyDisplayedAttachmentList()
		var validationParams = {
			parentFormID: imageObjectRef.parentFormID,
			imageID: imageObjectRef.imageID,
			attachments: currentAttachmentIDs
		}
		jsonAPIRequest("frm/image/validateInput", validationParams, function(validationResult) {
			if (validationResult.validationSucceeded) {
				$imageContainer.popover('destroy')
				validationCompleteCallback(true)
			} else {
				$imageContainer.popover({
					html: 'true',
					content: function() { return escapeHTML(validationResult.errorMsg) },
					trigger: 'manual',
					placement: 'auto left'
				})
				$imageContainer.popover('show')
				validationCompleteCallback(false)
			}
			
		})	
		
	}
	
	var getCurrentlyDisplayedAttachmentList = function() {
		var currentAttachmentIDs = []
		$imageInnerContainer.find(".attachGalleryThumbnailContainer").each(function() {
			var attachRef = $(this).data("attachRef")
			currentAttachmentIDs.push(attachRef.attachmentInfo.attachmentID)
		})
		return currentAttachmentIDs
	}

	function loadRecordIntoImage(imageElem, recordRef) {
	
		console.log("loadRecordIntoImage: loading record into html editor: " + JSON.stringify(recordRef))
	
		var imageObjectRef = imageElem.data("objectRef")
		var imageContainerID = imageObjectRef.imageID
		
		function initImageContainer(imageURL) {
		
		}
	
		var imageFieldID = imageObjectRef.properties.fieldID

		console.log("loadRecordIntoImage: Field ID to load data:" + imageFieldID)
	
		var componentIsReadOnly = formComponentIsReadOnly(imageObjectRef.properties.permissions)
	
		if(componentIsReadOnly) {
			imageElem.find(".imageComponentManageAttachmentsButtton").hide()
			imageElem.find(".attachmentComponentAddLinkButton").hide()
		} else {
			imageElem.find(".imageComponentManageAttachmentsButtton").show()
			imageElem.find(".attachmentComponentAddLinkButton").show()
		}
	
	
		function saveRecordUpdateWithCurrentlyDisplayedAttachmentList() {
			
			validateInput(function(inputIsValid) {
				if(inputIsValid) {
					// Build an up to date list of the currently displayed attachments from attachments displayed in
					// the current gallery.
		
					var currentAttachmentIDs = getCurrentlyDisplayedAttachmentList()
		
					console.log("Saving updated attachment list: " + JSON.stringify(currentAttachmentIDs))
		
		
					// The record proxy is saved as part of initialization.
					// TODO - Pass the record proxy into the load record functions.
					var recordProxy = imageElem.data("viewFormConfig").recordProxy
					var currRecordRef = recordProxy.getRecordFunc()
				
					var recordUpdateParams = {
						parentDatabaseID:currRecordRef.parentDatabaseID,
						fieldID: imageFieldID, 
						recordID: currRecordRef.recordID,
						changeSetID: recordProxy.changeSetID,
						valueFormatContext: "image",
						valueFormatFormat: "general",
						attachments: currentAttachmentIDs }
					console.log("Attachment: Setting file field value: " + JSON.stringify(recordUpdateParams))
					jsonAPIRequest("recordUpdate/setFileFieldValue", recordUpdateParams, function(updatedRecord) {
						console.log("Attachment: Done uploading file: updated record ref = " + JSON.stringify(updatedRecord))
						recordProxy.updateRecordFunc(updatedRecord)
					})
				}
			})
		
		
		}
	

		// Populate the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(imageFieldID)) {
		
			var fieldVal = recordRef.fieldValues[imageFieldID]
		
			var getRefParams = { attachmentIDs: fieldVal.attachments }
			jsonAPIRequest("attachment/getReferences", getRefParams, function(attachRefs) {
				$imageInnerContainer.empty()
				for(var refIndex = 0; refIndex < attachRefs.length; refIndex++) {
				
					var attachRef = attachRefs[refIndex]
								
					var $thumbnailContainer = attachmentGalleryThumbnailContainer(attachRef,
									saveRecordUpdateWithCurrentlyDisplayedAttachmentList,componentIsReadOnly)
					$thumbnailContainer.data("attachRef",attachRef)
					$imageInnerContainer.append($thumbnailContainer)
				
				}
				initAttachmentContainerPopupGallery($imageInnerContainer)
			})
	
		} else {
			// There's no value in the current record for this field, so clear the value in the container
			$imageInnerContainer.empty()
		}	
		
	}

	
	function initAttachmentEditBehavior() {
		initAttachmentFormComponentViewModeGeometry($imageContainer,imageObjectRef)
	
		var imageFieldID = imageObjectRef.properties.fieldID
	
		var uploadImageURL = "/api/attachment/upload"
	
		var uploadedFileResults = []
		
		var $imageUploadInput = imageUploadInputFromImageComponentContainer($imageContainer)
	
		function saveRecordUpdateWithAttachmentListAdditions(newAttachmentList) {
			var currRecordRef = recordProxy.getRecordFunc()
		
			var updatedAttachmentList = []
			if(currRecordRef.fieldValues.hasOwnProperty(imageFieldID)) {
				// If there are existing attachments, merge with the list of existing attachments.
				// Otherwise, use the new list of attachments to initially set the attachment list.
				updatedAttachmentList = currRecordRef.fieldValues[imageFieldID].attachments.slice(0)
			}
			updatedAttachmentList = $.merge(updatedAttachmentList,newAttachmentList)
		
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
					addAttachmentsCallback: saveRecordUpdateWithAttachmentListAdditions
				}
				openAddAttachmentsDialog(manageAttachmentParams)
			})
		
		var $addLinkButton = addLinkButtonFromAttachmentComponentContainer($imageContainer)
		initButtonControlClickHandler($addLinkButton,function() {
			console.log("Add URL Link button clicked")
			var attachLinkParams = {
				parentDatabaseID: componentContext.databaseID,
				addLinkCallback: function(newAttachmentID) {
					var newAttachList = [newAttachmentID]
					saveRecordUpdateWithAttachmentListAdditions(newAttachList)
				}
			}
			openAttachLinkDialog(attachLinkParams)
		})
		
	}
	initAttachmentEditBehavior()
	
	$imageContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoImage,
		recordProxy: recordProxy,
		validateValue: validateInput
	})
	
		
}