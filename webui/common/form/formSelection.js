function populateFormSelectionMenu(params) {
	
	function populateSelectionMenuFormList(formsInfo, menuSelector) {
		$(menuSelector).empty()		
		$(menuSelector).append(defaultSelectOptionPromptHTML("Select a form"))
		$.each(formsInfo, function(index, formInfo) {
			$(menuSelector).append(selectFieldHTML(formInfo.formID, formInfo.name))		
		})
	}	
	
	
	var listParams =  { parentTableID: params.parentTableID }
	jsonAPIRequest("frm/list",listParams,function(formsInfo) {
		
		populateSelectionMenuFormList(formsInfo,params.menuSelector)
		
		if (params.hasOwnProperty("initialFormID")) {
			$(params.menuSelector).val(params.initialFormID)
		}
	})
	
}