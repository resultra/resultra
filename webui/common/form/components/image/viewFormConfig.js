


function initImageRecordEditBehavior($container,componentContext,recordProxy, imageObjRef,remoteValidationCallback) {
	
	var imageFieldID = imageObjRef.properties.fieldID
	
	
	function getCurrentImageVal() {
		var currRecordRef = recordProxy.getRecordFunc()
		var imageFieldID = imageObjRef.properties.fieldID
		var fieldVal = currRecordRef.fieldValues[imageFieldID]
		if(fieldVal === undefined || fieldVal === null) {
			return null
		} else {
			return fieldVal
		}
	}
	
	function setImageVal(imageVal) {
		
		// Validation is done in the popup.
		var imageTextValueFormat = {
			context:"image",
			format:"general"
		}
		var currRecordRef = recordProxy.getRecordFunc()
		var setRecordValParams = { 
			parentDatabaseID:currRecordRef.parentDatabaseID,
			recordID:currRecordRef.recordID, 
			changeSetID: recordProxy.changeSetID,
			fieldID:imageFieldID, 
			attachment:imageVal 
		}
		jsonAPIRequest("recordUpdate/setImageFieldValue",setRecordValParams,function(replyData) {
			// After updating the record, the local cache of records will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			recordProxy.updateRecordFunc(replyData)

		}) // set record's text field value
				
			
	}
	
	
	function initImageEditPopup(attachmentID) {
		
		var $imageButton = $container.find(".imageEditLinkButton")
		
		$imageButton.unbind("click")
		$imageButton.find("input").remove()
		
		function initImageButtonForAddingImage() {
			
			var $addAttachmentInput = $('<input class="uploadInput" type="file">')
			$imageButton.append($addAttachmentInput)
			
			function setAttachmentFromDialog(newAttachmentID) {
				setImageVal(newAttachmentID)
			}		
			var addAttachmentParams = {
				parentDatabaseID: componentContext.databaseID,
				setAttachmentCallback: setAttachmentFromDialog,
				$addAttachmentInput: $addAttachmentInput
			}
			initAddAttachmentThenOpenInfoDialogButton(addAttachmentParams)
		}
		
		function initImageButtonForEditingExistingImage(attachmentID) {
					
			$imageButton.click(function(e) {			
				var attachmentParams = {
					attachmentID: attachmentID,
					parentDatabaseID: componentContext.databaseID,
					setAttachmentCallback: setImageVal
				}
				openSingleAttachmentDialog(attachmentParams)							
			})
		}
		
		// Initialize the button to either add a new image or 
		// edit/replace the existing one.
		if(attachmentID === null) {
			initImageButtonForAddingImage()
		} else {
			initImageButtonForEditingExistingImage(attachmentID)
		}
		
	}
	
	var validateImageInput = function(validationCompleteCallback) {	
		remoteValidationCallback(getCurrentImageVal(), function(validationResult) {
			setupFormComponentValidationPrompt($container,validationResult,validationCompleteCallback)	
		}) 
	}

	function loadRecordIntoImage($imageContainer, recordRef) {
	
		console.log("loadRecordIntoImage: loading record into text box: " + JSON.stringify(recordRef))
	
		function setImageDisplay(attachmentID) {
			var $imageDisplay = $imageContainer.find('.imageDisplay')
			initSingleAttachmentImagePopupLink($imageContainer,$imageDisplay,attachmentID)			
		}
	
		// text box is linked to a field value

		console.log("loadRecordIntoImage: Field ID to load data:" + imageFieldID)

		var fieldVal = recordRef.fieldValues[imageFieldID]
		if(fieldVal===undefined || fieldVal===null) {
			setImageDisplay(null)
			initImageEditPopup(null)
		} else {
			setImageDisplay(fieldVal)
			initImageEditPopup(fieldVal)
		}
	
	}



	function initImageFieldEditBehavior(componentContext, $container,
					recordProxy, imageObjRef) {
	
		var imageFieldID = imageObjRef.properties.fieldID
		var $clearValueButton = $container.find(".imageComponentClearValueButton")
	
		var fieldRef = getFieldRef(imageFieldID)
		if(fieldRef.isCalcField) {
			$imageInput.prop('disabled',true);
			return;  // stop initialization, the text box is read only.
		}
		
		initImageClearValueControl($container,imageObjRef)
	
		if(formComponentIsReadOnly(imageObjRef.properties.permissions)) {
			//$imageInput.prop('disabled',true);
		} else {
			//$imageInput.prop('disabled',false);
		}
			
					
		initButtonControlClickHandler($clearValueButton,function() {
				console.log("Clear value clicked for text box")
			setImageVal(null)	
		})
			
	}	// initImageFieldEditBehavior
	
	
	initImageEditPopup(null)
	
	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoImage,
		validateValue: validateImageInput
	})
	
	$container.data("componentContext",componentContext)
	
	
	initImageFieldEditBehavior(componentContext, $container,
			recordProxy, imageObjRef)
	
}


function initImageFormRecordEditBehavior($container,componentContext,recordProxy, imageObjRef) {
		
	var validateImageInput = function(inputVal, validationCompleteCallback) {
		
		var validationParams = {
			parentFormID: imageObjRef.parentFormID,
			imageID: imageObjRef.imageID,
			inputVal: inputVal
		}
		jsonAPIRequest("frm/image/validateInput", validationParams, function(validationResult) {
			validationCompleteCallback(validationResult)
		})	
		
	}
	initImageRecordEditBehavior($container,componentContext,recordProxy, imageObjRef,validateImageInput)
	
}


function initImageTableRecordEditBehavior($container,componentContext,recordProxy, imageObjRef) {
		
	var validateImageInput = function(inputVal, validationCompleteCallback) {
		
		var validationParams = {
			parentTableID: imageObjRef.parentTableID,
			imageID: imageObjRef.imageID,
			inputVal: inputVal
		}
		jsonAPIRequest("tableView/image/validateInput", validationParams, function(validationResult) {
			validationCompleteCallback(validationResult)
		})	
		
	}
	initImageRecordEditBehavior($container,componentContext,recordProxy, imageObjRef,validateImageInput)
	
}