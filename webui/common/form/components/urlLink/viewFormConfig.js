// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.



function initUrlLinkRecordEditBehavior($container,componentContext,recordProxy, urlLinkObjRef,remoteValidationCallback) {
	
	
	function getCurrentUrlLinkVal() {
		var urlLinkFieldID = urlLinkObjRef.properties.fieldID
		var currRecordRef = recordProxy.getRecordFunc()
		var fieldVal = currRecordRef.fieldValues[urlLinkFieldID]
		if(fieldVal === undefined || fieldVal === null) {
			return null
		} else {
			return fieldVal
		}
	}
	
	var validateUrlLinkInput = function(validationCompleteCallback) {
		var currVal = getCurrentUrlLinkVal()	
		remoteValidationCallback(currVal, function(validationResult) {
			setupFormComponentValidationPrompt($container,validationResult,validationCompleteCallback)	
		}) 
	}

	function loadRecordIntoUrlLink($urlLinkContainer, recordRef) {
	
		console.log("loadRecordIntoUrlLink: loading record into text box: " + JSON.stringify(recordRef))
	
		function setUrlDisplay(url) {
			var $urlLinkDisplay = $urlLinkContainer.find('.urlLinkDisplay')
			$urlLinkDisplay.text(url)
			$urlLinkDisplay.attr("href",url)
		}
	
		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		var urlLinkFieldID = urlLinkObjRef.properties.fieldID
		var fieldVal = recordRef.fieldValues[urlLinkFieldID]
		if(fieldVal === undefined || fieldVal === null) {
			setUrlDisplay(null)
		} // If record has a value for the current container's associated field ID.
		else
		{
			setUrlDisplay(fieldVal)
		}	
	
	}
	
	function setUrlLinkVal(urlLinkVal) {
		
		var currRecordRef = recordProxy.getRecordFunc()
		var urlLinkFieldID = urlLinkObjRef.properties.fieldID
		var setRecordValParams = { 
			parentDatabaseID:currRecordRef.parentDatabaseID,
			recordID:currRecordRef.recordID, 
			changeSetID: recordProxy.changeSetID,
			fieldID:urlLinkFieldID, 
			value:urlLinkVal 
		}
		jsonAPIRequest("recordUpdate/setUrlLinkFieldValue",setRecordValParams,function(replyData) {
			// After updating the record, the local cache of records will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			recordProxy.updateRecordFunc(replyData)

		}) // set record's text field value
				
			
	}
	

	function initUrlLinkFieldEditBehavior(componentContext, $container,
					recordProxy, urlLinkObjRef) {
	
		var urlLinkFieldID = urlLinkObjRef.properties.fieldID
		var $clearValueButton = $container.find(".urlLinkComponentClearValueButton")
	
		var fieldRef = getFieldRef(urlLinkFieldID)
						
		if(fieldRef.isCalcField) {
			$urlLinkInput.prop('disabled',true);
			return;  // stop initialization, the text box is read only.
		}
		
		initUrlLinkClearValueControl($container,urlLinkObjRef)
		
		
		function initEditLinkPopup() {
			var $urlLink = $container.find(".urlLinkEditLinkButton")
			$urlLink.popover({
				html: 'true',
				content: function() { return urlLinkEditPopupViewContainerHTML() },
				trigger: 'manual',
				placement: 'auto',
				container:'body'
			})

			$urlLink.click(function(e) {
				$(this).popover('toggle')
				e.stopPropagation()
			})
			$urlLink.on('shown.bs.popover', function()
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
					$urlLink.popover('hide')	
				})
		
				var $closePopupButton = $popover.find(".closeLinkEditorPopup")
				initButtonControlClickHandler($closePopupButton,function() {
					$urlLink.popover('hide')
				})
			
				var $urlLinkInput = $popover.find(".urlLinkComponentInput")
				$urlLinkInput.val(getCurrentUrlLinkVal())
				
				
				$urlLinkInput.focusout(function () {
					// Retrieve the "raw input" value entered by the user and 
					// update the "rawVal" data setting on the text box.
					var inputVal = $urlLinkInput.val()
					remoteValidationCallback(inputVal, function(validationResult) {
						if(validationResult.validationSucceeded) {
							setUrlLinkVal(inputVal)
							clearFormComponentValidationPrompt($container)
						}
					}) 
	
				}) // focus out
				
			});
			
		}
		initEditLinkPopup()
		
		var $editLinkButton = $container.find(".urlLinkEditLinkButton")
		var $editLinkDisplay = $container.find(".formInputStaticInputContainer")
		if(formComponentIsReadOnly(urlLinkObjRef.properties.permissions)) {
			$editLinkButton.hide()
			$editLinkDisplay.css("background-color","rgb(238,238,238)")
		} else {
			$editLinkButton.show()
			$editLinkDisplay.css("background-color","")
		}
	
						
		initButtonControlClickHandler($clearValueButton,function() {
				console.log("Clear value clicked for text box")
			setUrlLinkVal(null)	
		})
			
	}	
	

	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoUrlLink,
		validateValue: validateUrlLinkInput
	})
	
	$container.data("componentContext",componentContext)
	
	
	initUrlLinkFieldEditBehavior(componentContext, $container,
			recordProxy, urlLinkObjRef)
	
}


function initUrlLinkFormRecordEditBehavior($container,componentContext,recordProxy, urlLinkObjRef) {
		
	var validateUrlLinkInput = function(inputVal, validationCompleteCallback) {
		
		var validationParams = {
			parentFormID: urlLinkObjRef.parentFormID,
			urlLinkID: urlLinkObjRef.urlLinkID,
			inputVal: inputVal
		}
		jsonAPIRequest("frm/urlLink/validateInput", validationParams, function(validationResult) {
			validationCompleteCallback(validationResult)
		})	
		
	}
	initUrlLinkRecordEditBehavior($container,componentContext,recordProxy, urlLinkObjRef,validateUrlLinkInput)
	
}


function initUrlLinkTableRecordEditBehavior($container,componentContext,recordProxy, urlLinkObjRef) {
		
	var validateUrlLinkInput = function(inputVal, validationCompleteCallback) {
		
		var validationParams = {
			parentTableID: urlLinkObjRef.parentTableID,
			urlLinkID: urlLinkObjRef.urlLinkID,
			inputVal: inputVal
		}
		jsonAPIRequest("tableView/urlLink/validateInput", validationParams, function(validationResult) {
			validationCompleteCallback(validationResult)
		})	
		
	}
	initUrlLinkRecordEditBehavior($container,componentContext,recordProxy, urlLinkObjRef,validateUrlLinkInput)
	
}