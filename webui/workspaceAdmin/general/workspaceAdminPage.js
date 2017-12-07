$(document).ready(function() {	
	
	
	initWorkspaceAdminSettingsPageLayout($('#generalAdminPage'))	
	
	initWorkspaceAdminPageHeader()
	
	initWorkspaceAdminSettingsTOC("settingsTOCGeneral")
	
	
	initWorkspaceNameProperty(workspaceAdminContext.workspaceName)
	
	var infoParams = {}
	jsonAPIRequest("workspace/getInfo",infoParams,function(workspaceInfo) {
		initWorkspacePermissionSettings(workspaceInfo)
	})
	
	
}); // document ready
