// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package displayTable

import "github.com/resultra/resultra/server/generic/uniqueID"

type DisplayTableProperties struct {
	OrderedColumns []string `json:"orderedColumns"`
}

func newDefaultDisplayTableProperties() DisplayTableProperties {
	props := DisplayTableProperties{
		OrderedColumns: []string{}}
	return props
}

func (srcProps DisplayTableProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*DisplayTableProperties, error) {

	destProps := srcProps

	destProps.OrderedColumns = uniqueID.CloneIDList(remappedIDs, srcProps.OrderedColumns)

	return &destProps, nil
}
