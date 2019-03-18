// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function formComponentIsReadOnly(permissions) {
	
	if(GlobalFormPagePrivs === undefined) {
		return true
	}
	if(GlobalFormPagePrivs === "edit") {
		// Page-level privileges are edit (read or write) => 
		// revert to the privileges for individual components.
		
		if(permissions.permissionMode === "readOnly") {
			return true
		} else {
			return false
		}
	} else {
		// Page-level privileges are view only => make all
		// components read-only
		return true
	}
	
}