	
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#valueListAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(valueListAdminContext.databaseID,"settingsTOCValueLists",valueListAdminContext.isSingleUserWorkspace)
			
	initAdminValueListListSettings(valueListAdminContext.databaseID)
	
	
	appendPageSpecificBreadcrumbHeader("/admin/valuelists/"+valueListAdminContext.databaseID,"Value Lists")
	
})