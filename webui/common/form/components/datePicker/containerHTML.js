

function datePickerInputFromContainer($datePickerContainer) {
	return 	$datePickerContainer.find(".datePickerComponentInput")
}

function datePickerControlHTML() {
	return '<div class="datePickerInputContainer">'+
				'<div class="input-group date datePickerComponentInput">' +
					'<input type="text" name="symbol"  class="form-control " placeholder="">' +
					'<span class="input-group-addon datePickerCalendarButton">' +
                 	   '<span class=" glyphicon glyphicon-calendar"></span>' +
                	'</span>' +
					clearValueButtonHTML("datePickerComponentClearValueButton") +
				'</div>'+
			'</div>';
}

function datePickerContainerHTML(elementID)
{	
	var containerHTML = ''+
	'<div class="layoutContainer datePickerContainer datePickerFormContainer">' +
		'<div class="form-group">'+
			'<label>Date Picker</label>'+
			datePickerControlHTML() +
		'</div>'+
	'</div>';
	
	return containerHTML
}

function datePickerTableViewCellContainerHTML() {
	return ''+
		'<div class="layoutContainer datePickerContainer datePickerTableCellContainer">' +
			'<div class="form-group">' +
				datePickerControlHTML() +
			'</div>'+
		'</div>';
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
		allowInputToggle:false,
		format: momentFormat, // moment format - time will be displayed if the format includes a time component
		showTodayButton: true, // show the today button as a default
		keepOpen: false, // keep the popup open after selecting a date
	    daysOfWeekDisabled: [0,6], // Disable weekends
		showClear:true, // Show the button to clear the date (trash icon)
		sideBySide: false, // When editing both date and time, show the time side-by-side with the date.
		inline: false, // Display the date & time picker inline without the need to of an input field
		stepping: 5, // number of minutes the up and down arrows will step in clicking the up and down arrows.
		useCurrent: false, // When the date picker is shown, set the picker to the current date/time
		widgetParent: 'body',
		widgetPositioning: {
			horizontal: 'left',
			vertical: 'bottom'
		}
		// daysOfWeekHighlighted: "1,2,3,4,5" // highlight weekdays
		// 
	})
	
	
	
	// To ensure the date/time picker shows on top of other parts of the page, it needs to be attached to the body and 
	// repositioned near the input element. See: https://github.com/Eonasdan/bootstrap-datetimepicker/issues/790 
	$datePickerInput.on('dp.show', function() {
	      var datepicker = $('body').find('.bootstrap-datetimepicker-widget:last');
	      if (datepicker.hasClass('bottom')) {
	        var top = $(this).offset().top + $(this).outerHeight();
	        var left = $(this).offset().left;
	        datepicker.css({
	          'top': top + 'px',
	          'bottom': 'auto',
	          'left': left + 'px',
			   'z-index': 99999999 // needed for when date picker shown in bootstrap dialog or popup
	        });
	      } else if (datepicker.hasClass('top')) {
	        var top = $(this).offset().top - datepicker.outerHeight();
	        var left = $(this).offset().left;
	        datepicker.css({
	          'top': top + 'px',
	          'bottom': 'auto',
	          'left': left + 'px',
			  'z-index': 99999999
	        });
	      } 
		  
		  $datePickerInput.resize(function() {
			  console.log("Date picker resized/moved")
		  })
	  });
}

function setDatePickerFormComponentDate($datePicker, datePickerRef, momentDate) {
	
	var $datePickerInput = datePickerInputFromContainer($datePicker)	
	$datePickerInput.data("DateTimePicker").date(momentDate)
}

function getDatePickerFormComponentUTCDate($datePicker) {
	var $datePickerInput = datePickerInputFromContainer($datePicker)
	var datePicker = $datePickerInput.data("DateTimePicker")
	
	var currDate =  datePicker.date()
	if (currDate != null) {
		var dateParam = currDate.toISOString()
		return dateParam
	} else {
		return null
	}
}

function setDatePickerComponentLabel($datePickerContainer,datePickerRef) {
	var $label = $datePickerContainer.find('label')
	
	setFormComponentLabel($label,datePickerRef.properties.fieldID,
			datePickerRef.properties.labelFormat)
}

function datePickerComponentIsDisabled($datePickerContainer) {
	var $datePickerInput = datePickerInputFromContainer($datePickerContainer)
	var disabled = $datePickerInput.prop("disabled")
	return disabled
	
}