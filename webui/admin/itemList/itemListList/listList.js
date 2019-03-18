// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function addListToAdminItemListList(databaseID,listInfo) {
 
	var $itemListListItem = $('#adminListListItemTemplate').clone()
	$itemListListItem.attr("id","")
	
	$itemListListItem.attr("data-listID",listInfo.listID)
	
	var $editListPropsButton = $itemListListItem.find(".editListPropsButton")
	$editListPropsButton.click(function(e) {
		e.preventDefault()
		$editListPropsButton.blur()
		navigateToItemListProps(databaseID,listInfo)
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