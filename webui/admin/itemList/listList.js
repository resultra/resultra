function addListToAdminItemListList(listInfo) {
 
	var $itemListListItem = $('#adminListListItemTemplate').clone()
	$itemListListItem.attr("id","")
	
	$itemListListItem.attr("data-listID",listInfo.listID)

	var $nameLabel = $itemListListItem.find(".listNameLabel")
	$nameLabel.text(listInfo.name)
 	
	$('#adminListList').append($itemListListItem)		
}



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
	
	
	var listsInfoParams = { databaseID: databaseID }
	jsonAPIRequest("itemList/list",listsInfoParams,function(listsInfo) {
		console.log("Got item lists info: " + JSON.stringify(listsInfo))
		
		$('#adminListList').empty()
		for (var listInfoIndex = 0; listInfoIndex < listsInfo.length; listInfoIndex++) {
			var listInfo = listsInfo[listInfoIndex]
			addListToAdminItemListList(listInfo)
		}
		
	})
	
	
	initButtonClickHandler('#adminNewListButton',function() {
		console.log("New list button clicked")
		openNewListDialog(databaseID)
	})
	
	
	
}