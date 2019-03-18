// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initActiveTrackerPropertyPanel(trackerDatabaseInfo) {
	
	var $props = $('#adminGeneralActiveTracker')
		
	initCheckboxChangeHandler('#activeTrackerPropIsActive', 
		trackerDatabaseInfo.isActive, function(isActive) {
			var setDescParams = {
				databaseID:trackerDatabaseInfo.databaseID,
				isActive:isActive
			}
			jsonAPIRequest("database/setActive",setDescParams,function(dbInfo) {
			})
	})
	
		
			
}