$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#roleAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(roleAdminContext.databaseID,"settingsTOCRoles")
			
	initUserRoleSettings(roleAdminContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/roles/"+roleAdminContext.databaseID,"Roles")
	
	
})