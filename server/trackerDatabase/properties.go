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
