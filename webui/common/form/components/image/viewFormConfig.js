




function initAttachmentRecordEditBehavior($imageContainer, componentContext,recordProxy,imageObjectRef,remoteValidationFunc) {
	
	var $imageInnerContainer = imageInnerContainerFromImageComponentContainer($imageContainer)
	
	var validateInput = function(validationCompleteCallback) {
		
		if(formComponentIsReadOnly(imageObjectRef.properties.permissions)) {
			validationCompleteCallback(true)
			return
		}
		var currentAttachmentIDs = getCurrentlyDisplayedAttachmentList()
		remoteValidationFunc(currentAttachmentIDs,function(validationResult) {
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

function initAttachmentFormRecordEditBehavior($imageContainer, componentContext,recordProxy,imageObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		
		var validationParams = {
			parentFormID: imageObjectRef.parentFormID,
			imageID: imageObjectRef.imageID,
			attachments: inputVal
		}
		jsonAPIRequest("frm/image/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	initAttachmentFormComponentViewModeGeometry($imageContainer,imageObjectRef)
	
	initAttachmentRecordEditBehavior($imageContainer, componentContext,recordProxy,imageObjectRef,validateInput)
}


function initAttachmentTableViewRecordEditBehavior($attachContainer, componentContext,recordProxy,attachObjectRef) {
	
	
	function validateInput(inputVal,validationResultCallback) {
		
		var validationParams = {
			parentTableID: attachObjectRef.parentTableID,
			attachmentID: attachObjectRef.attachmentID,
			attachments: inputVal
		}
		jsonAPIRequest("tableView/attachment/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	var currRecordRef = null
	var loadRecordIntoPopupAttachmentEditor = null
	function loadRecordIntoAttachmentEditor($attachContainer, recordRef) {
		currRecordRef = recordRef
		if(loadRecordIntoPopupAttachmentEditor != null) {
			loadRecordIntoPopupAttachmentEditor()
		}
	}
		
	var $attachmentPopupLink = $attachContainer.find(".attachmentEditPopop")
		
	$attachmentPopupLink.popover({
		html: 'true',
		content: function() { return attachmentTableViewPopupEditContainerHTML() },
		trigger: 'click',
		placement: 'auto left',
		container: "body"
	})
	
	$attachmentPopupLink.on('shown.bs.popover', function()
	{
	    //get the actual shown popover
	    var $popover = $(this).data('bs.popover').tip();
		
		// By default the popover takes on the maximum size of it's containing
		// element. Overridding this size allows the size to grow as needed.
		$popover.css("max-width","300px")
		// The max-height needs to be large enough to allow the comment box to
		// expand somewhat.
		$popover.css("max-height","600px")
		
		// Override the popover's default z-index to be less than the popup used to display the
		// attachments and the dialog box used to add new attachments.
		$popover.css("z-index","550")
		
		console.log("Popover html: " + $popover.html())
		
		var $attachmentEditorContainer = $popover.find(".attachmentEditorPopupContainer")
		
		initAttachmentTableCellComponentViewModeGeometry($attachmentEditorContainer)
				
		var $closePopupButton = $attachmentEditorContainer.find(".closeEditorPopup")
		initButtonControlClickHandler($closePopupButton,function() {
			$attachmentPopupLink.popover('hide')
			loadRecordIntoPopupAttachmentEditor = null
		})
		
		
		console.log("Popover html: " + $attachmentEditorContainer.html())
		
		function loadCurrentRecordIntoPopup() {
			if(currRecordRef != null) {
				var viewConfig = $attachmentEditorContainer.data("viewFormConfig")
				viewConfig.loadRecord($attachmentEditorContainer,currRecordRef)
			}			
		}
				
		initAttachmentRecordEditBehavior($attachmentEditorContainer, componentContext,
					recordProxy, attachObjectRef,validateInput)
		loadCurrentRecordIntoPopup()
		
		// Save the function pointer to load the record into the popup. If the comment is updated, this is needed, so to 
		// list of comments can be indirectly updated in the popup. 
		loadRecordIntoPopupAttachmentEditor = loadCurrentRecordIntoPopup

	});
	
	$attachContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoAttachmentEditor,
		validateValue: validateInput
	})
	
	
}


