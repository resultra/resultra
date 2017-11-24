$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(formsAdminContext.databaseID,"settingsTOCForms")
			
	initAdminFormSettings(formsAdminContext.databaseID)
	
})