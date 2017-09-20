function getRecordFilterRuleListRules(elemPrefix) {
	var filterRules = []
	
	var filterRuleListSelector = createPrefixedSelector(elemPrefix,
					'RecordFilterFilterRuleList')
	
	$(filterRuleListSelector + " .recordFilterPanelRuleListItem").each(function() {
	
		var filterRuleConfigFunc = $(this).data("filterRuleConfigFunc")
		var ruleConfig = filterRuleConfigFunc()
		
		if(ruleConfig != null) {
			filterRules.push(ruleConfig)
		}
	})

	console.log("filterRules rules: " + JSON.stringify(filterRules))
	
	var $matchLogic = $(createPrefixedSelector(elemPrefix,
					'RecordFilterMatchLogicSelection'))
	
	var filterRuleSet = {
		matchLogic: $matchLogic.val(),
		filterRules: filterRules
	}
	
	return filterRuleSet
}


function updateFilterRules(panelParams) {
	
	var filterRuleSet = getRecordFilterRuleListRules(panelParams.elemPrefix)
		
	panelParams.updateFilterRules(filterRuleSet)
}

function createFilterListRuleListItem(panelParams,fieldName) {
	
	var $filterRuleListItem = $('#recordFilterRuleListItem').clone()
	$filterRuleListItem.attr("id","")
	
	var $fieldLabel = $filterRuleListItem.find(".filterRuleListItemLabel")
	$fieldLabel.text(fieldName)
	
	var $deleteButton = $filterRuleListItem.find(".filterRuleListItemDeleteRuleButton")
	initButtonControlClickHandler($deleteButton,function() {
		$filterRuleListItem.remove()
		updateMatchLogicSelectionVisibility(panelParams)
		updateFilterRules(panelParams)
	})

	return $filterRuleListItem
}

function mapRuleConditionsByOperatorID(ruleInfo) {
	var ruleOpererators = {}
	for(var paramIndex = 0; paramIndex < ruleInfo.conditions.length; paramIndex++) {
		var currCondition = ruleInfo.conditions[paramIndex]
		ruleOpererators[currCondition.operatorID] = currCondition
	}
	return ruleOpererators
}

function dateFilterPanelRuleItem(panelParams,fieldInfo,defaultRuleInfo) {
	
	var $ruleControls = $('#recordFilterDateFieldRuleListItem').clone()
	$ruleControls.attr("id","")
	
	var $startEndDateControls = $ruleControls.find(".filterCustomStartEndDateControls")
	$startEndDateControls.hide()
	
	var $startDateInput = $ruleControls.find(".filterDateRangeStartInput")
	var $endDateInput = $ruleControls.find(".filterDateRangeEndInput")
	
	var $dateFilterModeSelection = $ruleControls.find(".filterDateRuleModeSelection")
	$dateFilterModeSelection.empty()
	$dateFilterModeSelection.append(defaultSelectOptionPromptHTML("Filter for"))
	
	var dateFilterModes = {
		"any": {
			label: "Any date (no filtering)",
			modeSelected: function() {
				$startEndDateControls.hide()
				updateFilterRules(panelParams)
			},
			modeConfig: function() {
				var condition = { conditionID: "anyDate" }
				var ruleConfig = { fieldID: fieldInfo.fieldID,
					ruleID: "any",
					conditions: [condition]}
				return ruleConfig				
			},
			initDefaultVals: function(ruleInfo) {
				$startEndDateControls.hide()			
			}
		},
		"notBlank": {
			label: "Date is set (not blank)",
			modeSelected: function() {
				$startEndDateControls.hide()
				updateFilterRules(panelParams)
			},
			modeConfig: function() {
				var condition = { conditionID: "notBlank" }
				var ruleConfig = { fieldID: fieldInfo.fieldID,
					ruleID: "notBlank",
					conditions: [condition]}
				return ruleConfig				
			},
			initDefaultVals: function(ruleInfo) {
				$startEndDateControls.hide()			
			}
		},
		"isBlank": {
			label: "Date is not set (blank)",
			modeSelected: function() {
				$startEndDateControls.hide()
				updateFilterRules(panelParams)
			},
			modeConfig: function() {
				var condition = { conditionID: "isBlank" }
				var ruleConfig = { fieldID: fieldInfo.fieldID,
					ruleID: "isBlank",
					conditions: [condition]}
				return ruleConfig				
			},
			initDefaultVals: function(ruleInfo) {
				$startEndDateControls.hide()			
			}
		},
		"after": {
			label: "After",
			modeSelected: function() {
				$startDateInput.show()
				$endDateInput.hide()
				$startEndDateControls.show()
				updateFilterRules(panelParams)
			},
			modeConfig: function() {
				var startDate = $startDatePicker.data("DateTimePicker").date()
				if (startDate === null) { return null }
				var startDateUTC = startDate.utc()
				var conditions = [
					{ operatorID: "minDate", dateParam: startDateUTC }
				]
				var ruleConfig = { fieldID: fieldInfo.fieldID, 
					ruleID: "after", 
					conditions: conditions }
				return ruleConfig
			}, 
			initDefaultVals: function(ruleInfo) {
				var ruleConditions = mapRuleConditionsByOperatorID(ruleInfo)
				$startDateInput.show()
				$endDateInput.hide()
				$startEndDateControls.show()
				var startDate = moment(ruleConditions["minDate"].dateParam)
				$startDatePicker.data("DateTimePicker").date(startDate)
			}
		},
		"before": {
			label: "Before",
			modeSelected: function() {
				$startDateInput.hide()
				$endDateInput.show()
				$startEndDateControls.show()
				updateFilterRules(panelParams)
			},
			modeConfig: function() {
				var endDate = $endDatePicker.data("DateTimePicker").date()
				if (endDate === null) { return null }
				var endDateUTC = endDate.utc()
				var conditions = [
					{ operatorID: "maxDate", dateParam: endDateUTC }
				]
				var ruleConfig = { fieldID: fieldInfo.fieldID, 
					ruleID: "before", 
					conditions: conditions }
				return ruleConfig
			}, 
			initDefaultVals: function(ruleInfo) {
				var ruleConditions = mapRuleConditionsByOperatorID(ruleInfo)
				$startDateInput.hide()
				$endDateInput.show()
				$startEndDateControls.show()
				var endDate = moment(ruleConditions["maxDate"].dateParam)
				$endDatePicker.data("DateTimePicker").date(endDate)
			}
		},
		"dateRange": {
			label: "Custom date range",
			modeSelected: function() {
				$startDateInput.show()
				$endDateInput.show()
				$startEndDateControls.show()
				updateFilterRules(panelParams)
			},
			modeConfig: function() {
				var startDate = $startDatePicker.data("DateTimePicker").date()
				if (startDate === null) { return null }
				var startDateUTC = startDate.utc()
				var endDate = $endDatePicker.data("DateTimePicker").date()
				if (endDate === null) { return null }
				var endDateUTC = endDate.utc()
				var conditions = [
					{ operatorID: "minDate", dateParam: startDateUTC },
					{ operatorID: "maxDate", dateParam: endDateUTC }
				]
				var ruleConfig = { fieldID: fieldInfo.fieldID, 
					ruleID: "dateRange", 
					conditions: conditions }
				return ruleConfig
			}, 
			initDefaultVals: function(ruleInfo) {
				var ruleConditions = mapRuleConditionsByOperatorID(ruleInfo)
				$startDateInput.show()
				$endDateInput.show()
				$startEndDateControls.show()
				var startDate = moment(ruleConditions["minDate"].dateParam)
				$startDatePicker.data("DateTimePicker").date(startDate)
				var endDate = moment(ruleConditions["maxDate"].dateParam)
				$endDatePicker.data("DateTimePicker").date(endDate)
			}
		},
	}
	for(var modeID in dateFilterModes) {
	 	var selectModeHTML = selectOptionHTML(modeID, dateFilterModes[modeID].label)
	 	$dateFilterModeSelection.append(selectModeHTML)				
	}
	initSelectControlChangeHandler($dateFilterModeSelection,function(modeID) {
		var modeInfo = dateFilterModes[modeID]
		console.log("Date rule mode selection change: " + modeID)
		modeInfo.modeSelected()
	})
	
	var dateTimePickerOptions = {
		format: getDateComponentDateTimeMomentFormat("date"),
		useCurrent: false //Important! See issue #1075
	}
	
	// Initialize the start and end date controls
	var $startDatePicker = $ruleControls.find(".filterDateRangeStartInput")
	$startDatePicker.datetimepicker(dateTimePickerOptions)
	
	var $endDatePicker = $ruleControls.find(".filterDateRangeEndInput")
	$endDatePicker.datetimepicker(dateTimePickerOptions);
		
		
		
	// Initialization of default values needs to happen *before* the 
	// event handlers are setup for the date pickers. This ensures
	// the event handlers aren't triggered while initializing the defaults.
	if(defaultRuleInfo !== null) {
		var ruleInfo = dateFilterModes[defaultRuleInfo.ruleID]
		$dateFilterModeSelection.val(defaultRuleInfo.ruleID)
		ruleInfo.initDefaultVals(defaultRuleInfo)	
	}
		
	
	// Link the start and end date controls based to ensure
    // the range is preserved.
    $startDatePicker.on("dp.change", function (e) {
		console.log("Custom start date changed: " + e.date)
        $endDatePicker.data("DateTimePicker").minDate(e.date);
		updateFilterRules(panelParams)
    });
    $endDatePicker.on("dp.change", function (e) {
		console.log("Custom end date changed: " + e.date)
        $startDatePicker.data("DateTimePicker").maxDate(e.date);
		updateFilterRules(panelParams)
    });
	
	var $filterListItem = createFilterListRuleListItem(panelParams,fieldInfo.name)
	$filterListItem.data("filterRuleConfigFunc",function() {
		console.log("Date filter rule - getting config")
		var modeID = $dateFilterModeSelection.val()
		if(modeID !== null && modeID.length > 0) {
			var modeInfo = dateFilterModes[modeID]
			return modeInfo.modeConfig()
		} else {
			return null
		}
		
	})
	
	$filterListItem.append($ruleControls)
		
	return $filterListItem
	
	
}


function numberFilterPanelRuleItem(panelParams,fieldInfo,defaultRuleInfo) {
	
	var $ruleControls = $('#recordFilterNumberFieldRuleListItem').clone()
	$ruleControls.attr("id","")
	
	var $ruleSelection = $ruleControls.find("select")
	$ruleSelection.empty()
	$ruleSelection.append(defaultSelectOptionPromptHTML("Filter for"))
	
	for(var ruleID in filterRulesNumber) {
	 	var selectRuleHTML = selectOptionHTML(ruleID, filterRulesNumber[ruleID].label)
	 	$ruleSelection.append(selectRuleHTML)				
	}
		
	var $paramInput = $ruleControls.find("input")
	$paramInput.hide()
	
	
	if(defaultRuleInfo !== null) {
		var ruleInfo = filterRulesNumber[defaultRuleInfo.ruleID]
		$ruleSelection.val(defaultRuleInfo.ruleID)
		if(ruleInfo.hasParam) {
			var ruleConditions = mapRuleConditionsByOperatorID(defaultRuleInfo)
			var numParam = ruleConditions[defaultRuleInfo.ruleID].numberParam
			$paramInput.val(numParam)
			$paramInput.show()
		}
	}
	
	
	$paramInput.blur(function() {
		var numberVal = Number($paramInput.val())
		if(!isNaN(numberVal)) {
			updateFilterRules(panelParams)	
		}
	})
	
		
	initSelectControlChangeHandler($ruleSelection,function(ruleID) {
		var ruleInfo = filterRulesNumber[ruleID]
		console.log("Rule selection change: " + ruleID)
		if(ruleInfo.hasParam) {
			$paramInput.val("")
			$paramInput.show()
		} else {
			$paramInput.hide()
			updateFilterRules(panelParams)
		}
	})
		
		
	var $filterListItem = createFilterListRuleListItem(panelParams,fieldInfo.name)
		
	$filterListItem.data("filterRuleConfigFunc",function() {
		var ruleID = $ruleSelection.val()
		var paramVal = $paramInput.val()
		if(ruleID !== null && ruleID.length > 0) {
			var ruleInfo = filterRulesNumber[ruleID]
			var conditions = []
			if(ruleInfo.hasParam) {
				conditions.push({ operatorID: ruleID, numberParam: Number(paramVal) })				
			} else {
				conditions.push({ operatorID: ruleID })				
			}
			
			var ruleConfig = { fieldID: fieldInfo.fieldID,
				ruleID: ruleID,
				conditions: conditions }
			return ruleConfig
		} else {
			return null
		}
	})
		
	$filterListItem.append($ruleControls)
		
	return $filterListItem
	
}

function boolFilterPanelRuleItem(panelParams,fieldInfo,defaultRuleInfo) {
	
	var $ruleControls = $('#recordFilterBoolFieldRuleListItem').clone()
	$ruleControls.attr("id","")
	
	var $ruleSelection = $ruleControls.find("select")
	$ruleSelection.empty()
	$ruleSelection.append(defaultSelectOptionPromptHTML("Filter for"))
	
	var filterRulesBool = {
		"any": {
			label: "Any value (no filtering)",
		},
		"isBlank": {
			label: "Value not set (blank)",
		},
		"notBlank": {
			label: "Value is set (not blank)",
		},
		"true": {
			label: "Value is true",
		},
		"notTrue": {
			label: "Value not true",
		},
		"false": {
			label: "Value is false",
		},
		"notFalse": {
			label: "Value is not False",
		}
	}
	
	for(var ruleID in filterRulesBool) {
	 	var selectRuleHTML = selectOptionHTML(ruleID, filterRulesBool[ruleID].label)
	 	$ruleSelection.append(selectRuleHTML)				
	}	
	
	if(defaultRuleInfo !== null) {
		var ruleInfo = filterRulesBool[defaultRuleInfo.ruleID]
		$ruleSelection.val(defaultRuleInfo.ruleID)
	}	
		
	initSelectControlChangeHandler($ruleSelection,function(ruleID) {
		var ruleInfo = filterRulesBool[ruleID]
		console.log("Rule selection change: " + ruleID)
		updateFilterRules(panelParams)
	})
		
		
	var $filterListItem = createFilterListRuleListItem(panelParams,fieldInfo.name)
		
	$filterListItem.data("filterRuleConfigFunc",function() {
		var ruleID = $ruleSelection.val()
		if(ruleID !== null && ruleID.length > 0) {
			var ruleInfo = filterRulesBool[ruleID]
			var conditions = []
			conditions.push({ operatorID: ruleID })				
			
			var ruleConfig = { fieldID: fieldInfo.fieldID,
				ruleID: ruleID,
				conditions: conditions }	
			return ruleConfig
		} else {
			return null
		}
	})
		
	$filterListItem.append($ruleControls)
		
	return $filterListItem
	
}


function textFilterPanelRuleItem(panelParams,fieldInfo,defaultRuleInfo) {
	
	var $ruleControls = $('#recordFilterTextFieldRuleListItem').clone()
	$ruleControls.attr("id","")
	
	var $ruleSelection = $ruleControls.find("select")
	$ruleSelection.empty()
	$ruleSelection.append(defaultSelectOptionPromptHTML("Filter for"))
	
	var $ruleParam = $ruleControls.find(".recordFilterRuleParam")
	
	var filterRulesText = {
		"isBlank": {
			label: "Value not set (blank)",
			hasParam: false,
		},
		"notBlank": {
			label: "Value is set (not blank)",
			hasParam: false,
		},
		"contains": {
			label: "Text contains",
			hasParam: true,
		}
	}
	
	for(var ruleID in filterRulesText) {
	 	var selectRuleHTML = selectOptionHTML(ruleID, filterRulesText[ruleID].label)
	 	$ruleSelection.append(selectRuleHTML)				
	}	
	
	if(defaultRuleInfo !== null) {
		var ruleInfo = filterRulesText[defaultRuleInfo.ruleID]
		$ruleSelection.val(defaultRuleInfo.ruleID)
		if (ruleInfo.hasParam) {
			$ruleParam.show()
		} else {
			$ruleParam.hide()
		}
	} else {
		// Parameter input initially hidden until a filtering rule is selected
		$ruleParam.hide()		
	}	
		
	initSelectControlChangeHandler($ruleSelection,function(ruleID) {
		var ruleInfo = filterRulesText[ruleID]
		console.log("Rule selection change: " + ruleID)
		if (ruleInfo.hasParam) {
			$ruleParam.show()
		} else {
			$ruleParam.hide()
		}
		updateFilterRules(panelParams)
	})
	
	
	$ruleParam.blur(function() {
		var textVal = $ruleParam.val()
		if(nonEmptyStringVal(textVal)) {
			updateFilterRules(panelParams)	
		}
	})
			
	var $filterListItem = createFilterListRuleListItem(panelParams,fieldInfo.name)
		
	$filterListItem.data("filterRuleConfigFunc",function() {
		var ruleID = $ruleSelection.val()
		if(ruleID !== null && ruleID.length > 0) {
			var ruleInfo = filterRulesText[ruleID]
			var conditions = []
			if(ruleInfo.hasParam) {
				var paramVal = $ruleParam.val()
				conditions.push({ operatorID: ruleID, textParam: paramVal })				
			} else {
				conditions.push({ operatorID: ruleID })				
			}
			
			var ruleConfig = { fieldID: fieldInfo.fieldID,
				ruleID: ruleID,
				conditions: conditions }	
			return ruleConfig
		} else {
			return null
		}
	})
		
	$filterListItem.append($ruleControls)
		
	return $filterListItem
	
}


function tagFilterPanelRuleItem(panelParams,fieldInfo,defaultRuleInfo) {
	
	var $ruleControls = $('#recordFilterTagFieldRuleListItem').clone()
	$ruleControls.attr("id","")
	
	var ruleID = "tags"
	
	var $tagSelection = $ruleControls.find(".recordFilterTagSelections")
	
	var $matchLogicSelection = $ruleControls.find(".recordFilterMatchLogicSelection")
	
				
	var $filterListItem = createFilterListRuleListItem(panelParams,fieldInfo.name)
	
	function populateTagControl(tags) {
		$tagSelection.empty()
		var tagAdded = {}
		for(var tagIndex in tags) {
			var currTag = tags[tagIndex]
			if(tagAdded[currTag] === undefined) { // don't add duplicates
				var newOption = new Option(currTag,currTag);
				tagAdded[currTag] = true
				$tagSelection.append(newOption)					
			}
		}
		
	}
	
	
	if(defaultRuleInfo !== null) {
		var ruleConditions = mapRuleConditionsByOperatorID(defaultRuleInfo)
		var defaultTagList = ruleConditions[ruleID].tagsParam
		populateTagControl(defaultTagList)
		
		$tagSelection.val(defaultTagList)
	} 
	
		
	$filterListItem.data("filterRuleConfigFunc",function() {
		var selectedTags = $tagSelection.val()
		if (selectedTags !== null && selectedTags.length > 0) {
			
			var matchLogicConditionID = "logic"
			
			var conditions = []
			conditions.push({operatorID:ruleID, tagsParam: selectedTags })
			conditions.push({operatorID:matchLogicConditionID, textParam: $matchLogicSelection.val()})
			var ruleConfig = { fieldID: fieldInfo.fieldID,
				ruleID: ruleID,
				conditions: conditions }	
			return ruleConfig
			
		} else {
			return null;			
		}
	})
	
	$tagSelection.select2({
		placeholder: "Enter tags", // TODO - Allow a property to configure the placeholder.
		width: 250,
		tags:true,
		minimumInputLength: 1,
		maximumInputLength:32,
		tokenSeparators: [',']
	});
	
	$tagSelection.on('change', function() {
		updateFilterRules(panelParams)
	});
	$matchLogicSelection.on('change', function() {
		updateFilterRules(panelParams)
	});
	
	$filterListItem.append($ruleControls)
		
	return $filterListItem
	
}




function createFilterRulePanelListItem(panelParams, fieldInfo,defaultRuleInfo) {
	
	switch (fieldInfo.type) {
	case fieldTypeNumber:
		return numberFilterPanelRuleItem(panelParams, fieldInfo, defaultRuleInfo)
	case fieldTypeTime: 
		return dateFilterPanelRuleItem(panelParams, fieldInfo, defaultRuleInfo)
	case fieldTypeBool: 
		return boolFilterPanelRuleItem(panelParams, fieldInfo, defaultRuleInfo)
	case fieldTypeText:
		return textFilterPanelRuleItem(panelParams,fieldInfo,defaultRuleInfo)
	case fieldTypeTag:
		return tagFilterPanelRuleItem(panelParams,fieldInfo,defaultRuleInfo)
	default:
		console.log("createFilterRulePanelListItem: Unsupported field type:  " + fieldInfo.type)
		return $("")
	}
	
}

function updateDefaultFilterRules(panelParams, updateDoneFunc) {
	loadFieldInfo(panelParams.databaseID,[fieldTypeAll],function(fieldsByID) {
		
		var filterRuleListSelector = createPrefixedSelector(panelParams.elemPrefix,
						'RecordFilterFilterRuleList')
		var $filterRuleList = $(filterRuleListSelector)		
		$filterRuleList.empty()
			
		var ruleList = panelParams.defaultFilterRules.filterRules
		var matchLogic = panelParams.defaultFilterRules.matchLogic
		
		for(var defaultRuleIndex = 0; 
				defaultRuleIndex < ruleList.length; defaultRuleIndex++) {
					
			var currRuleInfo = ruleList[defaultRuleIndex]
			
			var fieldInfo = fieldsByID[currRuleInfo.fieldID]
					
			$filterRuleList.append(createFilterRulePanelListItem(panelParams,fieldInfo,currRuleInfo))
				
		}
		
		var $matchLogic = $(createPrefixedSelector(panelParams.elemPrefix,
						'RecordFilterMatchLogicSelection'))
		$matchLogic.val(matchLogic)
			
		updateMatchLogicSelectionVisibility(panelParams)
		
		updateDoneFunc()
	})
	
}

function initDefaultFilterRules(panelParams) {
	updateDefaultFilterRules(panelParams,panelParams.initDone)	
}

// The match logic selection only needs to be visible if there are more than 1 filter rules visible. Otherwise, it
// it doesn't matter if the filter rules are AND'ed or OR'd together.
function updateMatchLogicSelectionVisibility(panelParams) {
	
	var $filterRuleList = $(createPrefixedSelector(panelParams.elemPrefix,'RecordFilterFilterRuleList'))
	var $matchLogicSelectionControls = $(createPrefixedSelector(panelParams.elemPrefix,'RecordFilterMatchLogicSelectionControls'))
	
	var numFilterRules = $filterRuleList.find(".recordFilterPanelRuleListItem").length
	
	if(numFilterRules >= 2) {
		$matchLogicSelectionControls.show()
	} else {
		$matchLogicSelectionControls.hide()
	}
	
}

function initMatchLogicSelection(panelParams) {
	
	var $matchLogicSelection = $(createPrefixedSelector(panelParams.elemPrefix,'RecordFilterMatchLogicSelection'))
	initSelectControlChangeHandler($matchLogicSelection,function(selectedLogic) {
		updateFilterRules(panelParams)
	})			
	
}

