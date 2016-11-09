



function initDesignFormProperties(table,formID) {
	
	
	jsonAPIRequest("frm/get",{formID:formID},function(formInfo) {
		
		var formElemPrefix = "form_"
		
		initDesignFormRolePrivProperties(formID)
		initFormPropertiesFormName(formInfo)
		initFormSortPropertyPanel(formInfo)
		
		
		var filterPropertyPanelParams = {
			elemPrefix: formElemPrefix,
			tableID: tableID,
			defaultFilterIDs: formInfo.properties.defaultFilterIDs,
			setDefaultFilterFunc: function(defaultFilterIDs) {
				
				var setDefaultFiltersParams = {
					formID: formID,
					defaultFilterIDs: defaultFilterIDs
				}
		
				jsonAPIRequest("frm/setDefaultFilters",setDefaultFiltersParams,function(updatedForm) {
					console.log(" Default filters updated")
				}) // set record's number field value
							
			},
			availableFilterIDs: formInfo.properties.availableFilterIDs,
			setAvailableFilterFunc: function(availFilterIDs) {
				var setAvailFiltersParams = {
					formID: formID,
					availableFilterIDs: availFilterIDs
				}
		
				jsonAPIRequest("frm/setAvailableFilters",setAvailFiltersParams,function(updatedForm) {
					console.log("Available filters updated")
				}) // set record's number field value
			
			}
		}
		initFilterPropertyPanel(filterPropertyPanelParams)
		
		
		
	})
}