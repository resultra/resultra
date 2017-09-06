


function initEmailAddrRecordEditBehavior($container,componentContext,recordProxy, emailAddrObjRef) {
	
	var validateEmailAddrInput = function(validationCompleteCallback) {
		
		if(checkboxComponentIsDisabled($container)) {
			validationCompleteCallback(true)
			return
		}
		
		var $emailAddrInput = $container.find('input')
		
		var currVal = $emailAddrInput.val()
		var validationParams = {
			parentFormID: emailAddrObjRef.parentFormID,
			emailAddrID: emailAddrObjRef.emailAddrID,
			inputVal: currVal
		}
		jsonAPIRequest("frm/emailAddr/validateInput", validationParams, function(validationResult) {
			setupFormComponentValidationPrompt($container,validationResult,validationCompleteCallback)			
		})	
		
	}

	function loadRecordIntoEmailAddr($emailAddrContainer, recordRef) {
	
		console.log("loadRecordIntoEmailAddr: loading record into text box: " + JSON.stringify(recordRef))
	
		var $emailAddrInput = $emailAddrContainer.find('input')
	
		// text box is linked to a field value
		var emailAddrFieldID = emailAddrObjRef.properties.fieldID

		console.log("loadRecordIntoEmailAddr: Field ID to load data:" + emailAddrFieldID)

		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(emailAddrFieldID)) {

			var fieldVal = recordRef.fieldValues[emailAddrFieldID]
		
			if(fieldVal === null) {
				$emailAddrInput.val("")
			} else {
				$emailAddrInput.val(fieldVal)
			
			}

		} // If record has a value for the current container's associated field ID.
		else
		{
			$emailAddrInput.val("") // clear the value in the container
		}	
	
	}



	function initEmailAddrFieldEditBehavior(componentContext, $container,$emailAddrInput,
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
			$emailAddrInput.prop('disabled',true);
		} else {
			$emailAddrInput.prop('disabled',false);
		}
	
		
		var fieldType = fieldRef.type
		
		function setEmailAddrVal(emailAddrVal) {
			
			
			validateEmailAddrInput(function(inputIsValid) {
				if (inputIsValid) {
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
					
				} // inputIsValid
				
			})
				
		}
		
		function setEmailAddrValueListValue(textVal) {
			var $emailAddrInput = $container.find('input')
			$emailAddrInput.val(textVal)
			setEmailAddrVal(textVal)
		}
		
		initButtonControlClickHandler($clearValueButton,function() {
				console.log("Clear value clicked for text box")
		
			var currRecordRef = recordProxy.getRecordFunc()
			setEmailAddrVal(null)
		
		})
		

		$emailAddrInput.focusout(function () {
			// Retrieve the "raw input" value entered by the user and 
			// update the "rawVal" data setting on the text box.
			var inputVal = $emailAddrInput.val()
			console.log("Text Box focus out:" + inputVal)
		
			var currRecordRef = recordProxy.getRecordFunc()
					
			if(currRecordRef != null) {
		
				// Only update the value if it has changed. Sometimes a user may focus on or tab
				// through a field but not change it. In this case we don't need to update the record.
				if(currRecordRef.fieldValues[emailAddrFieldID] !== inputVal) {
						setEmailAddrVal(inputVal)			
				} // if input value is different than currently cached value
			}
	
		}) // focus out
	
	}	
	
	var $emailAddrInput = $container.find("input")

	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoEmailAddr,
		validateValue: validateEmailAddrInput
	})
	
	$container.data("componentContext",componentContext)
	
	
	// When the user clicks on the text box input control, prevent the click from propagating higher.
	// This allows the user to change the values without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the input control.
	$emailAddrInput.click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
	initEmailAddrFieldEditBehavior(componentContext, $container,$emailAddrInput,
			recordProxy, emailAddrObjRef)
	
}