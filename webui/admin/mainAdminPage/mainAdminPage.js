
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
	
	
	
	function setFieldListSettingsPage() {
		const contentURL = '/admin/fields/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initFieldsSettingsPageContent(mainAdminPageContext)
		});
		
	}

	function setValueListsSettingsPage() {
		const contentURL = '/admin/valuelists/' + mainAdminPageContext.databaseID
	
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        $('#contentPageSection').html(pageContentData);
				initValueListAdminSettingsPageContent(mainAdminPageContext)
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