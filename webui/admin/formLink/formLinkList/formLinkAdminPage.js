$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formLinkAdminPage'))	
	initAdminPageHeader(formLinkAdminContext.isSingleUserWorkspace)
	initAdminSettingsTOC(formLinkAdminContext.databaseID,"settingsTOCFormLinks",formLinkAdminContext.isSingleUserWorkspace)
			
	initAdminFormLinkSettings(formLinkAdminContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/formlink/"+formLinkAdminContext.databaseID,"New Item Links")
	
	
})