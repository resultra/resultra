function loadRecordIntoTextBox(textBoxElem, recordRef) {
	
	console.log("loadRecordIntoTextBox: loading record into text box: " + JSON.stringify(recordRef))
	
	var textBoxObjectRef = textBoxElem.data("objectRef")
	
	if(textBoxObjectRef.properties.linkedValType == linkedComponentValTypeField) {
		// text box is linked to a field value
		var textBoxFieldID = textBoxObjectRef.properties.fieldID
	
		console.log("loadRecordIntoTextBox: Field ID to load data:" + textBoxFieldID)
	
		// In other words, we are populating the "intersection" of field values in the record
		// with the fields shown by the layout's containers.
		if(recordRef.fieldValues.hasOwnProperty(textBoxFieldID)) {

			var fieldVal = recordRef.fieldValues[textBoxFieldID]

			console.log("loadRecordIntoTextBox: Load value into container: " + $(this).attr("id") + " field ID:" + 
						textBoxFieldID + "  value:" + fieldVal)

			textBoxElem.find('input').val(fieldVal)
		} // If record has a value for the current container's associated field ID.
		else
		{
			textBoxElem.find('input').val("") // clear the value in the container
		}
		
	} else {
		// Text box is linked to a global value
		console.log("text box: skipping initialization for global value")
	}
	
	
	
}


function initTextBoxRecordEditBehavior(textFieldObjectRef) {
	
	var container = $('#'+textFieldObjectRef.textBoxID)

	container.data("viewFormConfig", {
		loadRecord: loadRecordIntoTextBox
	})
	
	if(textFieldObjectRef.properties.linkedValType == "field") {
		
		var fieldRef = getFieldRef(textFieldObjectRef.properties.fieldID)
		if(fieldRef.isCalcField) {
			container.find('input').prop('disabled',true);
			return;  // stop initialization, the text box is read only.
		}
	
	
	
		container.focusout(function () {
			var inputVal = container.find("input").val()
		
			// TODO - get edit information from single "objectRef", rather
			// than a scattering of different data values.
			var containerID = container.attr("id")
		
			var currTextObjRef = getElemObjectRef(containerID)
		
			var fieldID = currTextObjRef.properties.fieldID
			var fieldRef = getFieldRef(fieldID)
			var fieldType = fieldRef.type
			console.log("Text Box focus out:" 
			    + " containerID: " + containerID
				+ " ,fieldID: " + fieldID
			    + " ,fieldType: " + fieldType
				+ " , inputval:" + inputVal)
		
			currRecordRef = currRecordSet.currRecordRef()
			if(currRecordRef != null) {
			
				// Only update the value if it has changed. Sometimes a user may focus on or tab
				// through a field but not change it. In this case we don't need to update the record.
				if(currRecordRef.fieldValues[fieldID] != inputVal) {
				
					if(fieldType == "text") {
						currRecordRef.fieldValues[fieldID] = inputVal
						var setRecordValParams = { 
							parentTableID:viewFormContext.tableID,
							recordID:currRecordRef.recordID, 
							fieldID:fieldID, value:inputVal }
						jsonAPIRequest("recordUpdate/setTextFieldValue",setRecordValParams,function(replyData) {
							// After updating the record, the local cache of records in currentRecordSet will
							// be out of date. So after updating the record on the server, the locally cached
							// version of the record also needs to be updated.
							currRecordSet.updateRecordRef(replyData)
							// After changing the value, some of the calculated fields may have changed. For this
							// reason, it is necessary to reload the record into the layout/form, so the most
							// up to date values will be displayed.
							loadCurrRecordIntoLayout()
						}) // set record's text field value
					
					} else if (fieldType == "number") {
						var numberVal = Number(inputVal)
						if(!isNaN(numberVal)) {
							console.log("Change number val: "
								+ "fieldID: " + fieldID
							    + " ,number = " + numberVal)
							currRecordRef.fieldValues[fieldID] = numberVal
							var setRecordValParams = { 
								parentTableID:viewFormContext.tableID,
								recordID:currRecordRef.recordID, 
								fieldID:fieldID, 
								value:numberVal }
							jsonAPIRequest("recordUpdate/setNumberFieldValue",setRecordValParams,function(replyData) {
								// After updating the record, the local cache of records in currentRecordSet will
								// be out of date. So after updating the record on the server, the locally cached
								// version of the record also needs to be updated.
								currRecordSet.updateRecordRef(replyData)
							
								// After changing the value, some of the calculated fields may have changed. For this
								// reason, it is necessary to reload the record into the layout/form, so the most
								// up to date values will be displayed.
								loadCurrRecordIntoLayout()
							}) // set record's number field value
						}
					
					}
				
			
				} // if input value is different than currently cached value
			
			
			}
		
		}) // focus out
		
	} else { 
		console.log("Skipping text box initialization: global value")
	}

	
	
}