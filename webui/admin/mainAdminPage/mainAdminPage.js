
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#mainAdminPage'))	
	initAdminPageHeader(mainAdminPageContext.isSingleUserWorkspace)
	initAdminSettingsTOC(mainAdminPageContext.databaseID,
		"settingsTOCGeneral",mainAdminPageContext.isSingleUserWorkspace)
				
})