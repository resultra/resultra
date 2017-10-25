package dashboard

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/trackerDatabase"
)

type DashboardProperties struct {
	Layout componentLayout.ComponentLayout `json:"layout"`
}

func (srcProps DashboardProperties) Clone(cloneParams *trackerDatabase.CloneDatabaseParams) (*DashboardProperties, error) {

	destLayout := srcProps.Layout.Clone(cloneParams.IDRemapper)
	destProps := DashboardProperties{Layout: destLayout}

	return &destProps, nil
}
