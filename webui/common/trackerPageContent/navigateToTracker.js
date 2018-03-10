function navigateToTracker(trackerInfo) {
	
	const databaseID = trackerInfo.databaseID
	
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
			$tocItem.addClass("active")
			
		})
	}
	
	function dashboardClicked(dashboardID,$tocItem) {
		
		var contentConfig = {
			mainContentURL: "/dashboard/view/contentLayout",
			rhsSidebarContentURL: "/dashboard/view/sidebarLayout"
		}
		setMainWindowPageContent(contentConfig,function() {
			var contentLayout = new DashboardContentLayout()
			theMainWindowLayout.showRHSSidebar()
			loadDashboardView(contentLayout,databaseID, dashboardID)	
			$tocItem.addClass("active")		
		})
		
	}
	
	function newItemClicked(linkID,$tocItem) {
		console.log("Main window: new item clicked: " + linkID)
		
		theMainWindowLayout.disableRHSSidebar()
		
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
		
	}
	
	
	
	setLHSSidebarContent("/common/trackerPageContent/toc",function() {
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
	
	var headerButtonsContentURL = "/common/trackerPageContent/headerButtons/" + databaseID
	setMainWindowHeaderButtonsContent(headerButtonsContentURL,function() {
		
		function seeAllAlertsClicked() {	
			var contentConfig = {
				mainContentURL: "/alertListView/contentLayout"
			}
			setMainWindowPageContent(contentConfig,function() {
				var contentLayout = new AlertListContentLayout()
				theMainWindowLayout.disableRHSSidebar()
				initAlertNotificationList(contentLayout,databaseID)	
			})	
		}
		initAlertHeader(databaseID,seeAllAlertsClicked)
		
		function loadSettingsPageContent() {
			theMainWindowLayout.disableRHSSidebar()	
			
			var contentConfig = {
				mainContentURL: '/admin/'+databaseID,
				lhsSidebarContentURL: "/admin/common/settingsTOC",
				offPageContentURL: "/admin/offPageContent"
			}
		
			setMainWindowPageContent(contentConfig,function() {
				var pageContext = {
					databaseID: databaseID
				} // TODO - pass as parameter
				initTrackerAdminPageContent(pageContext)
			})				
		}
		var $adminButton = $("#adminSettingsHeaderButton")
		$adminButton.click(function(e) {
			e.preventDefault()
			$adminButton.blur()
			loadSettingsPageContent()
		})
		
		
	})
	
	resetWorkspaceBreadcrumbHeader()
	appendMainWindowContentSpecificBreadcrumbHeader(trackerInfo.databaseName,function() {
		navigateToTracker(trackerInfo)
	})
		
}
