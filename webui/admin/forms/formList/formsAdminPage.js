$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#formAdminPage'))	
	
	initAdminPageHeader()
	
	initAdminSettingsTOC(formsAdminContext.databaseID,"settingsTOCForms")
			
	initAdminFormSettings(formsAdminContext.databaseID)
	
	appendPageSpecificBreadcrumbHeader("/admin/forms/"+formsAdminContext.databaseID,"Forms",formsAdminContext.isSingleUserWorkspace)
	
})