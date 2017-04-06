

function formatNumberInputVal(fieldID, componentContext, rawInputVal, format) {

	var fieldRef = getFieldRef(fieldID)
	return formatNumberValue(format,rawInputVal)
}



function loadRecordIntoNumberInput($numberInputContainer, recordRef) {
	
	console.log("loadRecordIntoNumberInput: loading record into text box: " + JSON.stringify(recordRef))
	
	var numberInputObjectRef = $numberInputContainer.data("objectRef")
	var $numberInputInput = $numberInputContainer.find('input')
	var componentContext = $numberInputContainer.data("componentContext")
	
	
	function setRawInputVal(rawVal) { $numberInputInput.data("rawVal",rawVal) }

	// text box is linked to a field value
	var numberInputFieldID = numberInputObjectRef.properties.fieldID

	console.log("loadRecordIntoNumberInput: Field ID to load data:" + numberInputFieldID)

	// In other words, we are populating the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(numberInputFieldID)) {

		var rawFieldVal = recordRef.fieldValues[numberInputFieldID]
		
		if(rawFieldVal === null) {
			$numberInputInput.val("")
		} else {
			console.log("loadRecordIntoNumberInput: Load value into container: " + $(this).attr("id") + " field ID:" + 
						numberInputFieldID + "  value:" + rawFieldVal)
		
			setRawInputVal(rawFieldVal)
		
			var formattedVal = formatNumberInputVal(numberInputFieldID,componentContext,
					rawFieldVal,numberInputObjectRef.properties.valueFormat.format)

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
	
	var numberInputFieldID = numberInputObjectRef.properties.fieldID
	var $spinnerControls = $container.find(".numberInputSpinnerControls")
	var $clearValueButton = $container.find(".numberInputComponentClearValueButton")
	
	var fieldRef = getFieldRef(numberInputFieldID)
	if(fieldRef.isCalcField) {
		$numberInputInput.prop('disabled',true);
		$spinnerControls.hide()
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
	
	initButtonControlClickHandler($clearValueButton,function() {
			console.log("Clear value clicked for text box")
		
		var currRecordRef = recordProxy.getRecordFunc()
		setNumberVal(null)
		$numberInputInput.data("rawVal","")
		
	})
	
	if(numberInputObjectRef.properties.showValueSpinner) {
		$spinnerControls.show()
		var $addButton = $container.find(".addButton")
		initButtonControlClickHandler($addButton,function() {
				console.log("Clear value clicked for text box")
		
			var inputVal = $numberInputInput.data("rawVal")
			var numberVal = Number(inputVal)
			if(!isNaN(numberVal)) {
				numberVal = numberVal + numberInputObjectRef.properties.valueSpinnerStepSize
				setNumberVal(numberVal)						
			}		
		})
		var $subButton = $container.find(".subButton")
		initButtonControlClickHandler($subButton,function() {
				console.log("Clear value clicked for text box")
		
			var inputVal = $numberInputInput.data("rawVal")
			var numberVal = Number(inputVal)
			if(!isNaN(numberVal)) {
				numberVal = numberVal - numberInputObjectRef.properties.valueSpinnerStepSize
				setNumberVal(numberVal)						
			}		
		})
	} else {
		$spinnerControls.hide()
	}
	
	
		
	if(fieldType == fieldTypeNumber) {
		$numberInputInput.focusin(function() {
			// When focusing on the text input box, replaced the formatted value with 
			// the raw input value.
			var rawInputVal = $numberInputInput.data("rawVal")
			console.log("Focus in for number field: raw value for editing: " + rawInputVal)
			$numberInputInput.val(rawInputVal)
		})
	}
	

	$numberInputInput.focusout(function () {

		var currTextObjRef = getContainerObjectRef($container)		
		var fieldID = numberInputFieldID
		var fieldRef = getFieldRef(numberInputFieldID)
		var fieldType = fieldRef.type
		console.log("Text Box focus out:" 
			+ " ,fieldID: " + numberInputFieldID
		    + " ,fieldType: " + fieldType
			+ " , inputval:" + inputVal)

		// Retrieve the "raw input" value entered by the user and 
		// update the "rawVal" data setting on the text box.
		var inputVal = $numberInputInput.val()
		$numberInputInput.data("rawVal",inputVal)
		
		// Now that entry of the raw value is complete, revert the 
		// displayed value back to the format specified for the text box.
		var formattedVal = formatNumberInputVal(numberInputFieldID,componentContext,
						inputVal,currTextObjRef.properties.valueFormat.format)
		$numberInputInput.val(formattedVal)
		
		var currRecordRef = recordProxy.getRecordFunc()
			
			
		if(currRecordRef != null) {
		
			// Only update the value if it has changed. Sometimes a user may focus on or tab
			// through a field but not change it. In this case we don't need to update the record.
			if(currRecordRef.fieldValues[numberInputFieldID] != inputVal) {
				var numberVal = Number(inputVal)
				if(!isNaN(numberVal)) {
					setNumberVal(numberVal)						
				}			
			} // if input value is different than currently cached value
		}
	
	}) // focus out
	
}

function initNumberInputRecordEditBehavior($container,componentContext,recordProxy, numberInputObjectRef) {
	
	var $numberInputInput = $container.find("input")

	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoNumberInput
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