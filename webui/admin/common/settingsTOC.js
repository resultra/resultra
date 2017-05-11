function initAdminSettingsTOC(databaseID) {

	var generalLink = '/admin/general/' + databaseID
	$('#settingsTOCGeneral').attr("href",generalLink)

	var fieldsLink = '/admin/fields/' + databaseID
	$('#settingsTOCFields').attr("href",fieldsLink)
	
	var formsLink = "/admin/forms/" + databaseID
	$('#settingsTOCForms').attr("href",formsLink)

	var listsLink = "/admin/lists/" + databaseID
	$('#settingsTOCLists').attr("href",listsLink)

	var valueListsLink = "/admin/valuelists/" + databaseID
	$('#settingsTOCValueLists').attr("href",valueListsLink)
	
	var formLinkLink = "/admin/formlink/" + databaseID
	$('#settingsTOCFormLinks').attr("href",formLinkLink)

	var dashboardLink = "/admin/dashboards/" + databaseID
	$('#settingsTOCDashboards').attr("href",dashboardLink)
	
	var globalLink = "/admin/globals/" + databaseID
	$('#settingsTOCGlobals').attr("href",globalLink)
	
	var userLink = "/admin/users/" + databaseID
	$('#settingsTOCUsers').attr("href",userLink)

	var roleLink = "/admin/roles/" + databaseID
	$('#settingsTOCRoles').attr("href",roleLink)
	
}