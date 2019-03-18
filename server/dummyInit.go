// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package server

import (
	"resultra/tracker/server/adminController"
	"resultra/tracker/server/common/attachment"
	"resultra/tracker/server/dashboardController"
	"resultra/tracker/server/databaseController"
	"resultra/tracker/server/displayTable"
	"resultra/tracker/server/form"
	"resultra/tracker/server/formLink"
	"resultra/tracker/server/global"
	"resultra/tracker/server/recordReadController"
	"resultra/tracker/server/recordUpdate"
	"resultra/tracker/server/timelineController"
	"resultra/tracker/server/trackerDatabase"
	"resultra/tracker/server/userRoleController"
	"resultra/tracker/server/valueList"
)

// Dummy variables to force inclusion of the packages (and not trigger an error from the Golang compiler).
// This is needed since these packages are essentially plug-ins which register their own HTTP handlers upon startup.
var dummyUnusedDBParams = trackerDatabase.NewDatabaseParams{}
var dummyRecordUpdateParams = recordUpdate.DummyStructForInclude{}
var dummyRecordVals = recordReadController.DummyStructForInclude{}
var dummyDashboardControllerVals = dashboardController.DummyStructForInclude{}
var dummDBInfo = databaseController.DummyStructForInclude{}
var dummyFormInfo = form.DummyStructForInclude{}
var dummyAdminInfo = adminController.DummyStructForInclude{}
var dummyGlobalInfo = global.DummyStructForInclude{}
var dummyTimelineInfo = timelineController.DummyStructForInclude{}
var dummyNewItemInfo = formLink.DummyStructForInclude{}
var dummyUserRoleInfo = userRoleController.DummyStructForInclude{}
var dummyAttachmentInfo = attachment.DummyStructForInclude{}
var dummyValueListInfo = valueList.DummyStructForInclude{}
var dummyDisplayTableInfo = displayTable.DummyStructForInclude{}

func DummyFunctionForImportFromGoogleAppEngineProjectFolder() {
	// This dummy function is needed so standaline packages inside
	// the server will be compiled into the google app engine executable.
	// The stand-alone packages won't be compiled in unless they are included somewhere.
}
