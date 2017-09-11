	
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#valueListAdminPage'))	
	initUserDropdownMenu()
	initAlertHeader(valueListAdminContext.databaseID)
	initAdminSettingsTOC(valueListAdminContext.databaseID)
			
	initAdminValueListListSettings(valueListAdminContext.databaseID)
	
})