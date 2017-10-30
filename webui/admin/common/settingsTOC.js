function initAdminSettingsTOC(databaseID, activeID) {
	
	
	var $settingsTOC = $('#settingsTOC')
	$settingsTOC.find("li").removeClass("active")
	var $activeItem = $('#' + activeID)
	$activeItem.addClass("active")
	

	var generalLink = '/admin/general/' + databaseID
	$('#settingsTOCGeneral').find("a").attr("href",generalLink)

	var fieldsLink = '/admin/fields/' + databaseID
	$('#settingsTOCFields').find("a").attr("href",fieldsLink)
	
	var formsLink = "/admin/forms/" + databaseID
	$('#settingsTOCForms').find("a").attr("href",formsLink)
	
	var tablesLink = "/admin/tables/" + databaseID
	$('#settingsTOCTables').find("a").attr("href",tablesLink)

	var listsLink = "/admin/lists/" + databaseID
	$('#settingsTOCLists').find("a").attr("href",listsLink)

	var valueListsLink = "/admin/valuelists/" + databaseID
	$('#settingsTOCValueLists').find("a").attr("href",valueListsLink)
	
	var formLinkLink = "/admin/formlink/" + databaseID
	$('#settingsTOCFormLinks').find("a").attr("href",formLinkLink)

	var dashboardLink = "/admin/dashboards/" + databaseID
	$('#settingsTOCDashboards').find("a").attr("href",dashboardLink)
	
	var globalLink = "/admin/globals/" + databaseID
	$('#settingsTOCGlobals').find("a").attr("href",globalLink)
	
	var userLink = "/admin/collaborators/" + databaseID
	$('#settingsTOCUsers').find("a").attr("href",userLink)

	var roleLink = "/admin/roles/" + databaseID
	$('#settingsTOCRoles').find("a").attr("href",roleLink)
	
	var alertLink = "/admin/alerts/" + databaseID
	$('#settingsTOCAlerts').find("a").attr("href",alertLink)
	
	
}