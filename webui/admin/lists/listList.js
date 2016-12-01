function initAdminListSettings(databaseID) {
	
    $("#adminListList").sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			// Get the new sorted list of form IDs. The prefix needs to be stripped from the ID.
	/*		var prefixRegexp = new RegExp('^' + adminFormListElemPrefix)
			var sortedIDs =  $("#adminFormList").sortable("toArray").map(function(elem) {
				return elem.replace(prefixRegexp,'')
			})
			console.log("New sort order:" + JSON.stringify(sortedIDs))
	*/
		}
    });
	
	
	var getDBInfoParams = { databaseID: databaseID }
	jsonAPIRequest("database/getInfo",getDBInfoParams,function(dbInfo) {
		console.log("Got database info: " + JSON.stringify(dbInfo))
		
/*		$('#adminFormList').empty()
		for (var formInfoIndex = 0; formInfoIndex < dbInfo.formsInfo.length; formInfoIndex++) {
			var formInfo = dbInfo.formsInfo[formInfoIndex]
			addFormToAdminFormList(formInfo)
		}
*/		
	})
	
	
	initButtonClickHandler('#adminNewListButton',function() {
		console.log("New list button clicked")
//		openNewListDialog(databaseID)
	})
	
	
	
}