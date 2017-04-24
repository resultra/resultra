
function loadDatePickerProperties($container,datePickerRef) {
	console.log("Loading date picker properties")
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#datePickerProps')
	
	function initFormatProperties() {
		var $formatSelection = $('#adminDateComponentFormatSelection')
		$formatSelection.val(datePickerRef.properties.dateFormat)
		initSelectControlChangeHandler($formatSelection,function(newFormat) {
		
			var iconParams = {
				parentFormID: datePickerRef.parentFormID,
				datePickerID: datePickerRef.datePickerID,
				dateFormat: newFormat
			}
			jsonAPIRequest("frm/datePicker/setFormat",iconParams,function(updatedDatePicker) {
				setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)	
				initDatePickerFormComponentInput($container,updatedDatePicker)
				var sampleDate = moment("2013-02-08 09:30")
				setDatePickerFormComponentDate($container, updatedDatePicker, sampleDate)
			})
		
		})
		
	}
	initFormatProperties()
	
	function initValidationProperties() {
		
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
				var validationParams = {
					parentFormID: datePickerRef.parentFormID,
					datePickerID: datePickerRef.datePickerID,
					validation: validationConfig
				}
				jsonAPIRequest("frm/datePicker/setValidation", validationParams, function(updatedDatePicker) {
					setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)
				})	
			
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
	initValidationProperties()
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
				parentFormID: datePickerRef.parentFormID,
				datePickerID: datePickerRef.datePickerID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/datePicker/setLabelFormat", formatParams, function(updatedDatePicker) {
			setDatePickerComponentLabel($container,updatedDatePicker)
			setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)	
		})	
	}
	
	var elemPrefix = "datePicker_"
	
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: datePickerRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: datePickerRef.parentFormID,
			datePickerID: datePickerRef.datePickerID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/datePicker/setVisibility",params,function(updatedDatePicker) {
			setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)	
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: datePickerRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: datePickerRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: datePickerRef.parentFormID,
				datePickerID: datePickerRef.datePickerID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/datePicker/setPermissions",params,function(updatedDatePicker) {
				setContainerComponentInfo($container,updatedDatePicker,updatedDatePicker.datePickerID)	
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: datePickerRef.parentFormID,
		componentID: datePickerRef.datePickerID,
		componentLabel: 'date picker',
		$componentContainer: $container
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	toggleFormulaEditorForField(datePickerRef.properties.fieldID)
	
}