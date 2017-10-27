

function addDashboardLinkToTOCList(tocConfig,dashboardInfo) {
	// TODO - Link to "dashboard view" page instead of dashboard design page (view page isn't implemented yet)
	var dashboardListItemHTML = '' +
		'<li>' + 
			'<a href="#" class="">' + 
					dashboardInfo.name + 
			'</a>'
		'</li>'
	var $dashboardListItem = $(dashboardListItemHTML)
	var $dashboardLink = $dashboardListItem.find("a")

	$dashboardLink.click(function(e) {
		e.preventDefault()
		console.log("Item list TOC item clicked: dashboard id = " + dashboardInfo.dashboardID)
		$dashboardLink.blur()

		$('#tocWrapper').find("li").removeClass("active")
		$dashboardListItem.addClass("active")

		if(tocConfig.dashboardClickedCallback !== undefined) {
			tocConfig.dashboardClickedCallback(dashboardInfo.dashboardID)
		}
	})
	
	$('#tocDashboardList').append($dashboardListItem)		
}

function addItemListLinkToTOCList(tocConfig, listInfo) {
	var itemListItemHTML = '' + 
		'<li>' + 
			'<a href="#">' + 
					listInfo.name + 
					'<span class="badge pull-right"></span>' +
			'</a>' +
		'</li>'
	var $itemListItem = $(itemListItemHTML)
	
	var $itemListLink = $itemListItem.find("a")
	
	$itemListLink.click(function(e) {
		e.preventDefault()
		console.log("Item list TOC item clicked: list id = " + listInfo.listID)
		$itemListLink.blur()
		
		$('#tocWrapper').find("li").removeClass("active")
		$itemListItem.addClass("active")
		
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
	
	var formLinkListItemHTML = 
		'<li>' + 
			'<a href="#"></a>' +
		'</li>'
	var $formLinkListItem = $(formLinkListItemHTML)
	var $formLinkLink = $formLinkListItem.find("a")
	$formLinkLink.text(linkInfo.name)
	
	
	$formLinkLink.click(function(e) {
		e.preventDefault()
		console.log("Form link clicked: link id = " + linkInfo.linkID)
		$formLinkLink.blur()

		$('#tocWrapper').find("li").removeClass("active")
		$formLinkListItem.addClass("active")
	
		if(tocConfig.newItemLinkClickedCallback !== undefined) {
			tocConfig.newItemLinkClickedCallback(linkInfo.linkID)
		}
	})
	
	
	$('#tocFormList').append($formLinkListItem)
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
			addDashboardLinkToTOCList(tocConfig,dashboardInfo)
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