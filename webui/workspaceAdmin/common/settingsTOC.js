// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initWorkspaceAdminSettingsTOC(activeID) {
	
	
	var $settingsTOC = $('#settingsTOC')
	$settingsTOC.find("li").removeClass("active")
	var $activeItem = $('#' + activeID)
	$activeItem.addClass("active")
	
	var generalLink = '/workspace-admin'
	$('#settingsTOCGeneral').find("a").attr("href",generalLink)

	var usersLink = '/workspace-admin/users'
	$('#settingsTOCUsers').find("a").attr("href",usersLink)
		
	
}