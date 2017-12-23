$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formLinkAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(formLinkAdminContext.databaseID,"settingsTOCFormLinks")
			
	initAdminFormLinkSettings(formLinkAdminContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/formlink/"+formLinkAdminContext.databaseID,"New Item Links")
	
	
})