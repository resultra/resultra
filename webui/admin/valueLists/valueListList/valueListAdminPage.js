	
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#valueListAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(valueListAdminContext.databaseID,"settingsTOCValueLists")
			
	initAdminValueListListSettings(valueListAdminContext.databaseID)
	
	
	appendPageSpecificBreadcrumbHeader("/admin/valuelists/"+valueListAdminContext.databaseID,"Value Lists")
	
})