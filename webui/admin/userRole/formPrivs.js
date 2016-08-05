function addRoleFormPrivTableRow(formID,roleID) {
	
	var buttonsHTML = userRoleItemButtonsHTML()

	var privs = "Full Access"
	
	var rowHTML = '' +
		'<tr>' +
	         '<td style="vertical-align:middle;text-align:right;">' + formID +  '</td>' +
	         '<td>' + formRolePrivsButtonsHTML(roleID) +  '</td>' +
	     '</tr>'
	
	$('#roleFormPrivsTable').append(rowHTML)
}


function initRoleFormPrivSettingsTable() {
	$('#roleFormPrivsTable').empty()
	addRoleFormPrivTableRow("form1")
	addRoleFormPrivTableRow("form2")
	addRoleFormPrivTableRow("form3")
	
}