// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.


function initDateValidationProperties(params) {
	
	var $dateValidationForm = $('#datePickerValidationPropsForm')
	var $validationTypeSelection = $('#datePickerValidationSelection')
	var $rangeControls = $dateValidationForm.find('.datePickerValidationStartEndDateControls')
	var $compareDateInput = $dateValidationForm.find('.dateCompareVal')
	var $dateParamInput = $dateValidationForm.find('.dateInputParam')
	var $compareInput = $dateValidationForm.find('.dateCompareVal')
	
	var $startDatePicker = $dateValidationForm.find(".datePickerValidationRangeStartInput")
	var $endDatePicker = $dateValidationForm.find(".datePickerValidationRangeEndInput")
	var $compareDatePicker = $dateValidationForm.find('.dateCompareVal')
	
	function getValidationConfig() {
		var validationType = $validationTypeSelection.val()
		
		var validationConfig = null
		
		switch(validationType) {
		case "none":
		case "required":
		case "future":
		case "past":
			validationConfig = { rule: validationType }
			break;
		case "before":
		case "after":
			var compareDate = $compareDatePicker.data("DateTimePicker").date()
			if (compareDate === null) { return null }
			var compareDateUTC = compareDate.utc()
			validationConfig = {
				rule: validationType,
				compareDate: compareDateUTC
			}
			break;		
		case "between":
			var startDate = $startDatePicker.data("DateTimePicker").date()
			if (startDate === null) { return null }
			var startDateUTC = startDate.utc()
			var endDate = $endDatePicker.data("DateTimePicker").date()
			if (endDate === null) { return null }
			var endDateUTC = endDate.utc()
			validationConfig = {
				rule: validationType,
				startDate: startDateUTC,
				endDate: endDateUTC
			}
			
			break;
		}
		
		return validationConfig		
		
	}
	
	function updateValidationConfig() {
		var validationConfig = getValidationConfig()
		if (validationConfig !== null) {
			console.log("Validation config changed: " + JSON.stringify(validationConfig))
			params.setValidationConfig(validationConfig)
		
		}
	}
	
	function configureControlsForValidationType(validationType) {
		switch(validationType) {
		case "none":
		case "required":
		case "future":
		case "past":
			$dateParamInput.hide()
			break;
		case "before":
		case "after":
			$rangeControls.hide()
			$compareInput.show()
			break;		
		case "between":
			$compareInput.hide()
			$rangeControls.show()
			break;
		}			
	}
	
	// Initialize the start and end date controls
	var datePickerConfig = {
		showClose:true,
		format: getDateComponentDateTimeMomentFormat("date")
	}
	$startDatePicker.datetimepicker(datePickerConfig)
	$endDatePicker.datetimepicker(datePickerConfig)
	$compareDatePicker.datetimepicker(datePickerConfig)
	// Link the start and end date controls based to ensure
    // the range is preserved.
    $startDatePicker.on("dp.change", function (e) {
		console.log("Custom start date changed: " + e.date)
        $endDatePicker.data("DateTimePicker").minDate(e.date);
		updateValidationConfig()
    });
    $endDatePicker.on("dp.change", function (e) {
		console.log("Custom end date changed: " + e.date)
        $startDatePicker.data("DateTimePicker").maxDate(e.date);
		updateValidationConfig()
    });
    $compareDatePicker.on("dp.change", function (e) {
		updateValidationConfig()
	})
	
	var defaultValidationType = "required"
	$validationTypeSelection.val(defaultValidationType)
	configureControlsForValidationType(defaultValidationType)
	
	initSelectControlChangeHandler($validationTypeSelection,function(newValidationType) {
		configureControlsForValidationType(newValidationType)
		updateValidationConfig()
		//updateValidationSettingsIfValid()
	})
	
	
}
