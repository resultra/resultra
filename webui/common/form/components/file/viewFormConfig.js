


function initFileRecordEditBehavior($container,componentContext,recordProxy, fileObjRef,remoteValidationCallback) {
	
	var fileFieldID = fileObjRef.properties.fieldID
	
	
	function getCurrentFileVal() {
		var currRecordRef = recordProxy.getRecordFunc()
		var fileFieldID = fileObjRef.properties.fieldID
		var fieldVal = currRecordRef.fieldValues[fileFieldID]
		if(fieldVal === undefined || fieldVal === null) {
			return null
		} else {
			return fieldVal
		}
	}
	
	function setFileVal(fileVal) {
		
		// Validation is done in the popup.
		var fileTextValueFormat = {
			context:"file",
			format:"general"
		}
		var currRecordRef = recordProxy.getRecordFunc()
		var setRecordValParams = { 
			parentDatabaseID:currRecordRef.parentDatabaseID,
			recordID:currRecordRef.recordID, 
			changeSetID: recordProxy.changeSetID,
			fieldID:fileFieldID, 
			attachment:fileVal,
			valueFormat: fileTextValueFormat 
		}
		jsonAPIRequest("recordUpdate/setFileFieldValue",setRecordValParams,function(replyData) {
			// After updating the record, the local cache of records will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			recordProxy.updateRecordFunc(replyData)

		}) // set record's text field value
				
			
	}
	
	
	function initFileEditPopup(attachmentID) {
		
		var $fileButton = $container.find(".fileEditLinkButton")
		
		$fileButton.unbind("click")
		$fileButton.find("input").remove()
		
		function initFileButtonForAddingFile() {
			
			var $addAttachmentInput = $('<input class="uploadInput" type="file">')
			$fileButton.append($addAttachmentInput)
			
			function setAttachmentFromDialog(newAttachmentID) {
				setFileVal(newAttachmentID)
			}		
			var addAttachmentParams = {
				parentDatabaseID: componentContext.databaseID,
				setAttachmentCallback: setAttachmentFromDialog,
				$addAttachmentInput: $addAttachmentInput
			}
			initAddAttachmentThenOpenInfoDialogButton(addAttachmentParams)
		}
		
		function initFileButtonForEditingExistingFile(attachmentID) {
					
			$fileButton.click(function(e) {			
				var attachmentParams = {
					attachmentID: attachmentID,
					parentDatabaseID: componentContext.databaseID,
					setAttachmentCallback: setFileVal
				}
				openSingleAttachmentDialog(attachmentParams)							
			})
		}
		
		// Initialize the button to either add a new file or 
		// edit/replace the existing one.
		if(attachmentID === null) {
			initFileButtonForAddingFile()
		} else {
			initFileButtonForEditingExistingFile(attachmentID)
		}
		
	}
	
	var validateFileInput = function(validationCompleteCallback) {	
		remoteValidationCallback(getCurrentFileVal(), function(validationResult) {
			setupFormComponentValidationPrompt($container,validationResult,validationCompleteCallback)	
		}) 
	}

	function loadRecordIntoFile($fileContainer, recordRef) {
	
		console.log("loadRecordIntoFile: loading record into text box: " + JSON.stringify(recordRef))
	
		function setFileDisplay(attachmentID) {
			$fileDisplay = $fileContainer.find('.fileDisplay')

			if(attachmentID != null) {
				var getRefParams = { attachmentID: attachmentID }
				jsonAPIRequest("attachment/getReference", getRefParams, function(attachRef) {
					$fileDisplay.text(attachRef.attachmentInfo.title)
					$fileDisplay.attr("href",attachRef.url)
				})
			} else {
				$fileDisplay.text("")
				$fileDisplay.attr("href","")
			}	
			
		}
	
		// text box is linked to a field value

		console.log("loadRecordIntoFile: Field ID to load data:" + fileFieldID)

		var fieldVal = recordRef.fieldValues[fileFieldID]
		if(fieldVal===undefined || fieldVal===null) {
			setFileDisplay(null)
		} else {
			setFileDisplay(fieldVal)
		}
		initFileEditPopup(fieldVal)
	
	}



	function initFileFieldEditBehavior(componentContext, $container,
					recordProxy, fileObjRef) {
	
		var fileFieldID = fileObjRef.properties.fieldID
		var $clearValueButton = $container.find(".fileComponentClearValueButton")
	
		var fieldRef = getFieldRef(fileFieldID)
		if(fieldRef.isCalcField) {
			$fileInput.prop('disabled',true);
			return;  // stop initialization, the text box is read only.
		}
		
		initFileClearValueControl($container,fileObjRef)
	
		if(formComponentIsReadOnly(fileObjRef.properties.permissions)) {
			//$fileInput.prop('disabled',true);
		} else {
			//$fileInput.prop('disabled',false);
		}
			
					
		initButtonControlClickHandler($clearValueButton,function() {
				console.log("Clear value clicked for text box")
			setFileVal(null)	
		})
			
	}	// initFileFieldEditBehavior
	
	
	initFileEditPopup(null)
	
	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoFile,
		validateValue: validateFileInput
	})
	
	$container.data("componentContext",componentContext)
	
	
	initFileFieldEditBehavior(componentContext, $container,
			recordProxy, fileObjRef)
	
}


function initFileFormRecordEditBehavior($container,componentContext,recordProxy, fileObjRef) {
		
	var validateFileInput = function(inputVal, validationCompleteCallback) {
		
		var validationParams = {
			parentFormID: fileObjRef.parentFormID,
			fileID: fileObjRef.fileID,
			inputVal: inputVal
		}
		jsonAPIRequest("frm/file/validateInput", validationParams, function(validationResult) {
			validationCompleteCallback(validationResult)
		})	
		
	}
	initFileRecordEditBehavior($container,componentContext,recordProxy, fileObjRef,validateFileInput)
	
}


function initFileTableRecordEditBehavior($container,componentContext,recordProxy, fileObjRef) {
		
	var validateFileInput = function(inputVal, validationCompleteCallback) {
		
		var validationParams = {
			parentTableID: fileObjRef.parentTableID,
			fileID: fileObjRef.fileID,
			inputVal: inputVal
		}
		jsonAPIRequest("tableView/file/validateInput", validationParams, function(validationResult) {
			validationCompleteCallback(validationResult)
		})	
		
	}
	initFileRecordEditBehavior($container,componentContext,recordProxy, fileObjRef,validateFileInput)
	
}