$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#tableAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(tableAdminContext.databaseID,"settingsTOCTables")
			
	initAdminTableListSettings(tableAdminContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/tables/"+tableAdminContext.databaseID,"Table Views",
		tableAdminContext.isSingleUserWorkspace)
	
	
})