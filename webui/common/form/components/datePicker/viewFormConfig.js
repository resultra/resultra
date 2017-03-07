
function getDateComponentDateTimeFormat() {
	// Populate with multiple formats which are selectable as an option for this component.
	var formatMap = {
		"dateOnly": 'MM/DD/YYYY',
		"dateTime": 'YYYY-MM-DD hh:mma' // Date with 12 hour time
	}
	var dateTimeFormat = formatMap['dateTime']
	return dateTimeFormat
}

function loadRecordIntoDatePicker($datePicker, recordRef) {
	
	console.log("loadRecordIntoDatePicker: loading record into date picker: " + JSON.stringify(recordRef))
	
	var datePickerObjectRef = $datePicker.data("objectRef")
	var datePickerFieldID = datePickerObjectRef.properties.fieldID
	
	
	console.log("loadRecordIntoDatePicker: Field ID to load data:" + datePickerFieldID)

	
	var $datePickerInput = datePickerInputFromContainer($datePicker)

	if(recordRef.fieldValues.hasOwnProperty(datePickerFieldID)) {

		// If record has a value for the current container's associated field ID.
		var fieldVal = recordRef.fieldValues[datePickerFieldID]
	
		// The jQuery UI date picker only supports dates. So, until the Bootstrap datetime picker can be
		// integrated, only the date will be formatted and shown in the input field.
		var dateVal = moment(fieldVal).format(getDateComponentDateTimeFormat())

		var currDateVal = $datePickerInput.val()
	
		if(currDateVal != dateVal) {
			$datePickerInput.val(dateVal)
		}
	
	} else {
		// There's no value in the current record for this field, so clear the value in the container
		$datePickerInput.val("") 
	}
	
}


function initDatePickerFieldEditBehavior(componentContext,recordProxy, datePickerObjectRef,$datePickerContainer) {
	

	var $datePickerInput = datePickerInputFromContainer($datePickerContainer)
	
	var fieldID = datePickerObjectRef.properties.fieldID
	
	var fieldRef = getFieldRef(fieldID)
	if(fieldRef.isCalcField) {
		$(datePickerSelector).data("DateTimePicker").disable()
		return;  // stop initialization, the check box is read only.
	}

	// Bootstrap datetime control is not ready for integration - it's conflicting with Semantic UI
	// $(datePickerSelector).on("dp.change", function(e) { }
	
	$datePickerContainer.find(".datePickerInputContainer").click(function(e) {
		
		// This is important - if a click hits an object, then stop the propagation of the click
		// to the parent div(s), including the canvas itself. If the parent canvas
		// gets a click, it will deselect all the items (see initObjectCanvasSelectionBehavior)
		e.stopPropagation();
	})
		
	$datePickerInput.on('dp.change',function (e) {
	    console.log("date picker changed dates")
		// Get the most recent copy of the object reference. It could have changed between
		// initialization time and the time the checkbox was changed.
		
		// The date passed with the event will be false if the date has been cleared.
		if(e.date !== false) {
			var objectRef = getContainerObjectRef($datePickerContainer)
				
			var dateParam = e.date.toISOString()
		
			currRecordRef = recordProxy.getRecordFunc()
		
			var dateValueFormat = {
				context: "datePicker",
				format: "date"
			}
			var setRecordValParams = {
				parentDatabaseID:currRecordRef.parentDatabaseID,
				recordID:currRecordRef.recordID, 
				changeSetID: recordProxy.changeSetID,
				fieldID:fieldID, 
				value:dateParam,
				 valueFormat: dateValueFormat}
			console.log("Setting date value: " + JSON.stringify(setRecordValParams))
		
			jsonAPIRequest("recordUpdate/setTimeFieldValue",setRecordValParams,function(updatedRecordRef) {
			
				// After updating the record, the local cache of records in currentRecordSet will
				// be out of date. So after updating the record on the server, the locally cached
				// version of the record also needs to be updated.
				recordProxy.updateRecordFunc(updatedRecordRef)
			}) // set record's text field value
			
		}
		
		
	});	
	
}


function initDatePickerRecordEditBehavior($datePickerContainer, componentContext,recordProxy, datePickerObjectRef) {
		
	var dateTimeFormat = getDateComponentDateTimeFormat()
		
		
	var $datePickerInput = datePickerInputFromContainer($datePickerContainer)
	$datePickerInput.datetimepicker({
		format: dateTimeFormat, // moment format - time will be displayed if the format includes a time component
		showTodayButton: true, // show the today button as a default
		keepOpen: false, // keep the popup open after selecting a date
	    daysOfWeekDisabled: [0,6], // Disable weekends
		showClear:true, // Show the button to clear the date (trash icon)
		sideBySide: false, // When editing both date and time, show the time side-by-side with the date.
		inline: false, // Display the date & time picker inline without the need to of an input field
		stepping: 5, // number of minutes the up and down arrows will step in clicking the up and down arrows.
		useCurrent: false, // When the date picker is shown, set the picker to the current date/time
		// daysOfWeekHighlighted: "1,2,3,4,5" // highlight weekdays
		// 
	})
	
	$datePickerContainer.data("viewFormConfig", {
		loadRecord: loadRecordIntoDatePicker
	})
	initDatePickerFieldEditBehavior(componentContext,recordProxy, datePickerObjectRef,$datePickerContainer)
	
}