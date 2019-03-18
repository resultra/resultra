// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initUrlLinkColPropertiesImpl(urlLinkCol) {
	
	setColPropsHeader(urlLinkCol)
	hideSiblingsShowOne("#urlLinkColProps")
	
	var elemPrefix = "urlLink_"
		
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: urlLinkCol.parentTableID,
			urlLinkID: urlLinkCol.urlLinkID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/urlLink/setLabelFormat", formatParams, function(updatedNumberInput) {
			setColPropsHeader(updatedNumberInput)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: urlLinkCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: urlLinkCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: urlLinkCol.parentTableID,
				urlLinkID: urlLinkCol.urlLinkID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/urlLink/setPermissions",params,function(updatedUrlLink) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
		
	
	var validationParams = {
		initialValidationProps: urlLinkCol.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentTableID: urlLinkCol.parentTableID,
				urlLinkID: urlLinkCol.urlLinkID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("tableView/urlLink/setValidation",validationParams,function(updatedUrlLink) {
			})
		
		}
	}
	initTextInputValidationProperties(validationParams)
	
	var clearValueParams = {
		initialVal: urlLinkCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: urlLinkCol.parentTableID,
				urlLinkID: urlLinkCol.urlLinkID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/urlLink/setClearValueSupported",formatParams,function(updatedUrlLink) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: urlLinkCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: urlLinkCol.parentTableID,
				urlLinkID: urlLinkCol.urlLinkID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/urlLink/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
}

function initUrlLinkColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		urlLinkID: columnID
	}
	jsonAPIRequest("tableView/urlLink/get", getColParams, function(urlLinkCol) { 
		initUrlLinkColPropertiesImpl(urlLinkCol)
	})
	
	
	
}