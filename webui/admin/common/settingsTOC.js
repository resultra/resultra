function initAdminSettingsTOC(databaseID) {

	var generalLink = '/admin/' + databaseID
	$('#settingsTOCGeneral').attr("href",generalLink)

	var fieldsLink = '/admin/fields/' + databaseID
	$('#settingsTOCFields').attr("href",fieldsLink)
	
	var formsLink = "/admin/forms/" + databaseID
	$('#settingsTOCForms').attr("href",formsLink)
	
}