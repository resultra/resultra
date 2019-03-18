// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initDefaultValuesPropertyPanel(panelParams) {
		
	var fieldSelectionDropdownParams = {
		elemPrefix: panelParams.elemPrefix,
		databaseID: panelParams.databaseID,
		fieldTypes: [fieldTypeBool,fieldTypeNumber,fieldTypeTime,fieldTypeText],
		includeCalcFields: false,
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
