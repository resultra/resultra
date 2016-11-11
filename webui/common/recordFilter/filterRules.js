
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

function createFilterListRuleListItem(fieldName) {
	
	var filterListItemHTML =  '' +
		'<div class="list-group-item recordFilterPanelRuleListItem">' +
			'<label>' +
				fieldName +
			'</label>' +
		'</div>';
	
	var $filterRuleListItem = $(filterListItemHTML)
		
	return $filterRuleListItem
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
		
		
	var $filterListItem = createFilterListRuleListItem(fieldInfo.name)
		
	$filterListItem.data("filterRuleConfigFunc",function() {
		var ruleID = $ruleSelection.val()
		var paramVal = $paramInput.val()
		if(ruleID.length > 0) {
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
	default:
		console.log("createFilterRulePanelListItem: Unsupported field type:  " + fieldInfo.type)
		return $("")
	}
	
}