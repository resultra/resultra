


function recordFilterPanelRuleItem(panelParams, fieldInfo) {
	
	function updateFilterRules() {
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
			updateFilterRules()	
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
			updateFilterRules()
		}
	})
		
	var filterListItemHTML =  '' +
		'<div class="list-group-item recordFilterPanelRuleListItem">' +
			'<label>' +
				fieldInfo.name +
			'</label>' +
		'</div>';
		
	var $filterListItem = $(filterListItemHTML)
		
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



function initFilterPropertyPanel(panelParams) {
	
	var defaultFilterListSelector = createPrefixedSelector(panelParams.elemPrefix,'DefaultFilterList')			
	var availListSelector = createPrefixedSelector(panelParams.elemPrefix,'AvailableFilterList')	
	
	function changeOneAvailableFilterSelection(filterID, isChecked) {
	
		var selectedFilterIDs = getFilterCheckboxListSelectedFilterIDs(availListSelector)
		console.log("Available filters: updated selection: " + JSON.stringify(selectedFilterIDs))
		
		panelParams.setAvailableFilterFunc(selectedFilterIDs)
	
	}
	
	function changeOneDefaultFilterSelection(filterID, isChecked) {
	
		var selectedFilterIDs = getFilterCheckboxListSelectedFilterIDs(defaultFilterListSelector)
		console.log("Available filters: updated selection: " + JSON.stringify(selectedFilterIDs))
		
		panelParams.setDefaultFilterFunc(selectedFilterIDs)
	
	}
	
	jsonAPIRequest("filter/getList",{parentTableID:panelParams.tableID},function(filterList) {
							
		initializeFilterCheckboxList(defaultFilterListSelector,'defaultFilters_',
			filterList,panelParams.defaultFilterIDs,changeOneDefaultFilterSelection)
	
		initializeFilterCheckboxList(availListSelector,'availFilters_',
			filterList,panelParams.availableFilterIDs,changeOneAvailableFilterSelection)
		
	})
	
	
	var manageFiltersSelector = createPrefixedSelector(panelParams.elemPrefix,'FilterRecordsManageFiltersButton')	
	initButtonClickHandler(manageFiltersSelector,function() {
		openRecordFilterManageFiltersDialog(panelParams.tableID)	
	}) 
		
	var fieldSelectionDropdownParams = {
		elemPrefix: panelParams.elemPrefix,
		tableID: panelParams.tableID,
		fieldTypes: [fieldTypeAll],
		fieldSelectionCallback: function(fieldInfo) {
			var filterRuleListSelector = createPrefixedSelector(panelParams.elemPrefix,
							'RecordFilterFilterRuleList')
			var $filterRuleList = $(filterRuleListSelector)		
			$filterRuleList.append(recordFilterPanelRuleItem(panelParams,fieldInfo))
		}
	}
	initFieldSelectionDropdown(fieldSelectionDropdownParams)	
	
}
