function initTableOfContentsRefreshPollingLoop(refreshCallback) {
    var userActivityTimer;
	
	var userIsActive = true
	
	$(window).mousemove(resetUserActivityTimer)
	$(window).click(resetUserActivityTimer)
	$(window).mousedown(resetUserActivityTimer)
	$(window).keypress(resetUserActivityTimer)
	$(window).scroll(resetUserActivityTimer)

    function setUserInactive() {
		userIsActive = false
    }

    function resetUserActivityTimer() {
        clearTimeout(userActivityTimer);
		userIsActive = true
		// If the timer completes before user activity is seen, then 
		// disable the inactivity timer.
        userActivityTimer = setTimeout(setUserInactive, 10000);  // time is in milliseconds
    }
	
	function refresh() {
		if(userIsActive) {
			refreshCallback()
		} 
		setTimeout(refresh,5000)
	}
	refresh()
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
	})
	
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