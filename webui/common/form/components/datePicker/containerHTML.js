

function datePickerInputFromContainer($datePickerContainer) {
	return 	$datePickerContainer.find(".datePickerComponentInput")
}


function datePickerContainerHTML(elementID)
{	
	
	var containerHTML = ''+
	'<div class="layoutContainer datePickerContainer">' +
		'<div class="form-group">'+
			'<label>New Field</label>'+
			'<div class="datePickerInputContainer">' + 
				'<input type="text" name="symbol"  class="form-control datePickerComponentInput" placeholder="Select a date">' +
			'</div>'+
		'</div>'+
	'</div>';
	
	return containerHTML
}


function getDateComponentDateTimeMomentFormat(dateFormat) {
	// Populate with multiple formats which are selectable as an option for this component.
	var formatMap = {
		"longDate": 'MM/DD/YYYY',
		"date": 'M/D/YY',
		"dateTime": 'M/D/YY h:mm A' // Date with 12 hour time
	}
	
	
	var dateTimeFormat = formatMap[dateFormat]
	if (dateFormat === undefined) {
		return formatMap["dateOnly"]
	}
	
	return dateTimeFormat
}



function initDatePickerFormComponentInput($datePickerContainer, datePickerRef) {
	
	var $datePickerInput = datePickerInputFromContainer($datePickerContainer)
	var momentFormat = getDateComponentDateTimeMomentFormat(datePickerRef.properties.dateFormat)
	
	// Destroy the existing date picker (if it's present)
	var currDatePicker = $datePickerInput.data("DateTimePicker")
	if (currDatePicker !== undefined) {
		currDatePicker.destroy()
	}
	
	$datePickerInput.datetimepicker({
		format: momentFormat, // moment format - time will be displayed if the format includes a time component
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
}

function setDatePickerFormComponentDate($datePicker, datePickerRef, momentDate) {
	
	var $datePickerInput = datePickerInputFromContainer($datePicker)
	
	var dateFormat = getDateComponentDateTimeMomentFormat(datePickerRef.properties.dateFormat)
	var formattedDate = momentDate.format(dateFormat)
	
	var currDateVal = $datePickerInput.val()
	if(currDateVal !== formattedDate) {
		$datePickerInput.val(formattedDate)
	}
}