function initAlertNotificationList(databaseID) {
	
	var getAlertListParams = {
		parentDatabaseID: databaseID
	}
	
	jsonAPIRequest("alert/getNotificationList",getAlertListParams,function(notificationList) {
		
	})
	
}