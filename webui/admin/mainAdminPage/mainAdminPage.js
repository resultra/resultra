
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

	
	// Call-back for dynamically setting the settings page, depending on the link pressed in the settings TOC.
	function setSettingsPage(linkID) {
		
		if (linkID === "general") {
			setGeneralSettingsPage()
		} else if(linkID == "forms") {
			setFormsSettingsPage()
		} else if(linkID == "formLinks") {
			setFormLinksSettingsPage()
		} else {
			setGeneralSettingsPage()			
		}
		
	}
	
	
	
	initAdminSettingsPageLayout($('#mainAdminPage'))
		
	initAdminPageHeader(mainAdminPageContext.isSingleUserWorkspace)
	
	initAdminSettingsTOC(mainAdminPageContext.databaseID,
		"settingsTOCGeneral",mainAdminPageContext.isSingleUserWorkspace,setSettingsPage)
		
		
	setGeneralSettingsPage()
		
				
})