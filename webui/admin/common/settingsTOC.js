// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initAdminSettingsTOC(databaseID, activeID,changeLinkCallback) {
	
	
	var $settingsTOC = $('#settingsTOC')
	$settingsTOC.find("li").removeClass("active")
	
	
	function linkListItemSelector(linkID) {
		return '#settingsTOC_' + linkID
	}
	
	const activeItemSelector = linkListItemSelector(activeID)
	var $activeItem = $(activeItemSelector)
	$activeItem.addClass("active")
	

	function initSettingsLinkListItem(linkID) {
				
		var $listItem = $(linkListItemSelector(linkID))
		
		var $link = $listItem.find("a")
		if($link.length) {
			$link.click(function(e) {
				e.preventDefault()

				$link.blur()
		
				$settingsTOC.find("li").removeClass("active")
				$listItem.addClass("active")
			
				changeLinkCallback(linkID)
			})			
		}
	
		
	}
	
	initSettingsLinkListItem("general")
	initSettingsLinkListItem("forms")
	initSettingsLinkListItem("formLinks")
	initSettingsLinkListItem("tables")
	initSettingsLinkListItem("lists")
	initSettingsLinkListItem("fields")
	initSettingsLinkListItem("valueLists")
	initSettingsLinkListItem("dashboards")
	initSettingsLinkListItem("alerts")
	initSettingsLinkListItem("roles")
	initSettingsLinkListItem("globals")
	initSettingsLinkListItem("collaborators")	
	
}