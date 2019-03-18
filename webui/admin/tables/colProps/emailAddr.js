// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initEmailAddrColPropertiesImpl(emailAddrCol) {
	
	setColPropsHeader(emailAddrCol)
	hideSiblingsShowOne("#emailAddrColProps")
	
	var elemPrefix = "emailAddr_"
		
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: emailAddrCol.parentTableID,
			emailAddrID: emailAddrCol.emailAddrID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/emailAddr/setLabelFormat", formatParams, function(updatedNumberInput) {
			setColPropsHeader(updatedNumberInput)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: emailAddrCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: emailAddrCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: emailAddrCol.parentTableID,
				emailAddrID: emailAddrCol.emailAddrID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/emailAddr/setPermissions",params,function(updatedEmailAddr) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
		
	
	var validationParams = {
		initialValidationProps: emailAddrCol.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentTableID: emailAddrCol.parentTableID,
				emailAddrID: emailAddrCol.emailAddrID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("tableView/emailAddr/setValidation",validationParams,function(updatedEmailAddr) {
			})
		
		}
	}
	initTextInputValidationProperties(validationParams)
	
	var clearValueParams = {
		initialVal: emailAddrCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: emailAddrCol.parentTableID,
				emailAddrID: emailAddrCol.emailAddrID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/emailAddr/setClearValueSupported",formatParams,function(updatedEmailAddr) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: emailAddrCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: emailAddrCol.parentTableID,
				emailAddrID: emailAddrCol.emailAddrID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/emailAddr/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
}

function initEmailAddrColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		emailAddrID: columnID
	}
	jsonAPIRequest("tableView/emailAddr/get", getColParams, function(emailAddrCol) { 
		initEmailAddrColPropertiesImpl(emailAddrCol)
	})
	
	
	
}