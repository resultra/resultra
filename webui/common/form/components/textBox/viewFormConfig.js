

function formatTextBoxVal(componentLink, componentContext, rawInputVal, format) {
	
	if(componentLink.linkedValType == linkedComponentValTypeField) {
		var fieldID = componentLink.fieldID
		var fieldRef = getFieldRef(fieldID)
		var fieldType = fieldRef.type
		if(fieldType == fieldTypeNumber) {
			return formatNumberValue(format,rawInputVal)
		} else {
			return rawInputVal
		}
	} else {
		var globalInfo = componentContext.globalsByID[componentLink.globalID]
		if(globalInfo.type == globalTypeNumber) {
			return formatNumberValue(format,rawInputVal)
		} else {
			return rawInputVal
		}
	}
}



function loadRecordIntoTextBox($textBoxContainer, recordRef) {
	
	console.log("loadRecordIntoTextBox: loading record into text box: " + JSON.stringify(recordRef))
	
	var textBoxObjectRef = $textBoxContainer.data("objectRef")
	var $textBoxInput = $textBoxContainer.find('input')
	var componentContext = $textBoxContainer.data("componentContext")
	
	var componentLink = textBoxObjectRef.properties.componentLink
	
	function setRawInputVal(rawVal) { $textBoxInput.data("rawVal",rawVal) }
	
	
	if(componentLink.linkedValType == linkedComponentValTypeField) {
		// text box is linked to a field value
		var textBoxFieldID = componentLink.fieldID
	
		console.log("loadRecordIntoTextBox: Field ID to load data:" + textBoxFieldID)
	
		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(textBoxFieldID)) {

			var rawFieldVal = recordRef.fieldValues[textBoxFieldID]

			console.log("loadRecordIntoTextBox: Load value into container: " + $(this).attr("id") + " field ID:" + 
						textBoxFieldID + "  value:" + rawFieldVal)
			
			setRawInputVal(rawFieldVal)
			
			var formattedVal = formatTextBoxVal(componentLink,componentContext,
					rawFieldVal,textBoxObjectRef.properties.valueFormat.format)

			$textBoxInput.val(formattedVal)
		} // If record has a value for the current container's associated field ID.
		else
		{
			$textBoxInput.val("") // clear the value in the container
			setRawInputVal("")
		}
		
	} else {
		// Text box is linked to a global value
		
		var textBoxGlobalID = componentLink.globalID
		if(textBoxGlobalID in currGlobalVals) {
			var rawGlobalVal = currGlobalVals[textBoxGlobalID]
			
			$textBoxInput.val(rawGlobalVal)
			
			setRawInputVal(rawGlobalVal)
		}
		else
		{
			$textBoxInput.val("") // clear the value in the container
			setRawInputVal("")
		}		
	}
	
	
	
}

function initTextBoxFieldEditBehavior(componentContext, changeSetID, $container,$textBoxInput, 
		getRecordFunc, updateRecordFunc, textFieldObjectRef) {
	
	var componentLink = textFieldObjectRef.properties.componentLink
	
	var fieldRef = getFieldRef(componentLink.fieldID)
	if(fieldRef.isCalcField) {
		$textBoxInput.prop('disabled',true);
		return;  // stop initialization, the text box is read only.
	}
	
	var fieldType = fieldRef.type
		
	if(fieldType == fieldTypeNumber) {
		$textBoxInput.focusin(function() {
			// When focusing on the text input box, replaced the formatted value with 
			// the raw input value.
			var rawInputVal = $textBoxInput.data("rawVal")
			console.log("Focus in for number field: raw value for editing: " + rawInputVal)
			$textBoxInput.val(rawInputVal)
		})
	}
	

	$textBoxInput.focusout(function () {

		var currTextObjRef = getContainerObjectRef($container)		
		var componentLink = currTextObjRef.properties.componentLink
		var fieldID = componentLink.fieldID
		var fieldRef = getFieldRef(fieldID)
		var fieldType = fieldRef.type
		console.log("Text Box focus out:" 
			+ " ,fieldID: " + fieldID
		    + " ,fieldType: " + fieldType
			+ " , inputval:" + inputVal)

		// Retrieve the "raw input" value entered by the user and 
		// update the "rawVal" data setting on the text box.
		var inputVal = $textBoxInput.val()
		$textBoxInput.data("rawVal",inputVal)
		
		// Now that entry of the raw value is complete, revert the 
		// displayed value back to the format specified for the text box.
		var formattedVal = formatTextBoxVal(componentLink,componentContext,
						inputVal,currTextObjRef.properties.valueFormat.format)
		$textBoxInput.val(formattedVal)
		
		var currRecordRef = getRecordFunc()
			
			
		if(currRecordRef != null) {
		
			// Only update the value if it has changed. Sometimes a user may focus on or tab
			// through a field but not change it. In this case we don't need to update the record.
			if(currRecordRef.fieldValues[fieldID] != inputVal) {
			
				if(fieldType == fieldTypeText) {
					currRecordRef.fieldValues[fieldID] = inputVal
					
					var textBoxTextValueFormat = {
						context:"textBox",
						format:"general"
					}
					
					var setRecordValParams = { 
						parentDatabaseID:currRecordRef.parentDatabaseID,
						recordID:currRecordRef.recordID, 
						changeSetID: changeSetID,
						fieldID:fieldID, value:inputVal,
						 valueFormat: textBoxTextValueFormat }
					jsonAPIRequest("recordUpdate/setTextFieldValue",setRecordValParams,function(replyData) {
						// After updating the record, the local cache of records will
						// be out of date. So after updating the record on the server, the locally cached
						// version of the record also needs to be updated.
						updateRecordFunc(replyData)
						
					}) // set record's text field value
				
				} else if (fieldType == fieldTypeNumber) {
					var numberVal = Number(inputVal)
					if(!isNaN(numberVal)) {
						console.log("Change number val: "
							+ "fieldID: " + fieldID
						    + " ,number = " + numberVal)
						currRecordRef.fieldValues[fieldID] = numberVal
						var textBoxNumberValueFormat = {
							context:"textBox",
							format:"general"
						}
						var setRecordValParams = { 
							parentDatabaseID:currRecordRef.parentDatabaseID,
							recordID:currRecordRef.recordID,
							changeSetID: changeSetID,
							fieldID:fieldID, 
							value:numberVal,
							 valueFormat:textBoxNumberValueFormat
						}
						jsonAPIRequest("recordUpdate/setNumberFieldValue",setRecordValParams,function(replyData) {
							// After updating the record, the local cache of records will
							// be out of date. So after updating the record on the server, the locally cached
							// version of the record also needs to be updated.
							updateRecordFunc(replyData)
							
						}) // set record's number field value
					}
				
				}
			
		
			} // if input value is different than currently cached value
		
		
		}
	
	}) // focus out
	
}

function initTextBoxGlobalValBehavior(componentContext,$textBoxInput, textFieldObjectRef) {
	
	var componentLink = textFieldObjectRef.properties.componentLink
	
	var globalInfo = componentContext.globalsByID[componentLink.globalID]
		
	$textBoxInput.focusout(function () {
		var inputVal = $textBoxInput.val()
		if(globalInfo.type == globalTypeNumber) {
			var numberVal = Number(inputVal)
			if(!isNaN(numberVal)) {
				var setGlobalValParams = {
					parentDatabaseID: componentContext.databaseID,
					globalID: componentLink.globalID,
					value: numberVal
				}
				console.log("Setting global value (number): " + JSON.stringify(setGlobalValParams))
				jsonAPIRequest("global/setNumberValue",setGlobalValParams,function(replyData) {
				})
				
			}
		} else {
			var setGlobalValParams = {
				parentDatabaseID: componentContext.databaseID,
				globalID: componentLink.globalID,
				value: inputVal
			}
			console.log("Setting global value (text): " + JSON.stringify(setGlobalValParams))
			jsonAPIRequest("global/setTextValue",setGlobalValParams,function(replyData) {
			})	
		}
	})
}

function initTextBoxRecordEditBehavior($container,componentContext,changeSetID,
	getCurrentRecordFunc, updateCurrentRecordFunc, textFieldObjectRef) {
	
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
	
	
	var componentLink = textFieldObjectRef.properties.componentLink
	
	if(componentLink.linkedValType == linkedComponentValTypeField) {
		initTextBoxFieldEditBehavior(componentContext, changeSetID, $container,$textBoxInput,
				getCurrentRecordFunc, updateCurrentRecordFunc, textFieldObjectRef)
		
	} else { 
		assert(componentLink.linkedValType == linkedComponentValTypeGlobal)
		initTextBoxGlobalValBehavior(componentContext,$textBoxInput,textFieldObjectRef)
	}

	
	
}