package dashboard

import (
	"appengine"
	"fmt"
	"log"
	"resultra/datasheet/datamodel"
)

const dashboardEntityKind string = "Dashboard"

type Dashboard struct {
	Name string `json:"name"`
}

type DashboardRef struct {
	DashboardID string `json:"dashboardID"`
	Name        string `json:"name"`
}

func NewDashboard(appEngContext appengine.Context, dashboardName string) (*DashboardRef, error) {

	sanitizedName, sanitizeErr := datamodel.SanitizeName(dashboardName)
	if sanitizeErr != nil {
		return nil, sanitizeErr
	}

	var newDashboard = Dashboard{sanitizedName}
	dashboardID, insertErr := datamodel.InsertNewEntity(appEngContext, dashboardEntityKind, nil, &newDashboard)
	if insertErr != nil {
		return nil, insertErr
	}

	log.Printf("NewDashboard: Created new dashboard: id= %v, name='%v'", dashboardID, sanitizedName)
	dashboardRef := DashboardRef{dashboardID, sanitizedName}

	return &dashboardRef, nil

}

func GetDashboardRef(appEngContext appengine.Context, dashboardID string) (*DashboardRef, error) {

	var dashboard DashboardRef
	getErr := datamodel.GetRootEntityByID(appEngContext, dashboardEntityKind, dashboardID, &dashboard)
	if getErr != nil {
		return nil, fmt.Errorf("GetDashboardRef: Can't get dashboard: Error retrieving existing dashboard: dashboard ID=%v, err = %v", dashboardID, getErr)
	}

	return &DashboardRef{dashboardID, dashboard.Name}, nil

}
