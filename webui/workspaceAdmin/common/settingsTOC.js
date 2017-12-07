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