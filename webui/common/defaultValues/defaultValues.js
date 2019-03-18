// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function getDefaultValListDefaultVals(elemPrefix) {
	var defaultVals = []
	
	var defaultValListSelector = createPrefixedSelector(elemPrefix,'DefaultValuesList')
	
	$(defaultValListSelector + " .defaultValuesPanelListItem").each(function() {
	
		var defaultValConfigFunc = $(this).data("defaultValConfigFunc")
		var defaultValConfig = defaultValConfigFunc()
		
		if(defaultValConfig != null) {
			defaultVals.push(defaultValConfig)
		}
	})

	console.log("getDefaultValListDefaultVals config: " + JSON.stringify(defaultVals))
	
	return defaultVals
}


function updateDefaultValues(panelParams) {
	console.log("updateDefaultValues: " + JSON.stringify(panelParams))

	var defaultVals = getDefaultValListDefaultVals(panelParams.elemPrefix)
	
	panelParams.updateDefaultValues(defaultVals)
}

function createDefaultValueListItem(panelParams,fieldName) {
	
	var $defaultValListItem = $('#defaultValuesPanelListItem').clone()
	$defaultValListItem.attr("id","")
	
	var $fieldLabel = $defaultValListItem.find(".defaultValuesListItemLabel")
	$fieldLabel.text(fieldName)
	
	var $deleteButton = $defaultValListItem.find(".defaultValuesListItemDeleteDefaultValButton")
	initButtonControlClickHandler($deleteButton,function() {
		$defaultValListItem.remove()
		updateDefaultValues(panelParams)
	})

	return $defaultValListItem
}
 

function boolDefaultValueItem(panelParams,fieldInfo,defaultDefaultValInfo) {
	
	var $defaultValControls = $('#defaultValueBoolFieldListItem').clone()
	$defaultValControls.attr("id","")
	
	var $defaultValSelection = $defaultValControls.find("select")
	$defaultValSelection.empty()
	$defaultValSelection.append(defaultSelectOptionPromptHTML("Select a default value"))
	
	var defaultValuesBool = {
		"true": {
			label: "True",
		},
		"false": {
			label: "False",
		},
	}
	
	for(var defaultValID in defaultValuesBool) {
	 	var selectDefaultValHTML = selectOptionHTML(defaultValID, defaultValuesBool[defaultValID].label)
	 	$defaultValSelection.append(selectDefaultValHTML)				
	}	
	
	if(defaultDefaultValInfo !== null) {
		$defaultValSelection.val(defaultDefaultValInfo.defaultValueID)
	}	
		
	initSelectControlChangeHandler($defaultValSelection,function(defaultValID) {
		var defaultValInfo = defaultValuesBool[defaultValID]
		console.log("Default value selection change: " + defaultValID)
		updateDefaultValues(panelParams)
	})
		
		
	var $defaultValListItem = createDefaultValueListItem(panelParams,fieldInfo.name)
		
	$defaultValListItem.data("defaultValConfigFunc",function() {
		var defaultValID = $defaultValSelection.val()
		if(defaultValID !== null && defaultValID.length > 0) {
			var conditions = []
			
			var defaultValConfig = { fieldID: fieldInfo.fieldID,
				defaultValueID: defaultValID }	
			return defaultValConfig
		} else {
			return null
		}
	})
		
	$defaultValListItem.append($defaultValControls)
		
	return $defaultValListItem
	
}


function timeDefaultValueItem(panelParams,fieldInfo,defaultDefaultValInfo) {
	
	var $defaultValControls = $('#defaultValueTimeFieldListItem').clone()
	$defaultValControls.attr("id","")
	
	var $defaultValSelection = $defaultValControls.find("select")
	$defaultValSelection.empty()
	$defaultValSelection.append(defaultSelectOptionPromptHTML("Select a default value"))
	
	var defaultValuesTime = {
		"currentTime": {
			label: "Current date and time",
		}, 
		"clearValue": {
			label: "Clear value"
		}
	}
	
	for(var defaultValID in defaultValuesTime) {
	 	var selectDefaultValHTML = selectOptionHTML(defaultValID, defaultValuesTime[defaultValID].label)
	 	$defaultValSelection.append(selectDefaultValHTML)				
	}	
	
	if(defaultDefaultValInfo !== null) {
		$defaultValSelection.val(defaultDefaultValInfo.defaultValueID)
	}	
		
	initSelectControlChangeHandler($defaultValSelection,function(defaultValID) {
		var defaultValInfo = defaultValuesTime[defaultValID]
		console.log("Default value selection change: " + defaultValID)
		updateDefaultValues(panelParams)
	})
		
		
	var $defaultValListItem = createDefaultValueListItem(panelParams,fieldInfo.name)
		
	$defaultValListItem.data("defaultValConfigFunc",function() {
		var defaultValID = $defaultValSelection.val()
		if(defaultValID !== null && defaultValID.length > 0) {
			var conditions = []
			
			var defaultValConfig = { fieldID: fieldInfo.fieldID,
				defaultValueID: defaultValID }	
			return defaultValConfig
		} else {
			return null
		}
	})
		
	$defaultValListItem.append($defaultValControls)
		
	return $defaultValListItem
	
}


function numberDefaultValueItem(panelParams,fieldInfo,defaultDefaultValInfo) {
	
	var $defaultValControls = $('#defaultValueNumberFieldListItem').clone()
	$defaultValControls.attr("id","")
	
	var $defaultValNumberValInput = $defaultValControls.find(".defaultValueInput")
	
	
	var defaultValuesNumber = {
		"specificVal": {
			label: "Specific value",
			hasVal: true
		}, 
		"clearValue": {
			label: "Clear value",
			hasVal:false
		}
	}
	var $defaultValSelection = $defaultValControls.find("select")
	$defaultValSelection.empty()
	$defaultValSelection.append(defaultSelectOptionPromptHTML("Select a default value"))
	for(var defaultValID in defaultValuesNumber) {
	 	var selectDefaultValHTML = selectOptionHTML(defaultValID, defaultValuesNumber[defaultValID].label)
	 	$defaultValSelection.append(selectDefaultValHTML)				
	}
	
	function toggleNumberInputForDefaultVal(defaultValID) {
		var defaultValInfo = defaultValuesNumber[defaultValID]
		if (defaultValInfo.hasVal) {
			$defaultValNumberValInput.show()
		} else {
			$defaultValNumberValInput.hide()
		}
	}	
	
	initSelectControlChangeHandler($defaultValSelection,function(defaultValID) {
		toggleNumberInputForDefaultVal(defaultValID)
		console.log("Default value selection change: " + defaultValID)
		updateDefaultValues(panelParams)
	})
	
	
	// If the control is being configured with an existing/default value,
	// then initialize the controls accordingly.
	var $defaultValInput = $defaultValControls.find(".defaultValueInput")
	if(defaultDefaultValInfo !== null) {
		$defaultValSelection.val(defaultDefaultValInfo.defaultValueID)
		$defaultValInput.val(defaultDefaultValInfo.numberVal)
		toggleNumberInputForDefaultVal(defaultDefaultValInfo.defaultValueID)
	} else {
		$defaultValInput.hide()
	}
	
		
	$defaultValInput.blur(function() {
		var defaultVal = Number($defaultValInput.val())
		if(!isNaN(defaultVal)) {
			updateDefaultValues(panelParams)	
		}
	})
				
	var $defaultValListItem = createDefaultValueListItem(panelParams,fieldInfo.name)
		
	$defaultValListItem.data("defaultValConfigFunc",function() {
		
		
		var defaultValID = $defaultValSelection.val()
		if(defaultValID !== null && defaultValID.length > 0) {
			var conditions = []
			
			var defaultValConfig = { fieldID: fieldInfo.fieldID,
				defaultValueID: defaultValID }	
				
			var defaultValInfo = defaultValuesNumber[defaultValID]
				if (defaultValInfo.hasVal) {
					var defaultVal = Number($defaultValInput.val())
					if(!isNaN(defaultVal)) {
						defaultValConfig.numberVal = defaultVal
						return defaultValConfig
					} else {
						return null
					}
				} else {
					return defaultValConfig
				}
		} else {
			return null
		}
	})
		
	$defaultValListItem.append($defaultValControls)
		
	return $defaultValListItem
	
}

function textDefaultValueItem(panelParams,fieldInfo,defaultDefaultValInfo) {
	
	var $defaultValControls = $('#defaultValueTextFieldListItem').clone()
	$defaultValControls.attr("id","")
	
	var $defaultValTextValInput = $defaultValControls.find(".defaultValueInput")
	
	
	var defaultValuesText = {
		"specificVal": {
			label: "Specific value",
			hasVal: true
		}, 
		"clearValue": {
			label: "Clear value",
			hasVal:false
		}
	}
	var $defaultValSelection = $defaultValControls.find("select")
	$defaultValSelection.empty()
	$defaultValSelection.append(defaultSelectOptionPromptHTML("Select a default value"))
	for(var defaultValID in defaultValuesText) {
	 	var selectDefaultValHTML = selectOptionHTML(defaultValID, defaultValuesText[defaultValID].label)
	 	$defaultValSelection.append(selectDefaultValHTML)				
	}
	
	function toggleTextInputForDefaultVal(defaultValID) {
		var defaultValInfo = defaultValuesText[defaultValID]
		if (defaultValInfo.hasVal) {
			$defaultValTextValInput.show()
		} else {
			$defaultValTextValInput.hide()
		}
	}	
	
	initSelectControlChangeHandler($defaultValSelection,function(defaultValID) {
		toggleTextInputForDefaultVal(defaultValID)
		console.log("Default value selection change: " + defaultValID)
		updateDefaultValues(panelParams)
	})
	
	
	// If the control is being configured with an existing/default value,
	// then initialize the controls accordingly.
	var $defaultValInput = $defaultValControls.find(".defaultValueInput")
	if(defaultDefaultValInfo !== null) {
		$defaultValSelection.val(defaultDefaultValInfo.defaultValueID)
		$defaultValInput.val(defaultDefaultValInfo.textVal)
		toggleTextInputForDefaultVal(defaultDefaultValInfo.defaultValueID)
	} else {
		$defaultValInput.hide()
	}
	
		
	$defaultValInput.blur(function() {
		var defaultVal = $defaultValInput.val()
		if (defaultVal !== null) {
			updateDefaultValues(panelParams)
		}
	})
				
	var $defaultValListItem = createDefaultValueListItem(panelParams,fieldInfo.name)
		
	$defaultValListItem.data("defaultValConfigFunc",function() {
		
		
		var defaultValID = $defaultValSelection.val()
		if(defaultValID !== null && defaultValID.length > 0) {
			var conditions = []
			
			var defaultValConfig = { fieldID: fieldInfo.fieldID,
				defaultValueID: defaultValID }	
				
			var defaultValInfo = defaultValuesText[defaultValID]
				if (defaultValInfo.hasVal) {
					var defaultVal = $defaultValInput.val()
					if(defaultVal !== null) {
						defaultValConfig.textVal = defaultVal
						return defaultValConfig
					} else {
						return null
					}
				} else {
					return defaultValConfig
				}
		} else {
			return null
		}
	})
		
	$defaultValListItem.append($defaultValControls)
		
	return $defaultValListItem
	
}



function createDefaultValuePanelListItem(panelParams, fieldInfo,defaultValueInfo) {
	
	switch (fieldInfo.type) {
	case fieldTypeNumber:
		return numberDefaultValueItem(panelParams, fieldInfo, defaultValueInfo)
	case fieldTypeText:
		return textDefaultValueItem(panelParams, fieldInfo, defaultValueInfo)
	case fieldTypeTime: 
		return timeDefaultValueItem(panelParams, fieldInfo, defaultValueInfo)
	case fieldTypeBool: 
		return boolDefaultValueItem(panelParams, fieldInfo, defaultValueInfo)
	default:
		console.log("createFilterRulePanelListItem: Unsupported field type:  " + fieldInfo.type)
		return $("")
	}
	
}



function initDefaultDefaultValuePanelItems(panelParams) {
	
	loadFieldInfo(panelParams.databaseID,[fieldTypeAll],function(fieldsByID) {
		
		var defaultValListSelector = createPrefixedSelector(panelParams.elemPrefix,
						'DefaultValuesList')
		var $defaultValList = $(defaultValListSelector)		
		$defaultValList.empty()
		
		for(var defaultValIndex = 0; 
				defaultValIndex < panelParams.defaultDefaultValues.length; defaultValIndex++) {
					
			var currDefaultVal = panelParams.defaultDefaultValues[defaultValIndex]
			
			var fieldInfo = fieldsByID[currDefaultVal.fieldID]
					
			$defaultValList.append(createDefaultValuePanelListItem(panelParams,fieldInfo,currDefaultVal))
				
		}
	})
	
}
