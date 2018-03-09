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