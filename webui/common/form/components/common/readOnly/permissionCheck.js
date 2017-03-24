function formComponentIsReadOnly(permissions) {
	if(permissions.permissionMode === "readOnly") {
		return true
	} else {
		return false
	}
}