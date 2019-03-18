// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
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
