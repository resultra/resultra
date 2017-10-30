
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#generalAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(generalAdminPageContext.databaseID,"settingsTOCGeneral")
			
	initAdminGeneralProperties(generalAdminPageContext.databaseID)
	
})