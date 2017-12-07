$(document).ready(function() {	
	
	initUserDropdownMenu()
	initHelpDropdownMenu()
	
	initWorkspaceNameProperty(workspaceAdminContext.workspaceName)
	
	var infoParams = {}
	jsonAPIRequest("workspace/getInfo",infoParams,function(workspaceInfo) {
		initWorkspacePermissionSettings(workspaceInfo)
		initUserRegistrationProps(workspaceInfo)
	})
	
	
}); // document ready
