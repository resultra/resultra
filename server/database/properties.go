package database

import (
	"resultra/datasheet/server/generic/uniqueID"
)

type DatabaseProperties struct {
	ListOrder      []string `json:"listOrder"`
	DashboardOrder []string `json:"dashboardOrder"`
}

func newDefaultDatabaseProperties() DatabaseProperties {
	props := DatabaseProperties{
		ListOrder:      []string{},
		DashboardOrder: []string{}}
	return props
}

func (srcProps DatabaseProperties) Clone(remappedIDs uniqueID.UniqueIDRemapper) (*DatabaseProperties, error) {

	destProps := srcProps

	destProps.ListOrder = uniqueID.CloneIDList(remappedIDs, srcProps.ListOrder)
	destProps.DashboardOrder = uniqueID.CloneIDList(remappedIDs, srcProps.DashboardOrder)

	return &destProps, nil
}
