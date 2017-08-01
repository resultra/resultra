function formComponentIsReadOnly(permissions) {
	
	if(GlobalFormPagePrivs === undefined) {
		return true
	}
	if(GlobalFormPagePrivs !== "edit") {
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