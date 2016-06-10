function loadRecordIntoDatePicker(datePickerElem, recordRef) {
	
	console.log("loadRecordIntoDatePicker: loading record into date picker: " + JSON.stringify(recordRef))
	
	var datePickerObjectRef = datePickerElem.data("objectRef")
	var datePickerFieldID = datePickerObjectRef.properties.fieldID
	
	console.log("loadRecordIntoDatePicker: Field ID to load data:" + datePickerFieldID)


	var datePickerContainerID = datePickerObjectRef.datePickerID
	var datePickerID = datePickerElemIDFromContainerElemID(datePickerContainerID)
	var datePickerSelector = '#' + datePickerID;
		var datePickerInputID = datePickerInputIDFromContainerElemID(datePickerContainerID)
		var datePickerInputSelector = '#'+datePickerInputID
	
	// Populate the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(datePickerFieldID)) {

		// If record has a value for the current container's associated field ID.
		var fieldVal = recordRef.fieldValues[datePickerFieldID]
		
		// The jQuery UI date picker only supports dates. So, until the Bootstrap datetime picker can be
		// integrated, only the date will be formatted and shown in the input field.
		var dateVal = moment(fieldVal).format("MM/DD/YYYY")

		console.log("loadRecordIntoDatePicker: Load value into container: " + $(datePickerElem).attr("id") + " field ID:" + 
					datePickerFieldID + "  value:" + fieldVal)

		var currDateVal = $(datePickerInputSelector).val()
		
		if(currDateVal != dateVal) {
			$(datePickerInputSelector).val(dateVal)
		}
		
	} else {
		// There's no value in the current record for this field, so clear the value in the container
		$(datePickerInputSelector).val("") 
	}
	
	
}


function initDatePickerRecordEditBehavior(datePickerObjectRef) {
	
	var datePickerContainerID = datePickerObjectRef.datePickerID
	var datePickerID = datePickerElemIDFromContainerElemID(datePickerContainerID)
	console.log("initDatePickerRecordEditBehavior: container ID =  " +datePickerContainerID + ' date picker ID = '+ datePickerID)
	
	var datePickerContainer = $('#'+datePickerContainerID)
	var datePickerSelector = '#'+datePickerID

	datePickerContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoDatePicker
	})


	var datePickerInputID = datePickerInputIDFromContainerElemID(datePickerContainerID)
	$('#'+datePickerInputID).datepicker()
	var fieldRef = getFieldRef(datePickerObjectRef.properties.fieldID)
	if(fieldRef.isCalcField) {
		$(datePickerSelector).data("DateTimePicker").disable()
		return;  // stop initialization, the check box is read only.
	}

	// Bootstrap datetime control is not ready for integration - it's conflicting with Semantic UI
	// $(datePickerSelector).on("dp.change", function(e) { }
		
	datePickerContainer.change(function () {
	    console.log("date picker changed dates")
		// Get the most recent copy of the object reference. It could have changed between
		// initialization time and the time the checkbox was changed.
		var containerID = datePickerObjectRef.datePickerID
		var objectRef = getElemObjectRef(containerID)
		var datePickerSelector = '#'+datePickerElemIDFromContainerElemID(containerID)
		
		var datePickerInputID = datePickerInputIDFromContainerElemID(containerID)
		
		var inputVal = $('#'+datePickerInputID).val()
		var dateVal = new Date(inputVal)
		var dateParam = moment(dateVal).toISOString()
		console.log("Date picker change value: input val = " + inputVal + " param="+dateParam)
		
		currRecordRef = currRecordSet.currRecordRef()
		var setRecordValParams = {
			parentTableID:viewFormContext.tableID,
			recordID:currRecordRef.recordID, 
			fieldID:objectRef.properties.fieldID, 
			value:dateParam }
		console.log("Setting date value: " + JSON.stringify(setRecordValParams))
		
		jsonAPIRequest("recordUpdate/setTimeFieldValue",setRecordValParams,function(updatedRecordRef) {
			
			// After updating the record, the local cache of records in currentRecordSet will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			currRecordSet.updateRecordRef(updatedRecordRef)
			// After changing the value, some of the calculated fields may have changed. For this
			// reason, it is necessary to reload the record into the layout/form, so the most
			// up to date values will be displayed.
			loadCurrRecordIntoLayout()
		}) // set record's text field value
		
	});	
	
}