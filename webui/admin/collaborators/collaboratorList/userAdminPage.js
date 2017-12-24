
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#userAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(userAdminPageContext.databaseID,"settingsTOCUsers")
	

	initUserListSettings(userAdminPageContext.databaseID)

	appendPageSpecificBreadcrumbHeader("/admin/collaborators/"+userAdminPageContext.databaseID,"Collaborators")

				
})