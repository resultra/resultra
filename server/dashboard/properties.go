// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package dashboard

import (
	"github.com/resultra/resultra/server/common/componentLayout"
	"github.com/resultra/resultra/server/trackerDatabase"
)

type DashboardProperties struct {
	Layout           componentLayout.ComponentLayout `json:"layout"`
	IncludeInSidebar bool                            `json:"includeInSidebar"`
}

func (srcProps DashboardProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*DashboardProperties, error) {

	destLayout := srcProps.Layout.Clone(cloneParams.IDRemapper)

	destProps := DashboardProperties{Layout: destLayout,
		IncludeInSidebar: srcProps.IncludeInSidebar}

	return &destProps, nil
}

func newDefaultDashboardProperties() DashboardProperties {
	props := DashboardProperties{
		Layout:           componentLayout.ComponentLayout{},
		IncludeInSidebar: true}
	return props
}
