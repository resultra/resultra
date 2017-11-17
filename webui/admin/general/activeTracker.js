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