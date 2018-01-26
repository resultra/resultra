
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#userAdminPage'))	
	initAdminPageHeader(false)
	initAdminSettingsTOC(userAdminPageContext.databaseID,"settingsTOCUsers")
	

	initUserListSettings(userAdminPageContext.databaseID)

	appendPageSpecificBreadcrumbHeader("/admin/collaborators/"+userAdminPageContext.databaseID,"Collaborators")

				
})