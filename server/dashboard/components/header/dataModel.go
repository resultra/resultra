package header

import (
	"database/sql"
	"fmt"
	"resultra/datasheet/server/common/componentLayout"
	"resultra/datasheet/server/dashboard/components/common"
	"resultra/datasheet/server/generic"
	"resultra/datasheet/server/generic/uniqueID"
	"resultra/datasheet/server/trackerDatabase"
)

const headerEntityKind string = "header"

// DashboardBarChart is the datastore object for dashboard bar charts.
type Header struct {
	ParentDashboardID string `json:"parentDashboardID"`

	HeaderID string `json:"headerID"`

	// DataSrcTable is the table the bar chart gets its data from
	Properties HeaderProps `json:"properties"`
}

type NewHeaderParams struct {
	ParentDashboardID string `json:"parentDashboardID"`

	Geometry componentLayout.LayoutGeometry `json:"geometry"`
}

func saveHeader(destDBHandle *sql.DB, newHeader Header) error {

	if saveErr := common.SaveNewDashboardComponent(destDBHandle, headerEntityKind,
		newHeader.ParentDashboardID, newHeader.HeaderID, newHeader.Properties); saveErr != nil {
		return fmt.Errorf("newHeader: Unable to save header component: error = %v", saveErr)
	}
	return nil

}

func newHeader(trackerDBHandle *sql.DB, params NewHeaderParams) (*Header, error) {

	if len(params.ParentDashboardID) <= 0 {
		return nil, fmt.Errorf("newHeader: Error creating summary table: missing parent dashboard ID")
	}

	if !componentLayout.ValidGeometry(params.Geometry) {
		return nil, fmt.Errorf("newHeader: Invalid geometry for bar chart: %+v", params.Geometry)
	}

	headerProps := newDefaultHeaderProps()
	headerProps.Geometry = params.Geometry

	newHeader := Header{
		ParentDashboardID: params.ParentDashboardID,
		HeaderID:          uniqueID.GenerateSnowflakeID(),
		Properties:        headerProps}

	if saveErr := saveHeader(trackerDBHandle, newHeader); saveErr != nil {
		return nil, fmt.Errorf("newHeader: Unable to save summary component with params=%+v: error = %v", params, saveErr)
	}

	return &newHeader, nil
}

func GetHeader(trackerDBHandle *sql.DB, parentDashboardID string, headerID string) (*Header, error) {

	headerProps := newDefaultHeaderProps()
	if getErr := common.GetDashboardComponent(trackerDBHandle, headerEntityKind, parentDashboardID, headerID, &headerProps); getErr != nil {
		return nil, fmt.Errorf("getBarChart: Unable to retrieve bar chart component: %v", getErr)
	}

	header := Header{
		ParentDashboardID: parentDashboardID,
		HeaderID:          headerID,
		Properties:        headerProps}

	return &header, nil

}

func getHeadersFromSrc(srcDBHandle *sql.DB, parentDashboardID string) ([]Header, error) {

	headers := []Header{}
	addHeader := func(headerID string, encodedProps string) error {

		headerProps := newDefaultHeaderProps()
		if decodeErr := generic.DecodeJSONString(encodedProps, &headerProps); decodeErr != nil {
			return fmt.Errorf("GetHeaders: can't decode properties: %v", encodedProps)
		}

		currHeader := Header{
			ParentDashboardID: parentDashboardID,
			HeaderID:          headerID,
			Properties:        headerProps}

		headers = append(headers, currHeader)

		return nil
	}
	if getErr := common.GetDashboardComponents(srcDBHandle, headerEntityKind, parentDashboardID, addHeader); getErr != nil {
		return nil, fmt.Errorf("getBarCharts: Can't get bar chart components: %v")
	}

	return headers, nil
}

func GetHeaders(trackerDBHandle *sql.DB, parentDashboardID string) ([]Header, error) {
	return getHeadersFromSrc(trackerDBHandle, parentDashboardID)
}

func CloneHeaders(cloneParams *trackerDatabase.CloneDatabaseParams, srcParentDashboardID string) error {

	remappedDashboardID, err := cloneParams.IDRemapper.GetExistingRemappedID(srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneHeaders: %v", err)
	}

	headers, err := getHeadersFromSrc(cloneParams.SrcDBHandle, srcParentDashboardID)
	if err != nil {
		return fmt.Errorf("CloneHeaders: %v", err)
	}

	for _, srcHeader := range headers {

		remappedHeaderID, err := cloneParams.IDRemapper.AllocNewRemappedID(srcHeader.HeaderID)
		if err != nil {
			return fmt.Errorf("CloneHeaders: %v", err)
		}

		clonedProps, err := srcHeader.Properties.Clone(cloneParams)
		if err != nil {
			return fmt.Errorf("CloneHeaders: %v", err)
		}

		destHeader := Header{
			ParentDashboardID: remappedDashboardID,
			HeaderID:          remappedHeaderID,
			Properties:        *clonedProps}

		if err := saveHeader(cloneParams.DestDBHandle, destHeader); err != nil {
			return fmt.Errorf("CloneHeaders: %v", err)
		}
	}

	return nil
}

func updateExistingHeader(trackerDBHandle *sql.DB, updatedHeader *Header) (*Header, error) {

	if updateErr := common.UpdateDashboardComponent(trackerDBHandle, headerEntityKind, updatedHeader.ParentDashboardID,
		updatedHeader.HeaderID, updatedHeader.Properties); updateErr != nil {
		return nil, fmt.Errorf("Error updating summary table %+v: %v", updatedHeader, updateErr)
	}

	return updatedHeader, nil

}
