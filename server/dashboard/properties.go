package dashboard

import (
	"resultra/tracker/server/common/componentLayout"
	"resultra/tracker/server/trackerDatabase"
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
