// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package trackerDatabase

import (
	"resultra/tracker/server/generic/uniqueID"
)

type DatabaseProperties struct {
	ListOrder      []string `json:"listOrder"`
	DashboardOrder []string `json:"dashboardOrder"`
	FormLinkOrder  []string `json:"formLinkOrder"`
}

func newDefaultDatabaseProperties() DatabaseProperties {
	props := DatabaseProperties{
		ListOrder:      []string{},
		DashboardOrder: []string{},
		FormLinkOrder:  []string{}}
	return props
}

func (srcProps DatabaseProperties) Clone(cloneParams *CloneDatabaseParams) (*DatabaseProperties, error) {

	destProps := srcProps

	destProps.ListOrder = uniqueID.CloneIDList(cloneParams.IDRemapper, srcProps.ListOrder)
	destProps.DashboardOrder = uniqueID.CloneIDList(cloneParams.IDRemapper, srcProps.DashboardOrder)
	destProps.FormLinkOrder = uniqueID.CloneIDList(cloneParams.IDRemapper, srcProps.FormLinkOrder)

	return &destProps, nil
}
