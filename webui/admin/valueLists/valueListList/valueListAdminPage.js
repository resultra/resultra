	
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#valueListAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(valueListAdminContext.databaseID,"settingsTOCValueLists")
			
	initAdminValueListListSettings(valueListAdminContext.databaseID)
	
})