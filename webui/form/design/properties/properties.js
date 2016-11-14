



function initDesignFormProperties(table,formID) {
	
	
	jsonAPIRequest("frm/get",{formID:formID},function(formInfo) {
		
		var formElemPrefix = "form_"
		
		initDesignFormRolePrivProperties(formID)
		initFormPropertiesFormName(formInfo)
		initFormSortPropertyPanel(formInfo)
		
		
		var filterPropertyPanelParams = {
			elemPrefix: formElemPrefix,
			tableID: tableID,
			/* TODO - restore a callback with functions like the following:
				var setDefaultFiltersParams = {
					formID: formID,
					defaultFilterIDs: defaultFilterIDs
				}
			*/
		}
		initFilterPropertyPanel(filterPropertyPanelParams)
		
		
		
	})
}