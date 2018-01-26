
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#generalAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(generalAdminPageContext.databaseID,"settingsTOCGeneral",generalAdminPageContext.isSingleUserWorkspace)
			
	initAdminGeneralProperties(generalAdminPageContext.databaseID)
	
})