	
$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#valueListAdminPage'))	
	initAdminPageHeader()
	initAdminSettingsTOC(valueListAdminContext.databaseID,"settingsTOCValueLists")
			
	initAdminValueListListSettings(valueListAdminContext.databaseID)
	
})