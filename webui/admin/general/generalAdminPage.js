
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#generalAdminPage'))	
	initAdminPageHeader(generalAdminPageContext.isSingleUserWorkspace)
	initAdminSettingsTOC(generalAdminPageContext.databaseID,"settingsTOCGeneral",generalAdminPageContext.isSingleUserWorkspace)
			
	initAdminGeneralProperties(generalAdminPageContext.databaseID)
	
})