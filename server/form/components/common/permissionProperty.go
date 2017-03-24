package common

const permModeReadOnly string = "readOnly"
const permModeReadWrite string = "readWrite"

// TODO const permModeConditional string = "conditional"

type ComponentValuePermissionsProperties struct {
	PermissionMode string `json:"permissionMode"`
}

func NewDefaultComponentValuePermissionsProperties() ComponentValuePermissionsProperties {
	return ComponentValuePermissionsProperties{
		PermissionMode: permModeReadWrite}
}
