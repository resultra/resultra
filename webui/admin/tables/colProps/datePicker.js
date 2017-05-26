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