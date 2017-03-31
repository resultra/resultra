
function loadRecordIntoTextBox($textBoxContainer, recordRef) {
	
	console.log("loadRecordIntoTextBox: loading record into text box: " + JSON.stringify(recordRef))
	
	var textBoxObjectRef = $textBoxContainer.data("objectRef")
	var $textBoxInput = $textBoxContainer.find('input')
	var componentContext = $textBoxContainer.data("componentContext")
	
	if(formComponentIsReadOnly(textBoxObjectRef.properties.permissions)) {
		$textBoxInput.prop('disabled',true);
	} else {
		$textBoxInput.prop('disabled',false);
		
	}
	
	
	function setRawInputVal(rawVal) { $textBoxInput.data("rawVal",rawVal) }

	// text box is linked to a field value
	var textBoxFieldID = textBoxObjectRef.properties.fieldID

	console.log("loadRecordIntoTextBox: Field ID to load data:" + textBoxFieldID)

	// In other words, we are populating the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(textBoxFieldID)) {

		var fieldVal = recordRef.fieldValues[textBoxFieldID]
		
		if(fieldVal === null) {
			$textBoxInput.val("")
		} else {
			$textBoxInput.val(fieldVal)
			
		}

	} // If record has a value for the current container's associated field ID.
	else
	{
		$textBoxInput.val("") // clear the value in the container
	}	
	
}

function initTextBoxFieldEditBehavior(componentContext, $container,$textBoxInput,recordProxy, textFieldObjectRef) {
	
	var textBoxFieldID = textFieldObjectRef.properties.fieldID
	
	var fieldRef = getFieldRef(textBoxFieldID)
	if(fieldRef.isCalcField) {
		$textBoxInput.prop('disabled',true);
		return;  // stop initialization, the text box is read only.
	}
	
	var fieldType = fieldRef.type
		
	function setTextVal(textVal) {
		var textBoxTextValueFormat = {
			context:"textBox",
			format:"general"
		}
		var currRecordRef = recordProxy.getRecordFunc()
		var setRecordValParams = { 
			parentDatabaseID:currRecordRef.parentDatabaseID,
			recordID:currRecordRef.recordID, 
			changeSetID: recordProxy.changeSetID,
			fieldID:textBoxFieldID, 
			value:textVal,
			valueFormat: textBoxTextValueFormat 
		}
		jsonAPIRequest("recordUpdate/setTextFieldValue",setRecordValParams,function(replyData) {
			// After updating the record, the local cache of records will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			recordProxy.updateRecordFunc(replyData)
			
		}) // set record's text field value
				
	}
	
	
	var $clearValueButton = $container.find(".textBoxComponentClearValueButton")
	initButtonControlClickHandler($clearValueButton,function() {
			console.log("Clear value clicked for text box")
		
		var currRecordRef = recordProxy.getRecordFunc()
		setTextVal(null)
		
	})
		

	$textBoxInput.focusout(function () {

		var currTextObjRef = getContainerObjectRef($container)		

		// Retrieve the "raw input" value entered by the user and 
		// update the "rawVal" data setting on the text box.
		var inputVal = $textBoxInput.val()
		console.log("Text Box focus out:" + inputVal)
		
		var currRecordRef = recordProxy.getRecordFunc()
					
		if(currRecordRef != null) {
		
			// Only update the value if it has changed. Sometimes a user may focus on or tab
			// through a field but not change it. In this case we don't need to update the record.
			if(currRecordRef.fieldValues[textBoxFieldID] != inputVal) {
					setTextVal(inputVal)			
			} // if input value is different than currently cached value
		}
	
	}) // focus out
	
}

function initTextBoxRecordEditBehavior($container,componentContext,recordProxy, textFieldObjectRef) {
	
	var $textBoxInput = $container.find("input")

	$container.data("viewFormConfig", {
		loadRecord: loadRecordIntoTextBox
	})
	
	$container.data("componentContext",componentContext)
	
	
	// When the user clicks on the text box input control, prevent the click from propagating higher.
	// This allows the user to change the values without selecting the form component itself.
	// The user can still select the component by clicking on the label or anywwhere outside
	// the input control.
	$textBoxInput.click(function (event){
		event.stopPropagation();
   	 	//   ... your code here
		return false;
	});
	initTextBoxFieldEditBehavior(componentContext, $container,$textBoxInput,
			recordProxy, textFieldObjectRef)
	
}