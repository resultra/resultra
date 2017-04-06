
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



function createDefaultValuePanelListItem(panelParams, fieldInfo,defaultValueInfo) {
	
	switch (fieldInfo.type) {
	case fieldTypeNumber:
		console.log("createDefaultValuePanelListItem: Unsupported field type:  " + fieldInfo.type)
		return $("")
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
