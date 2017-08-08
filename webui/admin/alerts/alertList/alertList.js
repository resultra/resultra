function initAdminAlertSettings(databaseID) {
	
	var $adminAlertList = $("#adminAlertList")	
	
	function addAlertListItem(alertRef) {
		
		var $alertListItem = $('#alertListItemTemplate').clone()
		$alertListItem.attr("id","")
		
		var $alertName = $alertListItem.find("label")
		$alertName.text(alertRef.name)
		
		var $editAlertButton = $alertListItem.find(".editAlertPropsButton")
		var editAlertLink = '/admin/alert/' + alertRef.alertID
		$editAlertButton.attr("href",editAlertLink)
		
		$adminAlertList.append($alertListItem)
	}
	
	
	var getAlertsParams = { 
		parentDatabaseID: databaseID 
	}
	jsonAPIRequest("alert/list",getAlertsParams,function(alerts) {
		
		$adminAlertList.empty()
		
		$.each(alerts,function(index,alertRef) {
			addAlertListItem(alertRef)			
		})
		
	})	
	
	
	initButtonClickHandler('#adminNewAlertButton',function() {
		console.log("New alert button clicked")
		openNewAlertDialog(databaseID)
	})
	
	
	
}