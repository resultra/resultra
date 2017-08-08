function initAdminAlertSettings(databaseID) {
	
	var $adminAlertList = $("#adminAlertList")
		
	var getAlertListParams = { databaseID: databaseID }

/*
	jsonAPIRequest("alert/getList",getAlertListParams,function(alertsInfo) {
		console.log("Got alerts info: " + JSON.stringify(dbInfo))
		
		$adminAlertList.empty()
		for (var alertInfoIndex = 0; alertInfoIndex < alertsInfo.length; alertInfoIndex++) {
			var formInfo = dbInfo.formsInfo[formInfoIndex]
//			addFormToAdminFormList(formInfo)
		}
		
	})
	*/
	
	
	initButtonClickHandler('#adminNewAlertButton',function() {
		console.log("New alert button clicked")
		openNewAlertDialog(databaseID)
	})
	
	
	
}