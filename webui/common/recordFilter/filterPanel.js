// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
		limitToFieldList: filterPaneParams.limitToFieldList,
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