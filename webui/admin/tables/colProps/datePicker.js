function initNumberDatePickerColPropertiesImpl(datePickerCol) {
	setColPropsHeader(datePickerCol)
	
	var elemPrefix = "datePicker_"
	hideSiblingsShowOne("#datePickerColProps")
	
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: datePickerCol.parentTableID,
			datePickerID: datePickerCol.datePickerID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/datePicker/setLabelFormat", formatParams, function(updateDatePicker) {
			setColPropsHeader(updateDatePicker)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: datePickerCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: datePickerCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: datePickerCol.parentTableID,
				datePickerID: datePickerCol.datePickerID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/datePicker/setPermissions",params,function(updateDatePicker) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var formatParams = {
		initialFormat: datePickerCol.properties.dateFormat,
		setFormat: function(newFormat) {
			var formatParams = {
				parentTableID: datePickerCol.parentTableID,
				datePickerID: datePickerCol.datePickerID,
				dateFormat: newFormat
			}
			jsonAPIRequest("tableView/datePicker/setFormat",formatParams,function(updatedDatePicker) {
			})
		}
	}
	initDateFormatProperties(formatParams)
	
	
	var validationParams = {
		setValidationConfig: function(validationConfig) {
			var validationParams = {
				parentTableID: datePickerCol.parentTableID,
				datePickerID: datePickerCol.datePickerID,
				validation: validationConfig
			}
			jsonAPIRequest("tableView/datePicker/setValidation", validationParams, function(updatedDatePicker) {
			})			
		}
	
	}
	initDateValidationProperties(validationParams)
	
	
	var clearValueParams = {
		initialVal: datePickerCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: datePickerCol.parentTableID,
				datePickerID: datePickerCol.datePickerID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/datePicker/setClearValueSupported",formatParams,function(updatedDatePicker) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	
	
	
}


function initDatePickerColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		datePickerID: columnID
	}
	jsonAPIRequest("tableView/datePicker/get", getColParams, function(datePickerCol) { 
		initNumberDatePickerColPropertiesImpl(datePickerCol)
	})
}