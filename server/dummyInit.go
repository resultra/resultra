// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
package server

import (
	"github.com/resultra/resultra/server/adminController"
	"github.com/resultra/resultra/server/common/attachment"
	"github.com/resultra/resultra/server/dashboardController"
	"github.com/resultra/resultra/server/databaseController"
	"github.com/resultra/resultra/server/displayTable"
	"github.com/resultra/resultra/server/form"
	"github.com/resultra/resultra/server/formLink"
	"github.com/resultra/resultra/server/global"
	"github.com/resultra/resultra/server/recordReadController"
	"github.com/resultra/resultra/server/recordUpdate"
	"github.com/resultra/resultra/server/timelineController"
	"github.com/resultra/resultra/server/trackerDatabase"
	"github.com/resultra/resultra/server/userRoleController"
	"github.com/resultra/resultra/server/valueList"
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
