$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formAdminPage'))	
	
	initAdminPageHeader(formsAdminContext.isSingleUserWorkspace)
	
	initAdminSettingsTOC(formsAdminContext.databaseID,"settingsTOCForms")
			
	initAdminFormSettings(formsAdminContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/forms/"+formsAdminContext.databaseID,"Forms",formsAdminContext.isSingleUserWorkspace)
	
})