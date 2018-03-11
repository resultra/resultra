function navigateToTracker(pageContext,trackerInfo) {
	
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
			
			// Listen for events to view a specific record/item in a particular form. This happens in response to
			// clicks to a form button deeper down in the DOM.
			$('#listViewContentLayout').on(viewFormInViewportEventName,function(e,params) {
				e.stopPropagation()
				console.log("Got formButton load form event: " + JSON.stringify(params))		
				loadExistingItemViewPageContent(params)
			})
			
			
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
		
		var $adminButton = $("#adminSettingsHeaderButton")
		$adminButton.click(function(e) {
			e.preventDefault()
			$adminButton.blur()
			navigateToAdminSettingsPageContent(pageContext,trackerInfo)
		})
		
		
	})
	
	resetWorkspaceBreadcrumbHeader()
	appendMainWindowContentSpecificBreadcrumbHeader(trackerInfo.databaseName,function() {
		navigateToTracker(pageContext,trackerInfo)
	})
		
}
