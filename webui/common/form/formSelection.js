// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function populateFormSelectionMenu(params) {
	
	var limitToFormsLookup = null
	if (params.hasOwnProperty("limitToFormIDs")) {
		limitToFormsLookup = new IDLookupTable(params.limitToFormIDs)
	}
	var includeDefaultSelection = true
	if (params.hasOwnProperty("includeDefaultSelection")) {
		includeDefaultSelection = params.includeDefaultSelection
	}
	
	
	function populateSelectionMenuFormList(formsInfo, menuSelector) {
		$(menuSelector).empty()		
		if(includeDefaultSelection) {
			$(menuSelector).append(defaultSelectOptionPromptHTML("Select a form"))		
		}
		$.each(formsInfo, function(index, formInfo) {
			
			if (limitToFormsLookup !== null) {
				if(limitToFormsLookup.hasID(formInfo.formID)) {
					$(menuSelector).append(selectFieldHTML(formInfo.formID, formInfo.name))				
				}
			} else {
				$(menuSelector).append(selectFieldHTML(formInfo.formID, formInfo.name))		
			}	
		})
	}	
	
	
	var listParams =  { parentDatabaseID: params.parentDatabaseID }
	jsonAPIRequest("frm/list",listParams,function(formsInfo) {
		
		populateSelectionMenuFormList(formsInfo,params.menuSelector)
		
		if (params.hasOwnProperty("initialFormID")) {
			$(params.menuSelector).val(params.initialFormID)
		}
	})
	
}