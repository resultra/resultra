
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#generalAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(generalAdminPageContext.databaseID)
			
	initAdminGeneralProperties(generalAdminPageContext.databaseID)
	
})