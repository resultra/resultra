package header

import (
	"database/sql"
	"fmt"
	"log"
	"resultra/datasheet/server/common/componentLayout"
)

// The BarChartPropertyUpdater interface along with UpdateBarChartProps() implement a harness for
// property updates. All property updates consiste of: (1) Retrieve the entity from the datastore,
// (2) Do the update on the entity itself, (3) Save the updated entity back to the datastore.
// Steps (1) and (3) can be done in a wrapper function UpdateBarChartProps, while only (2) needs
// be defined for each different property update. The goal is to minimize code bloat of property
// setting code and also make property updating code more uniform and less error prone.
type HeaderPropertyUpdater interface {
	uniqueHeaderID() string
	parentDashboardID() string
	updateHeaderProps(header *Header) error
}

type HeaderUniqueIDHeader struct {
	ParentDashboardID string `json:"parentDashboardID"`
	HeaderID          string `json:"headerID"`
}

func (idHeader HeaderUniqueIDHeader) parentDashboardID() string {
	return idHeader.ParentDashboardID
}

func (idHeader HeaderUniqueIDHeader) uniqueHeaderID() string {
	return idHeader.HeaderID
}

func updateHeaderProps(trackingDBHandle *sql.DB, propUpdater HeaderPropertyUpdater) (*Header, error) {

	// Retrieve the bar chart from the data store
	headerForUpdate, getErr := GetHeader(trackingDBHandle, propUpdater.parentDashboardID(), propUpdater.uniqueHeaderID())
	if getErr != nil {
		return nil, fmt.Errorf("updateHeaderProps: Unable to get existing header: %v", getErr)
	}

	// Do the actual update
	propUpdateErr := propUpdater.updateHeaderProps(headerForUpdate)
	if propUpdateErr != nil {
		return nil, fmt.Errorf("updateHeaderProps: Unable to update existing header: %v", propUpdateErr)
	}

	// Save the updated bar chart back to the data store
	updatedHeader, updateErr := updateExistingHeader(trackingDBHandle, headerForUpdate)
	if updateErr != nil {
		return nil, fmt.Errorf("updateHeaderProps: Unable to update existing header: %v", updateErr)
	}

	return updatedHeader, nil

}

// Title Property

type SetHeaderTitleParams struct {
	// Embed a common header to reference the BarChart in the datastore. This header also supports
	// the niqueBarChartID() method to retrieve the unique ID. So, once decoded, the struct can be passed as an
	// BarChartPropertyUpdater interface to a generic/reusable function to process the property update.
	HeaderUniqueIDHeader
	NewTitle string `json:"newTitle"`
}

func (titleParam SetHeaderTitleParams) updateHeaderProps(header *Header) error {

	log.Printf("Updating header title: %v", titleParam.NewTitle)

	header.Properties.Title = titleParam.NewTitle

	return nil
}

// Dimensions Property

type SetHeaderDimensionsParams struct {
	HeaderUniqueIDHeader
	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func (params SetHeaderDimensionsParams) updateHeaderProps(header *Header) error {

	if !componentLayout.ValidGeometry(params.Geometry) {
		return fmt.Errorf("setBarChartDimensions: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	header.Properties.Geometry = params.Geometry

	return nil
}

type SetHeaderSizeParams struct {
	HeaderUniqueIDHeader
	Size string `json:"size"`
}

func (params SetHeaderSizeParams) updateHeaderProps(header *Header) error {

	header.Properties.Size = params.Size

	return nil
}

type SetHeaderUnderlineParams struct {
	HeaderUniqueIDHeader
	Underlined bool `json:"underlined"`
}

func (params SetHeaderUnderlineParams) updateHeaderProps(header *Header) error {

	header.Properties.Underlined = params.Underlined

	return nil
}
