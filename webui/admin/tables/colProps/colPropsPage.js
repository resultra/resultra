
function setColPropsHeader(colInfo) {
	
	var $header = $('#colPropsColHeader')
	setFormComponentLabel($header,colInfo.properties.fieldID,
			colInfo.properties.labelFormat)
	
}


$(document).ready(function() {
	
	
	initFieldInfo(colPropsAdminContext.databaseID, function() {
		
		initAdminSettingsPageLayout($('#colPropsAdminPage'))	
		initAdminPageHeader(colPropsAdminContext.isSingleUserWorkspace)
		initAdminSettingsTOC(colPropsAdminContext.databaseID,"settingsTOCTables",colPropsAdminContext.isSingleUserWorkspace)
		
		appendPageSpecificBreadcrumbHeader("/admin/tables/"+colPropsAdminContext.databaseID,"Table Views")
		appendPageSpecificBreadcrumbHeader("/admin/table/"+colPropsAdminContext.tableID,colPropsAdminContext.tableName)
		appendPageSpecificBreadcrumbHeader("/admin/tablecol/"+colPropsAdminContext.columnID,"Column Settings")
		
		
		switch (colPropsAdminContext.colType) {
		case 'numberInput':
			initNumberInputColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'rating':
			initRatingColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'textInput':
			initTextInputColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'datePicker':
			initDatePickerColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'userSelection':
			initUserSelectionColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'checkbox':
			initCheckBoxColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'toggle':
			initToggleColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'button':
			initFormButtonColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'attachment':
			initAttachmentColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'note':
			initNoteColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'comment':
			initCommentColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'progress':
			initProgressColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'socialButton':
			initSocialButtonColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'tags':
			initTagColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'emailAddr':
			initEmailAddrColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'urlLink':
			initUrlLinkColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'file':
			initFileColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		case 'image':
			initImageColProperties(colPropsAdminContext.tableID, colPropsAdminContext.columnID)
			break
		default:
			console.log("Unknown column type: " + colPropsAdminContext.colType)
		}
		
	})
	
	
})