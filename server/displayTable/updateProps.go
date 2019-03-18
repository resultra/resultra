// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package displayTable

import (
	"database/sql"
	"fmt"
)

type TableIDInterface interface {
	getTableID() string
}

type TableIDHeader struct {
	TableID string `json:"tableID"`
}

func (idHeader TableIDHeader) getTableID() string {
	return idHeader.TableID
}

type TablePropUpdater interface {
	TableIDInterface
	updateProps(table *DisplayTable) error
}

func updateTableProps(trackerDBHandle *sql.DB, propUpdater TablePropUpdater) (*DisplayTable, error) {

	// Retrieve the bar chart from the data store
	tableForUpdate, getErr := GetTable(trackerDBHandle, propUpdater.getTableID())
	if getErr != nil {
		return nil, fmt.Errorf("updateTableProps: Unable to get existing table: %v", getErr)
	}

	if propUpdateErr := propUpdater.updateProps(tableForUpdate); propUpdateErr != nil {
		return nil, fmt.Errorf("updateTableProps: Unable to update existing table properties: %v", propUpdateErr)
	}

	table, updateErr := updateExistingTable(trackerDBHandle, propUpdater.getTableID(), tableForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateTableProps: Unable to update existing table properties: datastore update error =  %v", updateErr)
	}

	return table, nil
}

type SetTableNameParams struct {
	TableIDHeader
	NewTableName string `json:"newTableName"`
}

func (updateParams SetTableNameParams) updateProps(table *DisplayTable) error {

	// TODO - Validate name

	table.Name = updateParams.NewTableName

	return nil
}

type SetOrderedColParams struct {
	TableIDHeader
	OrderedColumns []string `json:"orderedColumns"`
}

func (updateParams SetOrderedColParams) updateProps(table *DisplayTable) error {

	// TODO - Validate each ID

	table.Properties.OrderedColumns = updateParams.OrderedColumns

	return nil
}
