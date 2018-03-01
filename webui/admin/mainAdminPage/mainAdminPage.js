
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#mainAdminPage'))	
	initAdminPageHeader(mainAdminPageContext.isSingleUserWorkspace)
	initAdminSettingsTOC(mainAdminPageContext.databaseID,
		"settingsTOCGeneral",mainAdminPageContext.isSingleUserWorkspace)
		
		
		const contentURL = '/admin/general/' + mainAdminPageContext.databaseID
		
		jQuery.get(contentURL, function(pageContentData) { // Perform AJAX GET request
		        jQuery('#contentPageSection').html(pageContentData);
				initGeneralAdminPageContent(mainAdminPageContext.databaseID)
		});
				
})