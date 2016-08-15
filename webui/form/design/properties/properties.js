



function initDesignFormProperties(table,formID) {
	
	jsonAPIRequest("frm/get",{formID:formID},function(formInfo) {
		initDesignFormRolePrivProperties(formID)
		initDesignFormFilterProperties(tableID,formInfo)
		initFormPropertiesFormName(formInfo)
	})
}