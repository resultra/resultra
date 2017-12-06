$(document).ready(function() {	
	
	initUserDropdownMenu()
	initHelpDropdownMenu()
	
	initWorkspaceNameProperty(workspaceAdminContext.workspaceName)
	
	var infoParams = {}
	jsonAPIRequest("workspace/getInfo",infoParams,function(workspaceInfo) {
		initWorkspacePermissionSettings(workspaceInfo)
	})
	
	
}); // document ready
