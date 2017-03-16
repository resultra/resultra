
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
	
	toggleFormulaEditorForField(datePickerRef.properties.fieldID)
	
}