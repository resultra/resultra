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
	
	return filterRules
}


function updateFilterRules(panelParams) {
	
	var filterRules = getRecordFilterRuleListRules(panelParams.elemPrefix)
	
	panelParams.updateFilterRules(filterRules)
}

function createFilterListRuleListItem(panelParams,fieldName) {
	
	var $filterRuleListItem = $('#recordFilterRuleListItem').clone()
	$filterRuleListItem.attr("id","")
	
	var $fieldLabel = $filterRuleListItem.find(".filterRuleListItemLabel")
	$fieldLabel.text(fieldName)
	
	var $deleteButton = $filterRuleListItem.find(".filterRuleListItemDeleteRuleButton")
	initButtonControlClickHandler($deleteButton,function() {
		$filterRuleListItem.remove()
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
	
	var $dateFilterModeSelection = $ruleControls.find(".filterDateRuleModeSelection")
	$dateFilterModeSelection.empty()
	$dateFilterModeSelection.append(defaultSelectOptionPromptHTML("Filter for"))
	
	var dateFilterModes = {
		"anyDate": {
			label: "Any date",
			modeSelected: function() {
				$startEndDateControls.hide()
				updateFilterRules(panelParams)
			},
			modeConfig: function() {
				var condition = { conditionID: "anyDate" }
				var ruleConfig = { fieldID: fieldInfo.fieldID,
					ruleID: "anyDate",
					conditions: [condition]}
				return ruleConfig				
			},
			initDefaultVals: function(ruleInfo) {
				$startEndDateControls.hide()			
			}
		},
		"customDateRange": {
			label: "Custom date range",
			modeSelected: function() {
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
					ruleID: "customDateRange", 
					conditions: conditions }
				return ruleConfig
			}, 
			initDefaultVals: function(ruleInfo) {
				var ruleConditions = mapRuleConditionsByOperatorID(ruleInfo)
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
	
	// Initialize the start and end date controls
	var $startDatePicker = $ruleControls.find(".filterDateRangeStartInput")
	$startDatePicker.datetimepicker()
	
	var $endDatePicker = $ruleControls.find(".filterDateRangeEndInput")
	$endDatePicker.datetimepicker({
            useCurrent: false //Important! See issue #1075
        });
		
		
		
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


function createFilterRulePanelListItem(panelParams, fieldInfo,defaultRuleInfo) {
	
	switch (fieldInfo.type) {
	case fieldTypeNumber:
		return numberFilterPanelRuleItem(panelParams, fieldInfo,defaultRuleInfo)
	case fieldTypeTime: 
		return dateFilterPanelRuleItem(panelParams, fieldInfo,defaultRuleInfo)
	default:
		console.log("createFilterRulePanelListItem: Unsupported field type:  " + fieldInfo.type)
		return $("")
	}
	
}

function initDefaultFilterRules(panelParams) {
	
	loadFieldInfo(panelParams.tableID,[fieldTypeAll],function(fieldsByID) {
		
		var filterRuleListSelector = createPrefixedSelector(panelParams.elemPrefix,
						'RecordFilterFilterRuleList')
		var $filterRuleList = $(filterRuleListSelector)		
		$filterRuleList.empty()
		
		for(var defaultRuleIndex = 0; 
				defaultRuleIndex < panelParams.defaultFilterRules.length; defaultRuleIndex++) {
					
			var currRuleInfo = panelParams.defaultFilterRules[defaultRuleIndex]
			
			var fieldInfo = fieldsByID[currRuleInfo.fieldID]
					
			$filterRuleList.append(createFilterRulePanelListItem(panelParams,fieldInfo,currRuleInfo))
				
		}
		
		panelParams.initDone()
		
	})
	
	
	
}
