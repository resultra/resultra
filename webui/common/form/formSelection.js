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