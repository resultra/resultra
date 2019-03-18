// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initFormButtonColPropertiesImpl(pageContext,formButtonCol) {
	
//	setColPropsHeader(formButtonCol)
	
	var elemPrefix = "button_"
	hideSiblingsShowOne("#formButtonColProps")
	
	function initColorSchemeProperties() {
		var $schemeSelection = $('#adminButtonComponentColorSchemeSelection')
		$schemeSelection.val(formButtonCol.properties.colorScheme)
		initSelectControlChangeHandler($schemeSelection,function(newScheme) {
		
			var sizeParams = {
				parentTableID: formButtonCol.parentTableID,
				buttonID: formButtonCol.buttonID,
				colorScheme: newScheme
			}
			jsonAPIRequest("tableView/formButton/setColorScheme",sizeParams,function(updatedButton) {
			})
		
		})
		
	}
	initColorSchemeProperties()	
	
	
	function initIconProperties() {
		var $iconSelection = $('#adminButtonComponentIconSelection')
		$iconSelection.val(formButtonCol.properties.icon)
		initSelectControlChangeHandler($iconSelection,function(newIcon) {
		
			var iconParams = {
				parentTableID: formButtonCol.parentTableID,
				buttonID: formButtonCol.buttonID,
				icon: newIcon
			}
			jsonAPIRequest("tableView/formButton/setIcon",iconParams,function(updatedButton) {
			})
		
		})
		
	}
	initIconProperties()
	
	function initButtonSizeProperties() {
		var $sizeSelection = $('#adminButtonComponentSizeSelection')
		$sizeSelection.val(formButtonCol.properties.size)
		initSelectControlChangeHandler($sizeSelection,function(newSize) {
		
			var sizeParams = {
				parentTableID: formButtonCol.parentTableID,
				buttonID: formButtonCol.buttonID,
				size: newSize
			}
			jsonAPIRequest("tableView/formButton/setSize",sizeParams,function(updatedButton) {
			})		
		})
		
	}
	initButtonSizeProperties()
	
	function saveBehaviorProperties(popupBehavior) {
		var setPopupBehaviorParams = { 
				parentTableID: formButtonCol.parentTableID,
				buttonID: formButtonCol.buttonID,
			popupBehavior: popupBehavior
		}
		jsonAPIRequest("tableView/formButton/setPopupBehavior",setPopupBehaviorParams,function(updatedButtonRef) {
		})
	}
	initFormButtonBehaviorProperties(formButtonCol,saveBehaviorProperties)
	
	var elemPrefix = "button_"
	
	var defaultValPropParams = {
		databaseID: pageContext.databaseID,
		elemPrefix: elemPrefix,
		defaultDefaultValues: formButtonCol.properties.defaultValues,
		updateDefaultValues: function(updatedDefaultVals) {
			console.log("Updating default values for form button: " + JSON.stringify(updatedDefaultVals))
			
			var setDefaultValsParams = {
				parentTableID: formButtonCol.parentTableID,
				buttonID: formButtonCol.buttonID,
				defaultValues: updatedDefaultVals }
			
			jsonAPIRequest("tableView/formButton/setDefaultVals",setDefaultValsParams,function(updatedButtonRef) {
			})
		}
	}
	initDefaultValuesPropertyPanel(defaultValPropParams)
	
	function saveButtonLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: formButtonCol.parentTableID,
			buttonID: formButtonCol.buttonID,
			buttonLabelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/formButton/setButtonLabelFormat", formatParams, function(updateRating) {
		})	
	}
	var buttonLabelParams = {
		elemPrefix: elemPrefix,
		initialVal: formButtonCol.properties.buttonLabelFormat,
		saveLabelPropsCallback: saveButtonLabelProps
	}
	initFormButtonLabelPropertyPanel(buttonLabelParams)
	
}


function initFormButtonColProperties(pageContext,tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		buttonID: columnID
	}
	jsonAPIRequest("tableView/formButton/get", getColParams, function(formButtonCol) { 
		initFormButtonColPropertiesImpl(pageContext,formButtonCol)
	})
}