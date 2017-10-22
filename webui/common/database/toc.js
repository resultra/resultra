

function addDashboardLinkToTOCList(dashboardInfo) {
	// TODO - Link to "dashboard view" page instead of dashboard design page (view page isn't implemented yet)
	var dashboardListItemHTML = '<a href="/viewDashboard/' + dashboardInfo.dashboardID 
			+ '" class="list-group-item">' + dashboardInfo.name + '</a>'
	$('#tocDashboardList').append(dashboardListItemHTML)		
}

function addItemListLinkToTOCList(tocConfig, listInfo) {
	var itemListItemHTML = '' + 
		'<a href="" class="list-group-item">' + 
				listInfo.name + 
			'<span class="badge"></span>' +
		'</a>'
	var $itemListItem = $(itemListItemHTML)
	
	$itemListItem.click(function(e) {
		e.preventDefault()
		console.log("Item list TOC item clicked: list id = " + listInfo.listID)
		$itemListItem.blur()
		
		if(tocConfig.itemListClickedCallback !== undefined) {
			tocConfig.itemListClickedCallback(listInfo.listID)
		}
	})
	
	var listCountParams = {
		databaseID: tocConfig.databaseID,
		preFilterRules: listInfo.properties.preFilterRules,
	}
	
	jsonAPIRequest("recordRead/getFilteredRecordCount",listCountParams,function(listCount) {
		var $listCount = $itemListItem.find("span")
		if (listCount > 0) {
			$listCount.text(listCount)
		} else {
			$listCount.hide()
		}
	})
	
	
	$('#tocListList').append($itemListItem)		
	
}

function addFormLinkToTOCList(tocConfig, linkInfo) {
	
	var formLinkListItemHTML = '<a href="/newItem/' + linkInfo.linkID 
			+ '" class="list-group-item">' + linkInfo.name + '</a>'
	$('#tocFormList').append(formLinkListItemHTML)
}


function initDatabaseTOC(tocConfig) {
	
	
	$('#tocFormList').empty()
	var linkParams = { parentDatabaseID: tocConfig.databaseID }
	jsonAPIRequest("formLink/getUserList",linkParams,function(linkList) {
		for(var linkIndex = 0; linkIndex < linkList.length; linkIndex++) {
			var currLink = linkList[linkIndex]
			var numLinksInSidebar = 0
			if(currLink.includeInSidebar) {
				numLinksInSidebar++
				addFormLinkToTOCList(tocConfig,currLink)		
			}
			if(numLinksInSidebar === 0) {
				$('#tocNewItemContainer').hide()
			}
		}
	})
	
	
	var getDBInfoParams = { databaseID: tocConfig.databaseID }
	jsonAPIRequest("dashboard/getUserDashboardList",getDBInfoParams,function(dashboardsInfo) {
		console.log("Got dashboard info: " + JSON.stringify(dashboardsInfo))		
		$('#tocDashboardList').empty()
		for (var dashboardInfoIndex = 0; dashboardInfoIndex < dashboardsInfo.length; dashboardInfoIndex++) {
			var dashboardInfo = dashboardsInfo[dashboardInfoIndex]
			addDashboardLinkToTOCList(dashboardInfo)
		}
		if(dashboardsInfo.length === 0) {
			$('#tocDashboardsContainer').hide()
		}
	})
	
	jsonAPIRequest("itemList/getUserItemListList",getDBInfoParams,function(listsInfo) {
		console.log("Got database info: " + JSON.stringify(listsInfo))		
		
		$('#tocListList').empty()
		for(var listInfoIndex = 0; listInfoIndex < listsInfo.length; listInfoIndex++) {
			var listInfo = listsInfo[listInfoIndex]
			addItemListLinkToTOCList(tocConfig,listInfo)
		}
		if (listsInfo.length === 0) {
			$('#tocListsContainer').hide()
		}
		
	}) // getRecord
	
}