package dashboard

import (
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/generic/uniqueID"
)

type DashboardProperties struct {
	Layout componentLayout.ComponentLayout `json:"layout"`
}

func (srcProps DashboardProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*DashboardProperties, error) {

	destLayout := srcProps.Layout.Clone(remappedIDs)
	destProps := DashboardProperties{Layout: destLayout}

	return &destProps, nil
}
