
function initTrackerAdminPageContent(pageContext,trackerInfo) {
	
	function setGeneralSettingsPage() {
		const contentURL = '/admin/general/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initGeneralAdminPageContent(pageContext.databaseID)
		});
		
	}
	
	function setFormsSettingsPage() {
		const contentURL = '/admin/forms/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initFormListAdminPage(pageContext)
		});
		
	}
	registerPageContentLoader("forms",'/admin/forms/' + pageContext.databaseID,function() {
		initFormListAdminPage(pageContext)
	})

	function setFormLinksSettingsPage() {
		const contentURL = '/admin/formlink/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initFormLinkSettingsPage(pageContext)
		});
		
	}
	registerPageContentLoader("formLinks",'/admin/formlink/' + pageContext.databaseID,function() {
		initFormLinkSettingsPage(pageContext)
	})


	function setTableListSettingsPage() {
		const contentURL = '/admin/tables/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initTableListAdminPageContent(pageContext)
		});
		
	}
	registerPageContentLoader("tables",'/admin/tables/' + pageContext.databaseID,function() {
		initTableListAdminPageContent(pageContext)
	})


	
	function setItemListSettingsPage() {
		const contentURL = '/admin/lists/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initItemListAdminSettingsPage(pageContext)
		});
		
	}
	registerPageContentLoader("lists",'/admin/lists/' + pageContext.databaseID,function() {
		initItemListAdminSettingsPage(pageContext)
	})
	
	
	function setFieldListSettingsPage() {
		const contentURL = '/admin/fields/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initFieldsSettingsPageContent(pageContext)
		});
		
	}
	registerPageContentLoader("fields", '/admin/fields/' + pageContext.databaseID,function() {
		initFieldsSettingsPageContent(pageContext)
	})



	function setValueListsSettingsPage() {
		const contentURL = '/admin/valuelists/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initValueListAdminSettingsPageContent(pageContext)
		});
		
	}
	registerPageContentLoader("valueLists",'/admin/valuelists/' + pageContext.databaseID,function() {
		initValueListAdminSettingsPageContent(pageContext)
	})
	
	function setDashboardsSettingsPage() {
		const contentURL = '/admin/dashboards/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initDashboardsAdminSettingsPageContent(pageContext)
		});
		
	}
	registerPageContentLoader("dashboards",'/admin/dashboards/' + pageContext.databaseID,function() {
		initDashboardsAdminSettingsPageContent(pageContext)
	})
	
	
	function setAlertsSettingsPage() {
		const contentURL = '/admin/alerts/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initAlertListAdminSettingsPageContent(pageContext)
		});
		
	}
	registerPageContentLoader("alerts",'/admin/alerts/' + pageContext.databaseID,function() {
		initAlertListAdminSettingsPageContent(pageContext)
	})


	function setRolesSettingsPage() {
		const contentURL = '/admin/roles/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initUserRoleAdminSettingsPageContent(pageContext)
		});
		
	}
	registerPageContentLoader("roles",'/admin/roles/' + pageContext.databaseID,function() {
		initUserRoleAdminSettingsPageContent(pageContext)
	})
	
	
	function setCollaboratorsSettingsPage() {
		const contentURL = '/admin/collaborators/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initCollaboratorsSettingsPageContent(pageContext)
		});
		
	}
	registerPageContentLoader("collaborators",'/admin/collaborators/' + pageContext.databaseID,function() {
		initCollaboratorsSettingsPageContent(pageContext)
	})
	
	
	function setGlobalsSettingsPage() {
		const contentURL = '/admin/globals/' + pageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initGlobalSettingsPageContent(pageContext)
		});
		
	}
	

	// Call-back for dynamically setting the settings page, depending on the link pressed in the settings TOC.
	function setSettingsPage(linkID) {
		
		resetSettingsPageLayoutForStandardSettingsPages()
		
		if (linkID === "general") {
			setGeneralSettingsPage()
		} else if(linkID == "forms") {
			setFormsSettingsPage()
		} else if(linkID == "formLinks") {
			setFormLinksSettingsPage()
		} else if(linkID == "tables") {
			setTableListSettingsPage()
		} else if(linkID == "lists") {
			setItemListSettingsPage()
		} else if(linkID == "fields") {
			setFieldListSettingsPage()
		} else if(linkID == "valueLists") {
			setValueListsSettingsPage()
		} else if(linkID == "dashboards") {
			setDashboardsSettingsPage()
		} else if(linkID == "alerts") {
			setAlertsSettingsPage()
		} else if(linkID == "roles") {
			setRolesSettingsPage()
		} else if(linkID == "collaborators") {
			setCollaboratorsSettingsPage()
		} else if(linkID == "globals") {
			setGlobalsSettingsPage()
		} else {
			setGeneralSettingsPage()			
		}
		
		// Update the location in the browser. This is needed
		// to support the browser's back button in the case the user
		// navigates further down into the settings. Similarly if the 
		// user presses the refresh button, the most recent page content will
		// also be shown.
		window.location = window.location.origin + window.location.pathname + "#" + linkID
		
	}
	
	function getLinkIDAnchorName() {
		var linkID = window.location.hash.substr(1);
		if (linkID === null || linkID.length === 0) {
			return "general"
		}
		return linkID
	}
		
	initAdminPageHeader(pageContext.isSingleUserWorkspace)
	
	const currAnchorLinkID = getLinkIDAnchorName()
	
	initAdminSettingsTOC(pageContext.databaseID,
		currAnchorLinkID,setSettingsPage)
			
	setSettingsPage(currAnchorLinkID)
		
	resetWorkspaceBreadcrumbHeader()
	appendMainWindowContentSpecificBreadcrumbHeader(trackerInfo.databaseName,function() {
			navigateToTracker(pageContext,trackerInfo)
	})
	appendMainWindowContentSpecificBreadcrumbHeader("Settings",function() {
			navigateToAdminSettingsPageContent(pageContext,trackerInfo)
	})
	
				
}


function navigateToAdminSettingsPageContent(pageContext,trackerInfo) {
	
	theMainWindowLayout.disableRHSSidebar()	
	
	var contentConfig = {
		mainContentURL: '/admin/mainAdminPage/mainPageContent/'+pageContext.databaseID,
		lhsSidebarContentURL: "/admin/common/settingsTOC",
		offPageContentURL: "/admin/mainAdminPage/offPageContent"
	}

	setMainWindowPageContent(contentConfig,function() {
		var trackerPageContext = {
			databaseID: trackerInfo.databaseID,
			isSingleUserWorkspace:pageContext.isSingleUserWorkspace
		} 
		initTrackerAdminPageContent(trackerPageContext,trackerInfo)
	})				
}
