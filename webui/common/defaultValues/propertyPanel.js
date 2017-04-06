function initDefaultValuesPropertyPanel(panelParams) {
		
	var fieldSelectionDropdownParams = {
		elemPrefix: panelParams.elemPrefix,
		databaseID: panelParams.databaseID,
		fieldTypes: [fieldTypeBool,fieldTypeNumber,fieldTypeTime,fieldTypeText],
		fieldSelectionCallback: function(fieldInfo) {
			
			var defaultValuesListSelector = createPrefixedSelector(panelParams.elemPrefix,
							'DefaultValuesList')
			var $defaultValueList = $(defaultValuesListSelector)
			
			// Use null to signify no default rule information. This is true when
			// creating new rules, but will not be when re-loading the rules.
			var defaultValInfo = null
				
			$defaultValueList.append(createDefaultValuePanelListItem(panelParams,fieldInfo,defaultValInfo))
		}
	}
	initFieldSelectionDropdown(fieldSelectionDropdownParams)
	
	initDefaultDefaultValuePanelItems(panelParams)	
	
}
