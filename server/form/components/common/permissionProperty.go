// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
