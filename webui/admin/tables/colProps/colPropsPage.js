
function setColPropsHeader(colInfo) {
	
	var $header = $('#colPropsColHeader')
	setFormComponentLabel($header,colInfo.properties.fieldID,
			colInfo.properties.labelFormat)
	
}


$(document).ready(function() {
	
	initAdminSettingsPageLayout($('#colPropsAdminPage'))	
	initUserDropdownMenu()
	initAdminSettingsTOC(colPropsAdminContext.databaseID)
	
	switch (colPropsAdminContext.colType) {
	case 'numberInput':
		initNumberInputColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
	default:
		console.log("Unknown column type: " + colPropsAdminContext.colType)
	}
	
})