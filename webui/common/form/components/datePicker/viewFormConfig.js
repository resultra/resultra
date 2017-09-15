



function initDatePickerRecordEditBehavior($datePickerContainer, componentContext,recordProxy, datePickerObjectRef,remoteValidationFunc) {

	var validateDatePickerInput = function(validationCompleteCallback) {
		
		if(datePickerComponentIsDisabled($datePickerContainer)) {
			validationCompleteCallback(true)
			return
		}
		
		var currVal = getDatePickerFormComponentUTCDate($datePickerContainer)
		remoteValidationFunc(currVal,function(validationResult) {
			setupFormComponentValidationPrompt($datePickerContainer,validationResult,validationCompleteCallback)			
		})	
		
	}


	function initDatePickerFieldEditBehavior(componentContext,recordProxy, 
					datePickerObjectRef,$datePickerContainer) {
	

		var $datePickerControl = datePickerInputFromContainer($datePickerContainer)
		var $clearValueButton = $datePickerContainer.find(".datePickerComponentClearValueButton")
		var $calendarIcon = $datePickerContainer.find(".datePickerCalendarButton")
		var $datePickerInput = $datePickerContainer.find("input")

	
		var fieldID = datePickerObjectRef.properties.fieldID
	
		var fieldRef = getFieldRef(fieldID)
		if(fieldRef.isCalcField) {
			$datePickerInput.prop("disabled",true)
			$clearValueButton.css("display","none")
			$calendarIcon.css("display","none")
			return;  // stop initialization, the check box is read only.
		}
		
		initDatePickerAddonControls($datePickerContainer,datePickerObjectRef)
		
		$datePickerContainer.find(".datePickerInputContainer").click(function(e) {
		
			// This is important - if a click hits an object, then stop the propagation of the click
			// to the parent div(s), including the canvas itself. If the parent canvas
			// gets a click, it will deselect all the items (see initObjectCanvasSelectionBehavior)
			e.stopPropagation();
		})
		
	
		function setDateValue(dateVal) {
			validateDatePickerInput(function(inputIsValid) {
				if(inputIsValid) {
					var currRecordRef = recordProxy.getRecordFunc()
					var setRecordValParams = {
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID, 
						changeSetID: recordProxy.changeSetID,
						fieldID:fieldID, 
						value:dateVal}
					console.log("Setting date value: " + JSON.stringify(setRecordValParams))
	
					jsonAPIRequest("recordUpdate/setTimeFieldValue",setRecordValParams,function(updatedRecordRef) {
		
						// After updating the record, the local cache of records in currentRecordSet will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						recordProxy.updateRecordFunc(updatedRecordRef)
					}) // set record's text field value		
					
				} // inputIsValid
			})
		}
	
		initButtonControlClickHandler($clearValueButton,function() {
				console.log("Clear value clicked for date picker")
				setDateValue(null)
		})
		
	
		
		$datePickerControl.on('dp.change',function (e) {
		    console.log("date picker changed dates")		
			// The date passed with the event will be false if the date has been cleared.
			if(e.date !== false) {
				var objectRef = getContainerObjectRef($datePickerContainer)
				var dateParam = e.date.toISOString()
				setDateValue(dateParam)		
			} else {
				setDateValue(null) // clear the date
			}
		
		});	
	
	}


	function loadRecordIntoDatePicker($datePicker, recordRef) {
	
		console.log("loadRecordIntoDatePicker: loading record into date picker: " + JSON.stringify(recordRef))
	
		var datePickerObjectRef = $datePicker.data("objectRef")
		var datePickerFieldID = datePickerObjectRef.properties.fieldID
		
		console.log("loadRecordIntoDatePicker: Field ID to load data:" + datePickerFieldID)

		var $datePickerInput = datePickerInputFromContainer($datePicker)
		var $datePickerInputField = $datePicker.find(".datePickerInputField")
		
		function setConditionalFormat() {
			var rawFieldVal = null
			if(recordRef.fieldValues.hasOwnProperty(datePickerFieldID)) {
				var unformattedVal = recordRef.fieldValues[datePickerFieldID]
				rawFieldVal = moment(unformattedVal).toDate()
			}
			setBackgroundConditionalDateFormat($datePickerInputField,
					datePickerObjectRef.properties.conditionalFormats,rawFieldVal)
		}
		setConditionalFormat()
		

		if(recordRef.fieldValues.hasOwnProperty(datePickerFieldID)) {

			// If record has a value for the current container's associated field ID.
			var fieldVal = recordRef.fieldValues[datePickerFieldID]
		
			if (fieldVal === null) {
				$datePickerInput.data("DateTimePicker").clear()
			} else {
				// Values need to be set using the same format the date picker control was initialized with.
				var dateVal = moment(fieldVal)
				setDatePickerFormComponentDate($datePicker,datePickerObjectRef,dateVal)			
			}
		} else {
			// There's no value in the current record for this field, so clear the value in the container
			$datePickerInput.data("DateTimePicker").clear()
		}
	
	}

			
	$datePickerContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoDatePicker,
		validateValue: validateDatePickerInput
	})
	initDatePickerFieldEditBehavior(componentContext,recordProxy, datePickerObjectRef,$datePickerContainer)
	
}

function initFormDatePickerEditBehavior($datePickerContainer, componentContext,recordProxy, datePickerObjectRef) {
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentFormID: datePickerObjectRef.parentFormID,
			datePickerID: datePickerObjectRef.datePickerID,
			inputVal: inputVal
		}
		jsonAPIRequest("frm/datePicker/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	initDatePickerRecordEditBehavior($datePickerContainer,componentContext,recordProxy, datePickerObjectRef,validateInput)
}


function initTableViewDatePickerEditBehavior($datePickerContainer, componentContext,recordProxy, datePickerObjectRef) {
	
	function validateInput(inputVal,validationResultCallback) {
		var validationParams = {
			parentTableID: datePickerObjectRef.parentTableID,
			datePickerID: datePickerObjectRef.datePickerID,
			inputVal: inputVal
		}
		jsonAPIRequest("tableView/datePicker/validateInput", validationParams, function(validationResult) {
			validationResultCallback(validationResult)
		})
	}
	
	initDatePickerContainerControls($datePickerContainer,datePickerObjectRef)
	
	initDatePickerRecordEditBehavior($datePickerContainer,componentContext,recordProxy, datePickerObjectRef,validateInput)
}