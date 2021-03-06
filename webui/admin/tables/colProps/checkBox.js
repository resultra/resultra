// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initCheckBoxColPropertiesImpl(checkBoxCol) {
	
	setColPropsHeader(checkBoxCol)
	
	var elemPrefix = "checkBox_"
	hideSiblingsShowOne("#checkBoxColProps")
	
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: checkBoxCol.parentTableID,
			checkBoxID: checkBoxCol.checkBoxID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/checkBox/setLabelFormat", formatParams, function(updatedCheckBox) {
			setColPropsHeader(updatedCheckBox)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: checkBoxCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: checkBoxCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: checkBoxCol.parentTableID,
				checkBoxID: checkBoxCol.checkBoxID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/checkBox/setPermissions",params,function(updatedCheckBox) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	
	
	var formatParams = {
		initialColorScheme: checkBoxCol.properties.colorScheme,
		setColorScheme: function(newColorScheme) {
			var colorSchemeParams = {
				parentTableID: checkBoxCol.parentTableID,
				checkBoxID: checkBoxCol.checkBoxID,
				colorScheme: newColorScheme
			}
			console.log("Setting new color scheme: " + JSON.stringify(colorSchemeParams))
	
			jsonAPIRequest("tableView/checkBox/setColorScheme",colorSchemeParams,function(updatedCheckboxRef) {
			})
		},
		initialStrikethrough: false,
		setStrikethrough: function(strikethroughCompleted) {  } // no-op
	}
	initCheckBoxFormatProps(formatParams)
	
	
	var validationParams = {
		initialValidationConfig: checkBoxCol.properties.validation,
		setValidation: function(validationConfig) {
			var validationParams = {
				parentTableID: checkBoxCol.parentTableID,
				checkBoxID: checkBoxCol.checkBoxID,
				validation: validationConfig
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("tableView/checkBox/setValidation",validationParams,function(updatedCheckboxRef) {
			})
		
		}
	}
	initCheckBoxValidationProps(validationParams)	
	
	var clearValueParams = {
		initialVal: checkBoxCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: checkBoxCol.parentTableID,
				checkBoxID: checkBoxCol.checkBoxID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/checkBox/setClearValueSupported",formatParams,function(updatedCheckboxRef) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: checkBoxCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: checkBoxCol.parentTableID,
				checkBoxID: checkBoxCol.checkBoxID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/checkBox/setHelpPopupMsg",params,function(updatedCheckboxRef) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
}


function initCheckBoxColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		checkBoxID: columnID
	}
	jsonAPIRequest("tableView/checkBox/get", getColParams, function(checkBoxCol) { 
		initCheckBoxColPropertiesImpl(checkBoxCol)
	})
}