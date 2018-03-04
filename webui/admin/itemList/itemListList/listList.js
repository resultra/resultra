function addListToAdminItemListList(databaseID,listInfo) {
 
	var $itemListListItem = $('#adminListListItemTemplate').clone()
	$itemListListItem.attr("id","")
	
	$itemListListItem.attr("data-listID",listInfo.listID)
	
	var $editListPropsButton = $itemListListItem.find(".editListPropsButton")
	$editListPropsButton.click(function(e) {
		e.preventDefault()
		$editListPropsButton.blur()
		var editPropsContentURL = '/admin/itemList/' + listInfo.listID
		setSettingsPageContent(editPropsContentURL,function() {
			initItemListPropsSettingsPageContent(databaseID,listInfo)
		})
	})

	var $nameLabel = $itemListListItem.find(".listNameLabel")
	$nameLabel.text(listInfo.name)
 	
	$('#adminListList').append($itemListListItem)		
}



function initAdminListSettings(databaseID) {
	
	
	var $adminListList = $("#adminListList")
	
    $adminListList.sortable({
		placeholder: "ui-state-highlight",
		cursor:"move",
		update: function( event, ui ) {
			
			var sortedListIDs = []
			$adminListList.find(".adminListListItem").each(function() {
				var currListID = $(this).attr('data-listID')
				sortedListIDs.push(currListID)
			})
			var setOrderParams = {
				databaseID:databaseID,
				listOrder: sortedListIDs
			}
			console.log("New list sort order:" + JSON.stringify(sortedListIDs))
			jsonAPIRequest("database/setListOrder",setOrderParams,function(dbInfo) {
				console.log("Done changing database listOrder")
			})
		}
    });
	
	
	var listsInfoParams = { databaseID: databaseID }
	jsonAPIRequest("itemList/list",listsInfoParams,function(listsInfo) {
		console.log("Got item lists info: " + JSON.stringify(listsInfo))
		
		$adminListList.empty()
		for (var listInfoIndex = 0; listInfoIndex < listsInfo.length; listInfoIndex++) {
			var listInfo = listsInfo[listInfoIndex]
			addListToAdminItemListList(databaseID,listInfo)
		}
		
	})
	
	
	initButtonClickHandler('#adminNewListButton',function() {
		console.log("New list button clicked")
		openNewListDialog(databaseID)
	})
	
	
	
}