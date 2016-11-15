



function initDesignFormProperties(table,formID) {
	
	
	jsonAPIRequest("frm/get",{formID:formID},function(formInfo) {
		
		var formElemPrefix = "form_"
		
		initDesignFormRolePrivProperties(formID)
		initFormPropertiesFormName(formInfo)
		initFormSortPropertyPanel(formInfo)
		
		
		var filterPropertyPanelParams = {
			elemPrefix: formElemPrefix,
			tableID: tableID,
			defaultFilterRules: formInfo.properties.defaultFilterRules,
			updateFilterRules: function (updatedFilterRules) {
				var setDefaultFiltersParams = {
					formID: formID,
					filterRules: updatedFilterRules
				}
				jsonAPIRequest("frm/setDefaultFilterRules",setDefaultFiltersParams,function(updatedForm) {
					console.log(" Default filters updated")
				}) // set record's number field value
				
			}
			
		}
		initFilterPropertyPanel(filterPropertyPanelParams)
		
		
		
	})
}