



function initDesignFormProperties(formID) {
	
	
	jsonAPIRequest("frm/get",{formID:formID},function(formInfo) {
		
		var formElemPrefix = "form_"
		
		initFormPropertiesFormName(formInfo)				
	})
}