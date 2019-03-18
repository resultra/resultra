// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
	
	var helpPopupParams = {
		initialMsg: datePickerCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: datePickerCol.parentTableID,
				datePickerID: datePickerCol.datePickerID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/datePicker/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	var conditionalFormatParams = {
		initialFormats: datePickerCol.properties.conditionalFormats,
		setConditionalFormats: function(formats) {
			var params = {
				parentTableID: datePickerCol.parentTableID,
				datePickerID: datePickerCol.datePickerID,
				conditionalFormats: formats
			}
			jsonAPIRequest("tableView/datePicker/setConditionalFormats",params,function(updateCol) {
			})	
		}
	}
	initDateConditionalFormatPropertyPanel(conditionalFormatParams)
	
	
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