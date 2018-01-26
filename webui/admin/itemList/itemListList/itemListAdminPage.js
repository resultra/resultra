$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#itemListAdminPage'))	
	initAdminPageHeader(itemListAdminContext.isSingleUserWorkspace)
	initAdminSettingsTOC(itemListAdminContext.databaseID,"settingsTOCLists")
			
	initAdminListSettings(itemListAdminContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/lists/"+itemListAdminContext.databaseID,"Item Lists",
				itemListAdminContext.isSingleUserWorkspace)
	
})