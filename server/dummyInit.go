package server

import (
	"resultra/datasheet/server/adminController"
	"resultra/datasheet/server/dashboardController"
	"resultra/datasheet/server/database"
	"resultra/datasheet/server/databaseController"
	"resultra/datasheet/server/form"
	"resultra/datasheet/server/global"
	"resultra/datasheet/server/newItem"
	"resultra/datasheet/server/recordReadController"
	"resultra/datasheet/server/recordUpdate"
	"resultra/datasheet/server/timelineController"
)

// Dummy variables to force inclusion of the packages (and not trigger an error from the Golang compiler).
// This is needed since these packages are essentially plug-ins which register their own HTTP handlers upon startup.
var dummyUnusedDBParams = database.NewDatabaseParams{}
var dummyRecordUpdateParams = recordUpdate.DummyStructForInclude{}
var dummyRecordVals = recordReadController.DummyStructForInclude{}
var dummyDashboardControllerVals = dashboardController.DummyStructForInclude{}
var dummDBInfo = databaseController.DummyStructForInclude{}
var dummyFormInfo = form.DummyStructForInclude{}
var dummyAdminInfo = adminController.DummyStructForInclude{}
var dummyGlobalInfo = global.DummyStructForInclude{}
var dummyTimelineInfo = timelineController.DummyStructForInclude{}
var dummyNewItemInfo = newItem.DummyStructForInclude{}

func DummyFunctionForImportFromGoogleAppEngineProjectFolder() {
	// This dummy function is needed so standaline packages inside
	// the server will be compiled into the google app engine executable.
	// The stand-alone packages won't be compiled in unless they are included somewhere.
}
