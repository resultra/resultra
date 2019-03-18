// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function getGlobalInfoIndexedByID(databaseID, globalInfoIndexCallback) {
	
	var listParams =  { parentDatabaseID: databaseID }
	var globalInfoIndex = {}
	jsonAPIRequest("global/getList",listParams,function(globalsInfo) {
		$.each(globalsInfo, function(index, globalInfo) {
			globalInfoIndex[globalInfo.globalID] = globalInfo	
		})
		globalInfoIndexCallback(globalInfoIndex)
	})
}