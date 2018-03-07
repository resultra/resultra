
function navigateToTracker(databaseID) {
	
	function itemListClicked(listID,$tocItem) {
		
		var contentConfig = {
			mainContentURL: "/itemList/contentLayout",
			rhsSidebarContentURL: "/itemList/propertySidebarContent",
			offPageContentURL: "/itemList/offPageContent"
		}
		
		setMainWindowPageContent(contentConfig,function() {
			theMainWindowLayout.showRHSSidebar()
			var contentLayout = new ItemListContentLayout()
		
			loadItemListView(contentLayout,databaseID,listID)
//			mainWinLayout.clearSidebarNavigationSelection()
			$tocItem.addClass("active")
			
		})
				
		
		/*
		
		function loadView() {
			loadItemListView(mainWinLayout,mainWindowContext.databaseID,listID)
			mainWinLayout.clearSidebarNavigationSelection()
			$tocItem.addClass("active")
		}
		console.log("Main window: item list clicked: " + listID)
		loadView()
		loadLastViewCallback = loadView
		*/
		
	}
	
	function dashboardClicked(dashboardID,$tocItem) {
		
		var contentConfig = {
			mainContentURL: "/dashboard/view/contentLayout"
		}
		setMainWindowPageContent(contentConfig,function() {
			var contentLayout = new DashboardContentLayout()
			loadDashboardView(contentLayout,databaseID, dashboardID)	
			$tocItem.addClass("active")		
		})
		
		/*
		function loadView() {
			loadDashboardView(mainWinLayout,mainWindowContext.databaseID, dashboardID)	
			mainWinLayout.clearSidebarNavigationSelection()
			$tocItem.addClass("active")
		}
		loadView()
		loadLastViewCallback = loadView */
	}
	
	function newItemClicked(linkID,$tocItem) {
		console.log("Main window: new item clicked: " + linkID)
		
		var contentConfig = {
			mainContentURL: "/itemView/newItemContentLayout",
			offPageContentURL: "/itemView/newItemOffPageContent"
		}
		setMainWindowPageContent(contentConfig,function() {
			var newItemLayout = new NewItemContentLayout()
			function loadLastViewCallback() {
				// TBD
			}
		
			var newItemParams = {
				pageLayout: newItemLayout,
				databaseID: databaseID,
				formLinkID: linkID,
				loadLastViewCallback: loadLastViewCallback
			}
			loadNewItemView(newItemParams)
			$tocItem.addClass("active")
			
		})
		
		
		/*
		
		
		mainWinLayout.clearSidebarNavigationSelection()

		*/
	}
	
	function seeAllAlertsClicked() {
		/*
		function loadView() {
			mainWinLayout.clearSidebarNavigationSelection()
			initAlertNotificationList(mainWinLayout,mainWindowContext.databaseID)			
		}
		loadView()
		loadLastViewCallback = loadView
		*/
	}
	
	
	setLHSSidebarContent("/common/trackerTOC/toc",function() {
		var tocConfig = {
			databaseID: databaseID,
			newItemFormButtonFunc: openSubmitFormDialog,
			itemListClickedCallback: itemListClicked,
			dashboardClickedCallback: dashboardClicked,
			newItemLinkClickedCallback: newItemClicked
		}	
		initDatabaseTOC(tocConfig)
		theMainWindowLayout.showLHSSidebar()
		theMainWindowLayout.openLHSSidebar()
		
	})
		
}


function addTrackerListItem(trackerInfo) {

	var $trackerList = $("#myTrackerList")

	var $listItem = $('#trackerListItemTemplate').clone()
	$listItem.attr("id","")

	var $nameLabel = $listItem.find(".nameLabel")
	$nameLabel.text(trackerInfo.databaseName)
	
	
	// Only enable the link to open the tracker if the tracker is  active.
	if(trackerInfo.isActive) {
		
		// var openTrackerLink = '/main/' + trackerInfo.databaseID
		
		// $nameLabel.attr('href',openTrackerLink)	
		$nameLabel.click(function() {
		 	   console.log("tracker link clicked")
			navigateToTracker(trackerInfo.databaseID)
		})
		
		
		
	} else {
		$nameLabel.addClass("disabledTrackerLink")
	}

	var $settingsLink = $listItem.find(".adminEditPropsButton")

	if (trackerInfo.isAdmin) {
		var editPropsLink = '/admin/' + trackerInfo.databaseID
		$settingsLink.attr('href',editPropsLink)
		$settingsLink.tooltip()
	} else {
		$settingsLink.hide()
	}

	$trackerList.append($listItem)

}



function initTrackerList() {
	
	var $trackerList = $("#myTrackerList")

		
	function reloadTrackerList(includeInactive) {
		var getDBListParams = {
			includeInactive:includeInactive
		}
		jsonAPIRequest("database/getList",getDBListParams,function(trackerList) {
			$trackerList.empty()
			for (var trackerIndex=0; trackerIndex<trackerList.length; trackerIndex++) {	
				var trackerInfo = trackerList[trackerIndex]
				addTrackerListItem(trackerInfo)
			}
		})
	}
	reloadTrackerList(false)
	
	
	initButtonClickHandler('#newTrackerButton',function() {
		console.log("New form button clicked")
		openNewTrackerDialog()
	})
	
	initCheckboxChangeHandler('#showInactiveTrackers', false, function(includeInactive) {
		reloadTrackerList(includeInactive)
	})

	
}