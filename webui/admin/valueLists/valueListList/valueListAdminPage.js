	
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#valueListAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(valueListAdminContext.databaseID)
			
	initAdminValueListListSettings(valueListAdminContext.databaseID)
	
})