



function initNumberInputRecordEditBehavior($container,componentContext,recordProxy, numberInputObjectRef) {
	
	var $numberInputInput = $container.find('input')
	var numberInputFieldID = numberInputObjectRef.properties.fieldID
	
	// When the user focuses in/out of the control, the value displayed needs to be toggled
	// between the "raw input value" and the formatted value based upon the number input's properties.
	function getRawInputNumberVal() {	
		var rawVal = $numberInputInput.data("rawVal")
		return convertStringToNumber(rawVal)
	}
	function setRawInputVal(rawVal) { 
		var rawValStr = rawVal + '' // convert to a string for storing
		$numberInputInput.data("rawVal",rawValStr) 
	}
	function getRawInputVal() { return $numberInputInput.data("rawVal") }
	
	function getCurrentFieldValue() {
		var currRecordRef = recordProxy.getRecordFunc()
		if(currRecordRef.fieldValues.hasOwnProperty(numberInputFieldID)) {
			var rawFieldVal = currRecordRef.fieldValues[numberInputFieldID]
			return rawFieldVal
		} else {
			return null
		}
	}
		
	
	var validateNumberInput = function(validationCompleteCallback) {
		
		if(numberInputComponentDisabled($container)) {
			validationCompleteCallback(true)
			return
		}	
		var currVal = getRawInputNumberVal()
		var validationParams = {
			parentFormID: numberInputObjectRef.parentFormID,
			numberInputID: numberInputObjectRef.numberInputID,
			inputVal: currVal
		}
		jsonAPIRequest("frm/numberInput/validateInput", validationParams, function(validationResult) {
			if (validationResult.validationSucceeded) {
				$container.popover('destroy')
				validationCompleteCallback(true)
			} else {
				$container.popover({
					html: 'true',
					content: function() { return escapeHTML(validationResult.errorMsg) },
					trigger: 'manual',
					placement: 'auto left'
				})
				$container.popover('show')
				validationCompleteCallback(false)
			}
		})	
		
	}
	
	function loadRecordIntoNumberInput($numberInputContainer, recordRef) {

		console.log("loadRecordIntoNumberInput: loading record into text box: " + JSON.stringify(recordRef))

		var numberInputObjectRef = $numberInputContainer.data("objectRef")
		var $numberInputInput = $numberInputContainer.find('input')
		var componentContext = $numberInputContainer.data("componentContext")

		// text box is linked to a field value
		var numberInputFieldID = numberInputObjectRef.properties.fieldID

		console.log("loadRecordIntoNumberInput: Field ID to load data:" + numberInputFieldID)

		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(numberInputFieldID)) {

			var rawFieldVal = recordRef.fieldValues[numberInputFieldID]
	
			if(rawFieldVal === null) {
				$numberInputInput.val("")
				setRawInputVal("")
			} else {
				console.log("loadRecordIntoNumberInput: Load value into container: " + $(this).attr("id") + " field ID:" + 
							numberInputFieldID + "  value:" + rawFieldVal)
	
				setRawInputVal(rawFieldVal)
	
				var formattedVal = formatNumberValue(numberInputObjectRef.properties.valueFormat.format,rawFieldVal)
				$numberInputInput.val(formattedVal)
			}

		} // If record has a value for the current container's associated field ID.
		else
		{
			$numberInputInput.val("") // clear the value in the container
			setRawInputVal("")
		}	

	}
	
	
	function initNumberInputFieldEditBehavior(componentContext, $container,$numberInputInput,recordProxy, numberInputObjectRef) {
	
		
		var $clearValueButton = $container.find(".numberInputComponentClearValueButton")
	
		var fieldRef = getFieldRef(numberInputFieldID)
		if(fieldRef.isCalcField) {
			$numberInputInput.prop('disabled',true);
			$clearValueButton.hide()
			return;  // stop initialization, the text box is read only.
		}
	
		var fieldType = fieldRef.type
	
		if(formComponentIsReadOnly(numberInputObjectRef.properties.permissions)) {
			$numberInputInput.prop('disabled',true);
			$clearValueButton.hide()
		} else {
			$numberInputInput.prop('disabled',false);
			$clearValueButton.show()
		
		}
	
		function setNumberVal(numberVal) {
			validateNumberInput(function(inputIsValid) {
				if(inputIsValid) {
					
					var currSavedVal = getCurrentFieldValue()
					if (currSavedVal != numberVal) {
						var currRecordRef = recordProxy.getRecordFunc()
						var numberInputNumberValueFormat = {
							context:"numberInput",
							format:"general"
						}
						var setRecordValParams = { 
							parentDatabaseID:currRecordRef.parentDatabaseID,
							recordID:currRecordRef.recordID,
							changeSetID: recordProxy.changeSetID,
							fieldID:numberInputFieldID, 
							value:numberVal,
							valueFormat:numberInputNumberValueFormat
						}
						jsonAPIRequest("recordUpdate/setNumberFieldValue",setRecordValParams,function(replyData) {
							// After updating the record, the local cache of records will
							// be out of date. So after updating the record on the server, the locally cached
							// version of the record also needs to be updated.
							recordProxy.updateRecordFunc(replyData)
		
						}) // set record's number field value
						
					}
					
					
				} // if valid input
			}) // validateNumberInput
			
		
		}	
	
		initButtonControlClickHandler($clearValueButton,function() {
				console.log("Clear value clicked for text box")
		
			var currRecordRef = recordProxy.getRecordFunc()
			setRawInputVal("")
			$numberInputInput.val("")
			setNumberVal(null)
		})
	
		if(numberInputObjectRef.properties.showValueSpinner) {
			var $addButton = $container.find(".addButton")
			initButtonControlClickHandler($addButton,function() {
					console.log("Clear value clicked for text box")
				var numberVal = getRawInputNumberVal()
				if(numberVal != null) {
					numberVal = numberVal + numberInputObjectRef.properties.valueSpinnerStepSize
					setNumberVal(numberVal)						
				}		
			})
			var $subButton = $container.find(".subButton")
			initButtonControlClickHandler($subButton,function() {
					console.log("Clear value clicked for text box")
				var numberVal = getRawInputNumberVal()
				if(numberVal != null) {
					numberVal = numberVal - numberInputObjectRef.properties.valueSpinnerStepSize
					setNumberVal(numberVal)						
				}	
			})
		}

		$numberInputInput.focusin(function() {
			// When focusing on the text input box, replaced the formatted value with 
			// the raw input value.
			var rawInputVal = getRawInputVal()
			console.log("Focus in for number field: raw value for editing: " + rawInputVal)
			$numberInputInput.val(rawInputVal)
		})
	

		$numberInputInput.focusout(function () {

			// Retrieve the "raw input" value entered by the user and 
			// update the "rawVal" data setting on the text box.
			var inputVal = $numberInputInput.val()
			setRawInputVal(inputVal)
		
			// Now that entry of the raw value is complete, revert the 
			// displayed value back to the format specified for the text box.
			var formattedVal = formatNumberValue(numberInputObjectRef.properties.valueFormat.format,inputVal)
			$numberInputInput.val(formattedVal)
		
			var currRecordRef = recordProxy.getRecordFunc()
			
			if(currRecordRef != null) {
		
				// Only update the value if it has changed. Sometimes a user may focus on or tab
				// through a field but not change it. In this case we don't need to update the record.	
				var inputNumberVal =  getRawInputNumberVal()
				setNumberVal(inputNumberVal)
			}
	
		}) // focus out
	
	}

	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoNumberInput,
		validateValue: validateNumberInput
	})
	
	$container.data("componentContext",componentContext)
	
	
	// When the user clicks on the text box input control, prevent the click from propagating higher.
	// This allows the user to change the values without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the input control.
	$numberInputInput.click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
	initNumberInputFieldEditBehavior(componentContext, $container,$numberInputInput,
			recordProxy, numberInputObjectRef)
	
}