package database

import (
	"resultra/datasheet/server/generic/uniqueID"
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

func (srcProps DatabaseProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*DatabaseProperties, error) {

	destProps := srcProps

	destProps.ListOrder = uniqueID.CloneIDList(remappedIDs, srcProps.ListOrder)
	destProps.DashboardOrder = uniqueID.CloneIDList(remappedIDs, srcProps.DashboardOrder)
	destProps.FormLinkOrder = uniqueID.CloneIDList(remappedIDs, srcProps.FormLinkOrder)

	return &destProps, nil
}
