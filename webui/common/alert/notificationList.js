function loadNotificationListInfo(databaseID, loadDoneCallback) {
	var loadsRemaining = 3
	
	var notificationList
	var formsByID
	var fieldsByID
	
	function oneLoadComplete() {
		loadsRemaining--
		if (loadsRemaining <= 0) {
			loadDoneCallback(notificationList,formsByID,fieldsByID)
		}
	}
	
	var getAlertListParams = { parentDatabaseID: databaseID }
	jsonAPIRequest("alert/getNotificationList",getAlertListParams,function(notListReply) {
		notificationList = notListReply
		oneLoadComplete()
	})

	var getFormsParams = { parentDatabaseID: databaseID }
	jsonAPIRequest("frm/formsByID",getFormsParams,function(formsByIDReply) {
		formsByID = formsByIDReply
		oneLoadComplete()
	})
	
	loadFieldInfo(databaseID,[fieldTypeAll],function(fieldsByIDReply) {
		fieldsByID = fieldsByIDReply
		oneLoadComplete()			
	})		
	
}
