$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#roleAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(roleAdminContext.databaseID,"settingsTOCRoles",
		roleAdminContext.isSingleUserWorkspace)
			
	initUserRoleSettings(roleAdminContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/roles/"+roleAdminContext.databaseID,"Roles")
	
	
})