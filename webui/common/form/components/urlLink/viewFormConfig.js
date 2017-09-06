


function initUrlLinkRecordEditBehavior($container,componentContext,recordProxy, urlLinkObjRef,remoteValidationCallback) {
	
	var validateUrlLinkInput = function(validationCompleteCallback) {
		
		
		if(checkboxComponentIsDisabled($container)) {
			validationCompleteCallback(true)
			return
		}
		
		var $urlLinkInput = $container.find('input')
		var currVal = $urlLinkInput.val()
		
		remoteValidationCallback(currVal, function(validationResult) {
			setupFormComponentValidationPrompt($container,validationResult,validationCompleteCallback)	
		}) 
				
		
	}

	function loadRecordIntoUrlLink($urlLinkContainer, recordRef) {
	
		console.log("loadRecordIntoUrlLink: loading record into text box: " + JSON.stringify(recordRef))
	
		var $urlLinkInput = $urlLinkContainer.find('input')
	
		// text box is linked to a field value
		var urlLinkFieldID = urlLinkObjRef.properties.fieldID

		console.log("loadRecordIntoUrlLink: Field ID to load data:" + urlLinkFieldID)

		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(urlLinkFieldID)) {

			var fieldVal = recordRef.fieldValues[urlLinkFieldID]
		
			if(fieldVal === null) {
				$urlLinkInput.val("")
			} else {
				$urlLinkInput.val(fieldVal)
			
			}

		} // If record has a value for the current container's associated field ID.
		else
		{
			$urlLinkInput.val("") // clear the value in the container
		}	
	
	}



	function initUrlLinkFieldEditBehavior(componentContext, $container,$urlLinkInput,
					recordProxy, urlLinkObjRef) {
	
		var urlLinkFieldID = urlLinkObjRef.properties.fieldID
		var $clearValueButton = $container.find(".urlLinkComponentClearValueButton")
	
		var fieldRef = getFieldRef(urlLinkFieldID)
		if(fieldRef.isCalcField) {
			$urlLinkInput.prop('disabled',true);
			return;  // stop initialization, the text box is read only.
		}
		
		initUrlLinkClearValueControl($container,urlLinkObjRef)
	
		if(formComponentIsReadOnly(urlLinkObjRef.properties.permissions)) {
			$urlLinkInput.prop('disabled',true);
		} else {
			$urlLinkInput.prop('disabled',false);
		}
	
		
		var fieldType = fieldRef.type
		
		function setUrlLinkVal(urlLinkVal) {
			
			validateUrlLinkInput(function(inputIsValid) {
				if (inputIsValid) {
					var urlLinkTextValueFormat = {
						context:"urlLink",
						format:"general"
					}
					var currRecordRef = recordProxy.getRecordFunc()
					var setRecordValParams = { 
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID, 
						changeSetID: recordProxy.changeSetID,
						fieldID:urlLinkFieldID, 
						value:urlLinkVal,
						valueFormat: urlLinkTextValueFormat 
					}
					jsonAPIRequest("recordUpdate/setUrlLinkFieldValue",setRecordValParams,function(replyData) {
						// After updating the record, the local cache of records will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						recordProxy.updateRecordFunc(replyData)
			
					}) // set record's text field value
					
				} // inputIsValid
				
			})
				
		}
		
		function setUrlLinkValueListValue(textVal) {
			var $urlLinkInput = $container.find('input')
			$urlLinkInput.val(textVal)
			setUrlLinkVal(textVal)
		}
		
		initButtonControlClickHandler($clearValueButton,function() {
				console.log("Clear value clicked for text box")
			setUrlLinkVal(null)	
		})
		

		$urlLinkInput.focusout(function () {
			// Retrieve the "raw input" value entered by the user and 
			// update the "rawVal" data setting on the text box.
			var inputVal = $urlLinkInput.val()
			console.log("Text Box focus out:" + inputVal)
			setUrlLinkVal(inputVal)			
	
		}) // focus out
	
	}	
	
	var $urlLinkInput = $container.find("input")

	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoUrlLink,
		validateValue: validateUrlLinkInput
	})
	
	$container.data("componentContext",componentContext)
	
	
	// When the user clicks on the text box input control, prevent the click from propagating higher.
	// This allows the user to change the values without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the input control.
	$urlLinkInput.click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
	initUrlLinkFieldEditBehavior(componentContext, $container,$urlLinkInput,
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