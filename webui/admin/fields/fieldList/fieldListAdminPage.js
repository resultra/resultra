$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#fieldAdminPage'))	
	initAdminPageHeader(fieldListAdminContext.isSingleUserWorkspace)
	initAdminSettingsTOC(fieldListAdminContext.databaseID,"settingsTOCFields")
		
	initAdminFieldSettings(fieldListAdminContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/fields/"+fieldListAdminContext.databaseID,"Fields",
		fieldListAdminContext.isSingleUserWorkspace)
	
})