function addRoleFormPrivTableRow(formInfo) {
	
	var buttonsHTML = userRoleItemButtonsHTML()

	var privs = "Full Access"
	
	var rowHTML = '' +
		'<tr>' +
	         '<td style="vertical-align:middle;text-align:right;">' + formInfo.name +  '</td>' +
	         '<td>' + formRolePrivsButtonsHTML(formInfo.formID) +  '</td>' +
	     '</tr>'
	
	$('#roleFormPrivsTable').append(rowHTML)
}


function initRoleFormPrivSettingsTable(formsInfo) {
	
	$('#roleFormPrivsTable').empty()
	for(var formIndex = 0; formIndex < formsInfo.length; formIndex++) {
		var formInfo = formsInfo[formIndex]
		addRoleFormPrivTableRow(formInfo)
		
	}	
}