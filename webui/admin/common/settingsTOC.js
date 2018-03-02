function initAdminSettingsTOC(databaseID, activeID,isSingleUserWorkspace,changeLinkCallback) {
	
	
	var $settingsTOC = $('#settingsTOC')
	$settingsTOC.find("li").removeClass("active")
	
	var $activeItem = $('#' + activeID)
	$activeItem.addClass("active")
	

	function initSettingsLinkListItem(listItemSelector, linkID) {
		
		var $listItem = $(listItemSelector)
		
		var $link = $listItem.find("a")
	
		$link.click(function(e) {
			e.preventDefault()

			$link.blur()
		
			$settingsTOC.find("li").removeClass("active")
			$listItem.addClass("active")
			
			changeLinkCallback(linkID)
			/*
		
			if(tocConfig.itemListClickedCallback !== undefined) {
				tocConfig.itemListClickedCallback(listInfo.listID,$itemListItem)
			}
			setDefaultTOCItem(tocConfig.databaseID,listInfo.listID)
		
			*/
		})
		
	}
	
	initSettingsLinkListItem("#settingsTOCGeneral","general")
	initSettingsLinkListItem("#settingsTOCForms","forms")
	initSettingsLinkListItem("#settingsTOCFormLinks","formLinks")
	initSettingsLinkListItem("#settingsTOCTables","tables")
	initSettingsLinkListItem("#settingsTOCLists","lists")

/*

	var fieldsLink = '/admin/fields/' + databaseID
	$('#settingsTOCFields').find("a").attr("href",fieldsLink)
	
	var formsLink = "/admin/forms/" + databaseID
	$('#settingsTOCForms').find("a").attr("href",formsLink)
	
	var valueListsLink = "/admin/valuelists/" + databaseID
	$('#settingsTOCValueLists').find("a").attr("href",valueListsLink)
	
	var dashboardLink = "/admin/dashboards/" + databaseID
	$('#settingsTOCDashboards').find("a").attr("href",dashboardLink)
	
	var globalLink = "/admin/globals/" + databaseID
	$('#settingsTOCGlobals').find("a").attr("href",globalLink)
	
	if (!isSingleUserWorkspace) {
		var userLink = "/admin/collaborators/" + databaseID
		$('#settingsTOCUsers').find("a").attr("href",userLink)		
	}

	var roleLink = "/admin/roles/" + databaseID
	$('#settingsTOCRoles').find("a").attr("href",roleLink)
	
	var alertLink = "/admin/alerts/" + databaseID
	$('#settingsTOCAlerts').find("a").attr("href",alertLink)
	
*/
}