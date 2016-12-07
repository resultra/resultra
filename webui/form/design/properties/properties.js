



function initDesignFormProperties(table,formID) {
	
	
	jsonAPIRequest("frm/get",{formID:formID},function(formInfo) {
		
		var formElemPrefix = "form_"
		
		initDesignFormRolePrivProperties(formID)
		initFormPropertiesFormName(formInfo)
				
	})
}