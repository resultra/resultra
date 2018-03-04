
$(document).ready(function() {
	
	function setGeneralSettingsPage() {
		const contentURL = '/admin/general/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initGeneralAdminPageContent(mainAdminPageContext.databaseID)
		});
		
	}
	
	function setFormsSettingsPage() {
		const contentURL = '/admin/forms/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initFormListAdminPage(mainAdminPageContext)
		});
		
	}

	function setFormLinksSettingsPage() {
		const contentURL = '/admin/formlink/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initFormLinkSettingsPage(mainAdminPageContext)
		});
		
	}
	registerPageContentLoader("formLinks",'/admin/formlink/' + mainAdminPageContext.databaseID,function() {
		initFormLinkSettingsPage(mainAdminPageContext)
	})


	function setTableListSettingsPage() {
		const contentURL = '/admin/tables/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initTableListAdminPageContent(mainAdminPageContext)
		});
		
	}
	
	function setItemListSettingsPage() {
		const contentURL = '/admin/lists/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initItemListAdminSettingsPage(mainAdminPageContext)
		});
		
	}
	
	registerPageContentLoader("lists",'/admin/lists/' + mainAdminPageContext.databaseID,function() {
		initItemListAdminSettingsPage(mainAdminPageContext)
	})
	
	
	function setFieldListSettingsPage() {
		const contentURL = '/admin/fields/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initFieldsSettingsPageContent(mainAdminPageContext)
		});
		
	}
	registerPageContentLoader("fields", '/admin/fields/' + mainAdminPageContext.databaseID,function() {
		initFieldsSettingsPageContent(mainAdminPageContext)
	})



	function setValueListsSettingsPage() {
		const contentURL = '/admin/valuelists/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initValueListAdminSettingsPageContent(mainAdminPageContext)
		});
		
	}
	
	registerPageContentLoader("valueLists",'/admin/valuelists/' + mainAdminPageContext.databaseID,function() {
		initValueListAdminSettingsPageContent(mainAdminPageContext)
	})
	
	function setDashboardsSettingsPage() {
		const contentURL = '/admin/dashboards/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initDashboardsAdminSettingsPageContent(mainAdminPageContext)
		});
		
	}
	
	
	function setAlertsSettingsPage() {
		const contentURL = '/admin/alerts/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initAlertListAdminSettingsPageContent(mainAdminPageContext)
		});
		
	}

	registerPageContentLoader("alerts",'/admin/alerts/' + mainAdminPageContext.databaseID,function() {
		initAlertListAdminSettingsPageContent(mainAdminPageContext)
	})


	function setRolesSettingsPage() {
		const contentURL = '/admin/roles/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initUserRoleAdminSettingsPageContent(mainAdminPageContext)
		});
		
	}
	
	function setCollaboratorsSettingsPage() {
		const contentURL = '/admin/collaborators/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initCollaboratorsSettingsPageContent(mainAdminPageContext)
		});
		
	}
	
	
	function setGlobalsSettingsPage() {
		const contentURL = '/admin/globals/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initGlobalSettingsPageContent(mainAdminPageContext)
		});
		
	}
	

	// Call-back for dynamically setting the settings page, depending on the link pressed in the settings TOC.
	function setSettingsPage(linkID) {
		
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
	
	initAdminSettingsPageLayout($('#mainAdminPage'))
	
	function getLinkIDAnchorName() {
		var linkID = window.location.hash.substr(1);
		if (linkID === null || linkID.length === 0) {
			return "general"
		}
		return linkID
	}
		
	initAdminPageHeader(mainAdminPageContext.isSingleUserWorkspace)
	
	const currAnchorLinkID = getLinkIDAnchorName()
	
	initAdminSettingsTOC(mainAdminPageContext.databaseID,
		currAnchorLinkID,mainAdminPageContext.isSingleUserWorkspace,setSettingsPage)
			
	setSettingsPage(currAnchorLinkID)
		
				
})