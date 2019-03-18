// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initTextSelectionColPropertiesImpl(pageContext,textSelectionCol) {
	
	setColPropsHeader(textSelectionCol)
	hideSiblingsShowOne("#textSelectionColProps")
	
	var elemPrefix = "textSelection_"
		
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: textSelectionCol.parentTableID,
			selectionID: textSelectionCol.selectionID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/textSelection/setLabelFormat", formatParams, function(updatedNumberInput) {
			setColPropsHeader(updatedNumberInput)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: textSelectionCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: textSelectionCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: textSelectionCol.parentTableID,
				selectionID: textSelectionCol.selectionID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/textSelection/setPermissions",params,function(updatedTextSelection) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	function saveValueList(valueListID) {
		var setValueListParams = {
			parentTableID: textSelectionCol.parentTableID,
			selectionID: textSelectionCol.selectionID,
			valueListID: valueListID
		}
		jsonAPIRequest("tableView/textSelection/setValueList", setValueListParams, function(updatedTextBox) {
		})			
	}
	var valueListPropertyParams = {
		elemPrefix: elemPrefix,	
		databaseID: pageContext.databaseID,
		saveValueListCallback: saveValueList,
		defaultValueListID: textSelectionCol.properties.valueListID
	}
	initValueListSelectionPropertyPanel(valueListPropertyParams)
	
	
	var validationParams = {
		initialValidationProps: textSelectionCol.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentTableID: textSelectionCol.parentTableID,
				selectionID: textSelectionCol.selectionID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("tableView/textSelection/setValidation",validationParams,function(updatedTextSelection) {
			})
		
		}
	}
	initTextInputValidationProperties(validationParams)
	
	var clearValueParams = {
		initialVal: textSelectionCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: textSelectionCol.parentTableID,
				selectionID: textSelectionCol.selectionID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/textSelection/setClearValueSupported",formatParams,function(updatedTextSelection) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: textSelectionCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: textSelectionCol.parentTableID,
				selectionID: textSelectionCol.selectionID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/textSelection/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
}

function initTextSelectionColProperties(pageContext,tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		selectionID: columnID
	}
	jsonAPIRequest("tableView/textSelection/get", getColParams, function(textSelectionCol) { 
		initTextSelectionColPropertiesImpl(pageContext,textSelectionCol)
	})
	
	
	
}