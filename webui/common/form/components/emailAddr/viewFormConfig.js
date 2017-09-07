


function initEmailAddrRecordEditBehavior($container,componentContext,recordProxy, emailAddrObjRef,remoteValidationCallback) {
	
	
	function getCurrentEmailAddrVal() {
		var currRecordRef = recordProxy.getRecordFunc()
		var emailAddrFieldID = emailAddrObjRef.properties.fieldID
		var fieldVal = currRecordRef.fieldValues[emailAddrFieldID]
		if(fieldVal === undefined || fieldVal === null) {
			return null
		} else {
			return fieldVal
		}
	}
	
	var validateEmailAddrInput = function(validationCompleteCallback) {
			
		
		remoteValidationCallback(getCurrentEmailAddrVal(), function(validationResult) {
			setupFormComponentValidationPrompt($container,validationResult,validationCompleteCallback)	
		}) 
				
		
	}

	function loadRecordIntoEmailAddr($emailAddrContainer, recordRef) {
	
		console.log("loadRecordIntoEmailAddr: loading record into text box: " + JSON.stringify(recordRef))
	
		function setEmailAddressDisplay(emailAddr) {
			$emailAddrDisplay = $emailAddrContainer.find('.emailAddrDisplay')
			$emailAddrDisplay.text(emailAddr)
			$emailAddrDisplay.attr("href","mailto:" + emailAddr)
		}
	
		// text box is linked to a field value
		var emailAddrFieldID = emailAddrObjRef.properties.fieldID

		console.log("loadRecordIntoEmailAddr: Field ID to load data:" + emailAddrFieldID)

		var fieldVal = recordRef.fieldValues[emailAddrFieldID]
		if(fieldVal===undefined || fieldVal===null) {
			setEmailAddressDisplay(null)
		} else {
			setEmailAddressDisplay(fieldVal)
		}
	
	}



	function initEmailAddrFieldEditBehavior(componentContext, $container,
					recordProxy, emailAddrObjRef) {
	
		var emailAddrFieldID = emailAddrObjRef.properties.fieldID
		var $clearValueButton = $container.find(".emailAddrComponentClearValueButton")
	
		var fieldRef = getFieldRef(emailAddrFieldID)
		if(fieldRef.isCalcField) {
			$emailAddrInput.prop('disabled',true);
			return;  // stop initialization, the text box is read only.
		}
		
		initEmailAddrClearValueControl($container,emailAddrObjRef)
	
		if(formComponentIsReadOnly(emailAddrObjRef.properties.permissions)) {
			//$emailAddrInput.prop('disabled',true);
		} else {
			//$emailAddrInput.prop('disabled',false);
		}
			
		function setEmailAddrVal(emailAddrVal) {
			
			// Validation is done in the popup.
			var emailAddrTextValueFormat = {
				context:"emailAddr",
				format:"general"
			}
			var currRecordRef = recordProxy.getRecordFunc()
			var setRecordValParams = { 
				parentDatabaseID:currRecordRef.parentDatabaseID,
				recordID:currRecordRef.recordID, 
				changeSetID: recordProxy.changeSetID,
				fieldID:emailAddrFieldID, 
				value:emailAddrVal,
				valueFormat: emailAddrTextValueFormat 
			}
			jsonAPIRequest("recordUpdate/setEmailAddrFieldValue",setRecordValParams,function(replyData) {
				// After updating the record, the local cache of records will
				// be out of date. So after updating the record on the server, the locally cached
				// version of the record also needs to be updated.
				recordProxy.updateRecordFunc(replyData)
	
			}) // set record's text field value
					
				
		}
		
		
		function initEditAddrEditPopup() {
			var $emailAddrButton = $container.find(".emailAddrEditLinkButton")
			
			$emailAddrButton.popover({
				html: 'true',
				content: function() { return emailAddrEditPopupViewContainerHTML() },
				trigger: 'manual',
				placement: 'auto',
				container:'body'
			})

			$emailAddrButton.click(function(e) {
				$(this).popover('toggle')
				e.stopPropagation()
			})
			$emailAddrButton.on('shown.bs.popover', function()
			{
			    //get the actual shown popover
			    var $popover = $(this).data('bs.popover').tip();
				
				// If there's a click outside the popover, hide the popover.
				// This solution is described here:
				//    https://stackoverflow.com/questions/152975/how-do-i-detect-a-click-outside-an-element
				$popover.click(function(e) {
					e.stopPropagation()	
				})
				$('html').click(function() {
					$emailAddrButton.popover('hide')	
				})
		
				var $closePopupButton = $popover.find(".closeEmailAddrEditorPopup")
				initButtonControlClickHandler($closePopupButton,function() {
					$emailAddrButton.popover('hide')
				})
			
				var $emailAddrInput = $popover.find(".emailAddrComponentInput")
				$emailAddrInput.val(getCurrentEmailAddrVal())
				
				$emailAddrInput.focusout(function () {
					// Retrieve the "raw input" value entered by the user and 
					// update the "rawVal" data setting on the text box.
					var inputVal = $emailAddrInput.val()
					remoteValidationCallback(inputVal, function(validationResult) {
						if(validationResult.validationSucceeded) {
							setEmailAddrVal(inputVal)				
						}
					}) 
					
	
				}) // focus out
				
			});
			
		}
		initEditAddrEditPopup()
			
		initButtonControlClickHandler($clearValueButton,function() {
				console.log("Clear value clicked for text box")
			setEmailAddrVal(null)	
		})
			
	}	// initEmailAddrFieldEditBehavior
	
	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoEmailAddr,
		validateValue: validateEmailAddrInput
	})
	
	$container.data("componentContext",componentContext)
	
	
	initEmailAddrFieldEditBehavior(componentContext, $container,
			recordProxy, emailAddrObjRef)
	
}


function initEmailAddrFormRecordEditBehavior($container,componentContext,recordProxy, emailAddrObjRef) {
		
	var validateEmailAddrInput = function(inputVal, validationCompleteCallback) {
		
		var validationParams = {
			parentFormID: emailAddrObjRef.parentFormID,
			emailAddrID: emailAddrObjRef.emailAddrID,
			inputVal: inputVal
		}
		jsonAPIRequest("frm/emailAddr/validateInput", validationParams, function(validationResult) {
			validationCompleteCallback(validationResult)
		})	
		
	}
	initEmailAddrRecordEditBehavior($container,componentContext,recordProxy, emailAddrObjRef,validateEmailAddrInput)
	
}


function initEmailAddrTableRecordEditBehavior($container,componentContext,recordProxy, emailAddrObjRef) {
		
	var validateEmailAddrInput = function(inputVal, validationCompleteCallback) {
		
		var validationParams = {
			parentTableID: emailAddrObjRef.parentTableID,
			emailAddrID: emailAddrObjRef.emailAddrID,
			inputVal: inputVal
		}
		jsonAPIRequest("tableView/emailAddr/validateInput", validationParams, function(validationResult) {
			validationCompleteCallback(validationResult)
		})	
		
	}
	initEmailAddrRecordEditBehavior($container,componentContext,recordProxy, emailAddrObjRef,validateEmailAddrInput)
	
}