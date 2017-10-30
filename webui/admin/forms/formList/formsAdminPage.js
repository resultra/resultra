$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(formsAdminContext.databaseID,"settingsTOCForms")
			
	initAdminFormSettings(formsAdminContext.databaseID)
	
})