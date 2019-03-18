// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.



function initFilterPropertyPanel(panelParams) {
		
	var fieldSelectionDropdownParams = {
		elemPrefix: panelParams.elemPrefix,
		databaseID: panelParams.databaseID,
		fieldTypes: [fieldTypeAll],
		fieldSelectionCallback: function(fieldInfo) {
			var filterRuleListSelector = createPrefixedSelector(panelParams.elemPrefix,
							'RecordFilterFilterRuleList')
			var $filterRuleList = $(filterRuleListSelector)
			
			// Use null to signify no default rule information. This is true when
			// creating new rules, but will not be when re-loading the rules.
			var defaultRuleInfo = null
			$filterRuleList.append(createFilterRulePanelListItem(panelParams,fieldInfo,defaultRuleInfo))
			updateMatchLogicSelectionVisibility(panelParams)
		}
	}
	initFieldSelectionDropdown(fieldSelectionDropdownParams)
	
	initMatchLogicSelection(panelParams)	
	
	initDefaultFilterRules(panelParams, function() {})	
	
}
