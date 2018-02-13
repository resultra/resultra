
function setDefaultTOCItem(databaseID, itemID) {
	var itemKey = "resultra-default-toc-item:" + databaseID
	localStorage.setItem(itemKey,itemID)
}

function getDefaultTOCItem(databaseID) {
	var itemKey = "resultra-default-toc-item:" + databaseID
	return localStorage.getItem(itemKey)
}

function initTableOfContentsRefreshPollingLoop(refreshCallback) {
	initRefreshPollingLoop(5,refreshCallback)
}

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
			tocConfig.dashboardClickedCallback(dashboardInfo.dashboardID,$dashboardListItem)
		}
		
		setDefaultTOCItem(tocConfig.databaseID,dashboardInfo.dashboardID)
	})
	
	if(getDefaultTOCItem(tocConfig.databaseID)===dashboardInfo.dashboardID && 
		tocConfig.dashboardClickedCallback !== undefined) {
			tocConfig.dashboardClickedCallback(dashboardInfo.dashboardID,$dashboardListItem)
	}
	
	
	$('#tocDashboardList').append($dashboardListItem)		
}

function addItemListLinkToTOCList(tocConfig, listInfo) {
	
	
	var $itemListItem = $('#itemListTOCItemTemplate').clone()
	$itemListItem.attr("id","")

	$itemListItem.find(".tocListName").text(listInfo.name)
		
	var $itemListLink = $itemListItem.find("a")
	
	$itemListLink.click(function(e) {
		e.preventDefault()
		console.log("Item list TOC item clicked: list id = " + listInfo.listID)
		$itemListLink.blur()
		
		$('#tocWrapper').find("li").removeClass("active")
		$itemListItem.addClass("active")
		
		if(tocConfig.itemListClickedCallback !== undefined) {
			tocConfig.itemListClickedCallback(listInfo.listID,$itemListItem)
		}
		setDefaultTOCItem(tocConfig.databaseID,listInfo.listID)
		
	})
	
	var listCountParams = {
		databaseID: tocConfig.databaseID,
		preFilterRules: listInfo.properties.preFilterRules,
	}
	
	function refreshListCount() {
		jsonAPIRequest("recordRead/getFilteredRecordCount",listCountParams,function(listCount) {
			var $listCount = $itemListItem.find(".badge")
			if (listCount > 0) {
				$listCount.text(listCount)
			} else {
				$listCount.hide()
			}
		})		
	}
	
	initTableOfContentsRefreshPollingLoop(refreshListCount)
	
	// Load this list as the default list if it is set as a default.
	if(getDefaultTOCItem(tocConfig.databaseID)===listInfo.listID && 
		tocConfig.itemListClickedCallback !== undefined) {
			tocConfig.itemListClickedCallback(listInfo.listID,$itemListItem)			
	}
	
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
			tocConfig.newItemLinkClickedCallback(linkInfo.linkID,$formLinkListItem)
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
	
	function getTOCInfo(infoCallback) {
		var infoRemaining = 2
		var tocInfo = {}
		
		function processOneTOCInfo() {
			infoRemaining--
			if(infoRemaining <= 0) {
				infoCallback(tocInfo)
			}
		}
		
		var getDBInfoParams = { databaseID: tocConfig.databaseID }
		jsonAPIRequest("dashboard/getUserDashboardList",getDBInfoParams,function(dashboardsInfo) {
			tocInfo.dashboardsInfo = dashboardsInfo
			processOneTOCInfo()
		})
		
		jsonAPIRequest("itemList/getUserItemListList",getDBInfoParams,function(listsInfo) {
			tocInfo.listsInfo = listsInfo
			processOneTOCInfo()
		})
		
	}
	
	getTOCInfo(function(tocInfo) {
		
		// Set the default TOC item if it is undefined
		var defaultTOCItem = getDefaultTOCItem(tocConfig.databaseID)
		if(defaultTOCItem === null || defaultTOCItem === undefined) {
			if (tocInfo.listsInfo.length > 0) {
				var firstListInfo = tocInfo.listsInfo[0]
				setDefaultTOCItem(tocConfig.databaseID,firstListInfo.listID)
			} else if (tocInfo.dashboardsInfo.length > 0) {
				var firstDashboardInfo = tocInfo.dashboardsInfo[0]
				setDefaultTOCItem(tocConfig.databaseID,firstDashboardInfo.dashboardID)
			}
		}
		
		
		var listsInfo = tocInfo.listsInfo
		$('#tocListList').empty()
		for(var listInfoIndex = 0; listInfoIndex < listsInfo.length; listInfoIndex++) {
			var listInfo = listsInfo[listInfoIndex]
			addItemListLinkToTOCList(tocConfig,listInfo)
		}
		if (listsInfo.length === 0) {
			$('#tocListsContainer').hide()
		}
		
		var dashboardsInfo = tocInfo.dashboardsInfo
		$('#tocDashboardList').empty()
		for (var dashboardInfoIndex = 0; dashboardInfoIndex < dashboardsInfo.length; dashboardInfoIndex++) {
			var dashboardInfo = dashboardsInfo[dashboardInfoIndex]
			addDashboardLinkToTOCList(tocConfig,dashboardInfo)
		}
		if(dashboardsInfo.length === 0) {
			$('#tocDashboardsContainer').hide()
		}
		
	})
	
	
}