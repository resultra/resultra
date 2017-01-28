function loadRecordIntoDatePicker(datePickerElem, recordRef) {
	
	console.log("loadRecordIntoDatePicker: loading record into date picker: " + JSON.stringify(recordRef))
	
	var datePickerObjectRef = datePickerElem.data("objectRef")
	var componentLink = datePickerObjectRef.properties.componentLink
	
	var datePickerFieldID = componentLink.fieldID
	
	
	console.log("loadRecordIntoDatePicker: Field ID to load data:" + datePickerFieldID)


	var datePickerContainerID = datePickerObjectRef.datePickerID
	var datePickerID = datePickerElemIDFromContainerElemID(datePickerContainerID)
	var datePickerSelector = '#' + datePickerID;
		var datePickerInputID = datePickerInputIDFromContainerElemID(datePickerContainerID)
		var datePickerInputSelector = '#'+datePickerInputID
	
	if(componentLink.linkedValType == linkedComponentValTypeField) {
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
	} else {
		var datePickerGlobalID = componentLink.globalID
		if(datePickerGlobalID in currGlobalVals) {
			var globalVal = currGlobalVals[datePickerGlobalID]
			var dateVal = moment(globalVal).format("MM/DD/YYYY")
			$(datePickerInputSelector).val(dateVal)
		}
		else
		{
			$(datePickerInputSelector).val("") 
		}
		
	}
}

function getDataPickerDateVal(containerID) {
	var datePickerInputID = datePickerInputIDFromContainerElemID(containerID)
	
	var inputVal = $('#'+datePickerInputID).val()
	var dateVal = new Date(inputVal)
	var dateParam = moment(dateVal).toISOString()
	return dateParam
}

function initDatePickerFieldEditBehavior(componentContext,changeSetID,
		getRecordFunc, updateRecordFunc, datePickerObjectRef,$datePickerContainer) {
	
	
	var datePickerContainerID = datePickerObjectRef.datePickerID
	var datePickerID = datePickerElemIDFromContainerElemID(datePickerContainerID)
	
	var componentLink = datePickerObjectRef.properties.componentLink
	
	var fieldRef = getFieldRef(componentLink.fieldID)
	if(fieldRef.isCalcField) {
		$(datePickerSelector).data("DateTimePicker").disable()
		return;  // stop initialization, the check box is read only.
	}

	// Bootstrap datetime control is not ready for integration - it's conflicting with Semantic UI
	// $(datePickerSelector).on("dp.change", function(e) { }
		
	$datePickerContainer.change(function () {
	    console.log("date picker changed dates")
		// Get the most recent copy of the object reference. It could have changed between
		// initialization time and the time the checkbox was changed.
		var containerID = datePickerObjectRef.datePickerID
		var objectRef = getElemObjectRef(containerID)
		
		var componentLink = objectRef.properties.componentLink
		
		var dateParam = getDataPickerDateVal(containerID)
		
		currRecordRef = getRecordFunc()
		
		var dateValueFormat = {
			context: "datePicker",
			format: "date"
		}
		var setRecordValParams = {
			parentDatabaseID:currRecordRef.parentDatabaseID,
			recordID:currRecordRef.recordID, 
			changeSetID: changeSetID,
			fieldID:componentLink.fieldID, 
			value:dateParam,
			 valueFormat: dateValueFormat}
		console.log("Setting date value: " + JSON.stringify(setRecordValParams))
		
		jsonAPIRequest("recordUpdate/setTimeFieldValue",setRecordValParams,function(updatedRecordRef) {
			
			// After updating the record, the local cache of records in currentRecordSet will
			// be out of date. So after updating the record on the server, the locally cached
			// version of the record also needs to be updated.
			updateRecordFunc(updatedRecordRef)
		}) // set record's text field value
		
	});	
	
}

function initDatePickerGlobalEditBehavior(componentContext,datePickerObjectRef,$datePickerContainer) {
	$datePickerContainer.change(function () {
	    console.log("date picker changed dates (global)")
		// Get the most recent copy of the object reference. It could have changed between
		// initialization time and the time the checkbox was changed.
		var containerID = datePickerObjectRef.datePickerID
		var objectRef = getElemObjectRef(containerID)
		
		var componentLink = objectRef.properties.componentLink
				
		var dateParam = getDataPickerDateVal(containerID)
		
		
		var setGlobalValParams = {
			parentDatabaseID:componentContext.databaseID,
			globalID:componentLink.globalID, 
			value:dateParam }
		console.log("Setting date value (global): " + JSON.stringify(setGlobalValParams))
			
		jsonAPIRequest("global/setTimeValue",setGlobalValParams,function(updatedGlobalVal) {
		
			// TODO - Update the record set and global value
		}) // set record's text field value
		
		
	})
}



function initDatePickerRecordEditBehavior(componentContext,changeSetID,
		getRecordFunc, updateRecordFunc, datePickerObjectRef) {
	
	var datePickerContainerID = datePickerObjectRef.datePickerID
	var datePickerID = datePickerElemIDFromContainerElemID(datePickerContainerID)
	
	var datePickerInputID = datePickerInputIDFromContainerElemID(datePickerContainerID)
	$('#'+datePickerInputID).datepicker()
	
	
	console.log("initDatePickerRecordEditBehavior: container ID =  " +datePickerContainerID + ' date picker ID = '+ datePickerID)
	
	var $datePickerContainer = $('#'+datePickerContainerID)

	$datePickerContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoDatePicker
	})
	
	var componentLink = datePickerObjectRef.properties.componentLink
	
	if(componentLink.linkedValType == linkedComponentValTypeField) {
		initDatePickerFieldEditBehavior(componentContext,changeSetID,
					getRecordFunc, updateRecordFunc, datePickerObjectRef,$datePickerContainer)
		
	} else { 
		assert(componentLink.linkedValType == linkedComponentValTypeGlobal)
		initDatePickerGlobalEditBehavior(componentContext,datePickerObjectRef,$datePickerContainer)
	}
	
}