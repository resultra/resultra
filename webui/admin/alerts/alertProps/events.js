
function updateAlertConditions(panelParams) {
	console.log("Updating alert conditions ...")
	
	var $alertConditionList = $("#alertConditionList")
	var conditions = []
	$alertConditionList.find(".alertConditionListItem").each(function() {
		var condFunc = $(this).data("alertCondDefFunc")
		var condProp = condFunc()
		if (condProp != null) {
			conditions.push(condProp)
		}
	})
	console.log("alert conditions: " + JSON.stringify(conditions))
	
	var setConditionParams = { 
		alertID: panelParams.alertID,
		conditions: conditions
	}
	jsonAPIRequest("alert/setConditions",setConditionParams,function(alerts) {		
	})	
	
	
//	config.setConditionalFormats(conditions)
}

function createAlertConditionListItem(panelParams,fieldName) {
	
	var $listItem = $('#alertConditionListItem').clone()
	$listItem.attr("id","")
	
	var $fieldLabel = $listItem.find(".alertConditionItemLabel")
	$fieldLabel.text(fieldName)
	
	var $deleteButton = $listItem.find(".alertConditionDeleteButton")
	initButtonControlClickHandler($deleteButton,function() {
		$listItem.remove()
		updateAlertConditions(panelParams)
	})

	return $listItem
}

function dateAlertConditionListItem(propsParams,fieldInfo,defaultConditionInfo) {
	
	var $listItem = createAlertConditionListItem(propsParams,fieldInfo.name)
	
	var $alertProps = $("#alertDateFieldConditionProps").clone()
	$alertProps.attr("id","")
	var $dateParamInput = $alertProps.find(".condDateParamInput")
		$dateParamInput.datetimepicker()
	
	var $numParamInput = $alertProps.find(".condNumParamInput")
	
	
	var alertCondDefs = {
		"past": {
			label: "Date has passed",
			hasDateParam: false,
			hasNumberParam: false
		},
		"changed": {
			label: "Date changed",
			hasDateParam: false,
			hasNumberParam: false
		},
		"cleared": {
			label: "Date cleared",
			hasDateParam: false,
			hasNumberParam: false
		},
	}
	
	
	
	function initConditionDef(condDefID) {
		
		var condDef
		if (condDefID !== null) {
			condDef = alertCondDefs[condDefID]
		} else {
			condDef = {
				hasDateParam:false,
				hasNumberParam:false
			} 			
		}
		
		if (condDef.hasDateParam) {
			$dateParamInput.css("display","")
		} else {
			$dateParamInput.css("display","none")
		}
		
		if (condDef.hasNumberParam) {
			$numParamInput.show()
		} else {
			$numParamInput.hide()
		}
		
	}
	
	var $modeSelection = $alertProps.find(".alertConditionDateModeSelection")
	$modeSelection.empty()
	$modeSelection.append(defaultSelectOptionPromptHTML("Select a condition"))
	for(var condID in alertCondDefs) {
	 	var selectCondHTML = selectOptionHTML(condID, alertCondDefs[condID].label)
	 	$modeSelection.append(selectCondHTML)				
	}
	
	if (defaultConditionInfo !== null) {
		initConditionDef(defaultConditionInfo.conditionID)
		var condDef = alertCondDefs[defaultConditionInfo.conditionID]
		
		$modeSelection.val(defaultConditionInfo.conditionID)
		
		if (condDef.hasDateParam) {
			var dateMoment = moment(defaultConditionInfo.dateParam)
			$dateParamInput.data("DateTimePicker").date(dateMoment)			
		} else {
			$dateParamInput.data("DateTimePicker").date(null)
		}
		if (condDef.hasNumberParam) {
			$numParamInput.val(defaultConditionInfo.numberParam)
		} else {
			$numParamInput.val(null)
		}
	} else {
		initConditionDef(null)
		$dateParamInput.data("DateTimePicker").date(null)
		$numParamInput.val(null)
	}
	
	initSelectControlChangeHandler($modeSelection,function(conditionID) {
		initConditionDef(conditionID)
		updateAlertConditions(propsParams)
	})
	
    $dateParamInput.on("dp.change", function (e) {
		updateAlertConditions(propsParams)
    });
	$numParamInput.blur(function() {
		updateConditionProperties()			
	})
	
	$listItem.data("alertCondDefFunc",function() {
		var condID = $modeSelection.val()
		
		if(condID === null || condID.length <= 0) {
			return null
		}
		var condDef = { fieldID: fieldInfo.fieldID,
			conditionID: condID }	
		
		var condInfo = alertCondDefs[condID]
		
		if (condInfo.hasNumberParam) {
			var numberVal = convertStringToNumber($numParamInput.val())
			if(numberVal === null) {
				return null	
			}
			condDef.numberParam = numberVal
		}
			
		if (condInfo.hasStartDate) {
			var dateVal = $dateParamInput.data("DateTimePicker").date()
			if (dateVal == null) {
				return null
			}
			condDef.dateParam = dateVal.utc()
		}
		
		return condDef
	})
	
	
	
	$listItem.append($alertProps)
	
	return $listItem
}

function boolAlertConditionListItem(propsParams,fieldInfo,defaultConditionInfo) {
	
	var $listItem = createAlertConditionListItem(propsParams,fieldInfo.name)
	
	var $alertProps = $("#alertBoolFieldConditionProps").clone()
	$alertProps.attr("id","")
		
	
	var alertCondDefs = {
		"true": {
			label: "Value set to true"
		},
		"false": {
			label: "Value set to false"
		},
		"cleared": {
			label: "Value cleared"
		},
		"changed": {
			label: "Value changed"
		}
	}
	
	function initConditionDef(condDefID) {
		
		var condDef
		if (condDefID !== null) {
			condDef = alertCondDefs[condDefID]
		} else {
			condDef = {} 			
		}
				
	}
	
	var $modeSelection = $alertProps.find(".alertConditionBoolModeSelection")
	$modeSelection.empty()
	$modeSelection.append(defaultSelectOptionPromptHTML("Select a condition"))
	for(var condID in alertCondDefs) {
	 	var selectCondHTML = selectOptionHTML(condID, alertCondDefs[condID].label)
	 	$modeSelection.append(selectCondHTML)				
	}
	
	if (defaultConditionInfo !== null) {
		initConditionDef(defaultConditionInfo.conditionID)
		var condDef = alertCondDefs[defaultConditionInfo.conditionID]
		
		$modeSelection.val(defaultConditionInfo.conditionID)
		
	} else {
		initConditionDef(null)
	}
	
	initSelectControlChangeHandler($modeSelection,function(conditionID) {
		initConditionDef(conditionID)
		updateAlertConditions(propsParams)
	})
		
	$listItem.data("alertCondDefFunc",function() {
		var condID = $modeSelection.val()
		
		if(condID === null || condID.length <= 0) {
			return null
		}
		var condDef = { fieldID: fieldInfo.fieldID,
			conditionID: condID }	
		
		var condInfo = alertCondDefs[condID]
				
		return condDef
	})

	$listItem.append($alertProps)
	
	return $listItem
}

function numberAlertConditionListItem(propsParams,fieldInfo,defaultConditionInfo) {
	
	var $listItem = createAlertConditionListItem(propsParams,fieldInfo.name)
	
	var $alertProps = $("#alertNumberFieldConditionProps").clone()
	$alertProps.attr("id","")
	
	var $numParamInput = $alertProps.find(".condNumParamInput")
	
	
	var alertCondDefs = {
		"increased": {
			label: "Value increased",
			hasNumberParam: false
		},
		"decreased": {
			label: "Value decreased",
			hasNumberParam: false
		},
		"changed": {
			label: "Value changed",
			hasNumberParam: false
		},
		"cleared": {
			label: "Value cleared",
			hasNumberParam: false
		},
	}
	
	
	
	function initConditionDef(condDefID) {
		
		var condDef
		if (condDefID !== null) {
			condDef = alertCondDefs[condDefID]
		} else {
			condDef = {
				hasNumberParam:false
			} 			
		}
				
		if (condDef.hasNumberParam) {
			$numParamInput.show()
		} else {
			$numParamInput.hide()
		}
		
	}
	
	var $modeSelection = $alertProps.find(".alertConditionModeSelection")
	$modeSelection.empty()
	$modeSelection.append(defaultSelectOptionPromptHTML("Select a condition"))
	for(var condID in alertCondDefs) {
	 	var selectCondHTML = selectOptionHTML(condID, alertCondDefs[condID].label)
	 	$modeSelection.append(selectCondHTML)				
	}
	
	if (defaultConditionInfo !== null) {
		initConditionDef(defaultConditionInfo.conditionID)
		var condDef = alertCondDefs[defaultConditionInfo.conditionID]
		
		$modeSelection.val(defaultConditionInfo.conditionID)
		
		if (condDef.hasNumberParam) {
			$numParamInput.val(defaultConditionInfo.numberParam)
		} else {
			$numParamInput.val(null)
		}
	} else {
		initConditionDef(null)
		$numParamInput.val(null)
	}
	
	initSelectControlChangeHandler($modeSelection,function(conditionID) {
		initConditionDef(conditionID)
		updateAlertConditions(propsParams)
	})
		
	$listItem.data("alertCondDefFunc",function() {
		var condID = $modeSelection.val()
		
		if(condID === null || condID.length <= 0) {
			return null
		}
		var condDef = { fieldID: fieldInfo.fieldID,
			conditionID: condID }	
		
		var condInfo = alertCondDefs[condID]
		
		if (condInfo.hasNumberParam) {
			var numberVal = convertStringToNumber($numParamInput.val())
			if(numberVal === null) {
				return null	
			}
			condDef.numberParam = numberVal
		}
					
		return condDef
	})
	
	$listItem.append($alertProps)
	
	return $listItem
}


function createAlertPropsConditionItem(propsParams,fieldInfo,defaultConditionInfo) {

	switch (fieldInfo.type) {
	case fieldTypeTime: 
		return dateAlertConditionListItem(propsParams,fieldInfo,defaultConditionInfo)
	case fieldTypeBool: 
		return boolAlertConditionListItem(propsParams,fieldInfo,defaultConditionInfo)
	case fieldTypeNumber:
		return numberAlertConditionListItem(propsParams,fieldInfo,defaultConditionInfo)
	case fieldTypeText:
	default:
		console.log("createFilterRulePanelListItem: Unsupported field type:  " + fieldInfo.type)
		return $("<div>TBD</div>")
	}
	
}

function initAlertConditionProps(params) {
	
	var conditionFieldTypes = [fieldTypeTime,fieldTypeBool,fieldTypeNumber]
	
	function loadDefaultConditions() {
		loadFieldInfo(params.databaseID,conditionFieldTypes,function(fieldsByID) {
			
			var getAlertParams = { 
				alertID: params.alertID
			}
			jsonAPIRequest("alert/get",getAlertParams,function(alertInfo) {		
				var $alertConditionList = $("#alertConditionList")
				var conditions = alertInfo.properties.conditions
				
				for(var condIndex = 0; 
						condIndex < conditions.length; condIndex++) {
					
					var currCond = conditions[condIndex]
			
					var fieldInfo = fieldsByID[currCond.fieldID]
					
					$alertConditionList.append(createAlertPropsConditionItem(params,fieldInfo,currCond))
				
				}
			})			
		})
	
	}
	loadDefaultConditions()
	
			
	var fieldSelectionDropdownParams = {
		elemPrefix: "alertCondition_",
		databaseID: params.databaseID,
		fieldTypes: conditionFieldTypes,
		fieldSelectionCallback: function(fieldInfo) {
			
			var $alertConditionList = $("#alertConditionList")
			
			// Use null to signify no default condition information. This is true when
			// creating new rules, but will not be when re-loading the rules.
			var defaultConditionInfo = null
			$alertConditionList.append(createAlertPropsConditionItem(params,fieldInfo,defaultConditionInfo))
		}
	}
	initFieldSelectionDropdown(fieldSelectionDropdownParams)
	
	
}