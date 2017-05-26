
function setColPropsHeader(colInfo) {
	
	var $header = $('#colPropsColHeader')
	setFormComponentLabel($header,colInfo.properties.fieldID,
			colInfo.properties.labelFormat)
	
}


$(document).ready(function() {
	
	
	initFieldInfo(colPropsAdminContext.databaseID, function() {
		
		initAdminSettingsPageLayout($('#colPropsAdminPage'))	
		initUserDropdownMenu()
		initAdminSettingsTOC(colPropsAdminContext.databaseID)
		
		switch (colPropsAdminContext.colType) {
		case 'numberInput':
			initNumberInputColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'textInput':
			initTextInputColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'datePicker':
			initDatePickerColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		default:
			console.log("Unknown column type: " + colPropsAdminContext.colType)
		}
		
	})
	
	
})