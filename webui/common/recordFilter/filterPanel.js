
function initRecordFilterViewPanel(filterPaneParams) {
		
	initDefaultFilterRules(filterPaneParams) 	
	
	var resetFiltersButtonSelector = createPrefixedSelector(filterPaneParams.elemPrefix,'FilterRecordsResetFiltersButton')
	initButtonClickHandler(resetFiltersButtonSelector,function () {
		updateDefaultFilterRules(filterPaneParams, function () {
			filterPaneParams.updateFilterRules(filterPaneParams.defaultFilterRules)
			filterPaneParams.refilterWithCurrentFilterRules()		
		})		
	})
	
	var fieldSelectionDropdownParams = {
		elemPrefix: filterPaneParams.elemPrefix,
		databaseID: filterPaneParams.databaseID,
		fieldTypes: [fieldTypeAll],
		fieldSelectionCallback: function(fieldInfo) {
			var filterRuleListSelector = createPrefixedSelector(filterPaneParams.elemPrefix,
							'RecordFilterFilterRuleList')
			var $filterRuleList = $(filterRuleListSelector)
			
			// Use null to signify no default rule information. This is true when
			// creating new rules, but will not be when re-loading the rules.
			var defaultRuleInfo = null
			$filterRuleList.append(createFilterRulePanelListItem(filterPaneParams,fieldInfo,defaultRuleInfo))
			updateMatchLogicSelectionVisibility(filterPaneParams)
		}
	}
	initFieldSelectionDropdown(fieldSelectionDropdownParams)
	
	
}