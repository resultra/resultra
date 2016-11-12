
function updateFilterRules(panelParams) {
	var filterRules = []
	
	var filterRuleListSelector = createPrefixedSelector(panelParams.elemPrefix,
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

function dateFilterPanelRuleItem(panelParams,fieldInfo) {
	
	var $ruleControls = $('#recordFilterDateFieldRuleListItem').clone()
	$ruleControls.attr("id","")
	
	var $startEndDateControls = $ruleControls.find(".filterCustomStartEndDateControls")
	$startEndDateControls.hide()
	
	var $dateFilterModeSelection = $ruleControls.find(".filterDateRuleModeSelection")
	$dateFilterModeSelection.empty()
	$dateFilterModeSelection.append(defaultSelectOptionPromptHTML("Filter for"))
	
	var dateFilterModes = {
		"any": {
			label: "Any date",
			modeSelected: function() {
				$startEndDateControls.hide()
				updateFilterRules(panelParams)
			},
			modeConfig: function() {
				var ruleConfig = { fieldID: fieldInfo.fieldID, 
					ruleID: "anyDate", 
					param: "" }
				return ruleConfig				
			}
		},
		"customRange": {
			label: "Custom date range",
			modeSelected: function() {
				$startEndDateControls.show()
				updateFilterRules(panelParams)
			},
			modeConfig: function() {
				var ruleConfig = { fieldID: fieldInfo.fieldID, 
					ruleID: "dateRange", 
					param: "" }
				return ruleConfig
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


function numberFilterPanelRuleItem(panelParams, fieldInfo) {
	
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
	
	$paramInput.blur(function() {
		var numberVal = Number($paramInput.val())
		if(!isNaN(numberVal)) {
			updateFilterRules(panelParams)	
		}
	})
	
	$paramInput.hide()
		
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
			if(ruleInfo.hasParam) {				
			} else {
				paramVal = ""
			}
			var ruleConfig = { fieldID: fieldInfo.fieldID, 
				ruleID: ruleID, 
				param: paramVal }
			return ruleConfig
		} else {
			return null
		}
	})
		
	$filterListItem.append($ruleControls)
		
	return $filterListItem
	
}


function createFilterRulePanelListItem(panelParams, fieldInfo) {
	
	switch (fieldInfo.type) {
	case fieldTypeNumber:
		return numberFilterPanelRuleItem(panelParams, fieldInfo)
	case fieldTypeTime: 
		return dateFilterPanelRuleItem(panelParams, fieldInfo)
	default:
		console.log("createFilterRulePanelListItem: Unsupported field type:  " + fieldInfo.type)
		return $("")
	}
	
}