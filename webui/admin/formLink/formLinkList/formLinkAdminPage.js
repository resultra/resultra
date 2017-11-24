$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formLinkAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(formLinkAdminContext.databaseID,"settingsTOCFormLinks")
			
	initAdminFormLinkSettings(formLinkAdminContext.databaseID)
	
})