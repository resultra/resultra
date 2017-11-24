
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#generalAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(generalAdminPageContext.databaseID,"settingsTOCGeneral")
			
	initAdminGeneralProperties(generalAdminPageContext.databaseID)
	
})