

function initFilterPropertyPanel(panelParams) {
		
	var fieldSelectionDropdownParams = {
		elemPrefix: panelParams.elemPrefix,
		tableID: panelParams.tableID,
		fieldTypes: [fieldTypeAll],
		fieldSelectionCallback: function(fieldInfo) {
			var filterRuleListSelector = createPrefixedSelector(panelParams.elemPrefix,
							'RecordFilterFilterRuleList')
			var $filterRuleList = $(filterRuleListSelector)		
			$filterRuleList.append(createFilterRulePanelListItem(panelParams,fieldInfo))
		}
	}
	initFieldSelectionDropdown(fieldSelectionDropdownParams)	
	
}
