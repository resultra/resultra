// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initFormComponentPermissionsPropertyPanel(params) {
	
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