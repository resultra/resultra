$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formLinkAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(formLinkAdminContext.databaseID,"settingsTOCFormLinks")
			
	initAdminFormLinkSettings(formLinkAdminContext.databaseID)
	
})