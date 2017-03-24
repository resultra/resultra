function initFormComponentReadOnlyPropertyPanel(params) {
	
	var checkboxSelector = '#' + params.elemPrefix + "FormComponentReadOnlyProperty"
	
	var isReadOnly = true
	if(params.initialVal.permissionMode === "readWrite") {
		isReadOnly = false
	}
	
	initCheckboxChangeHandler(checkboxSelector, isReadOnly, function(isReadOnly) {
		
		var permMode = "readWrite"
		if(isReadOnly) {
			permMode = "readOnly"
		}
		
		var permissions = {
			permissionMode: permMode
		}
		
		params.permissionsChangedCallback(permissions)
		
	})
	
}