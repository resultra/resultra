

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
		}
	}
	initFieldSelectionDropdown(fieldSelectionDropdownParams)
	
	initDefaultFilterRules(panelParams)	
	
}
