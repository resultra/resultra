// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function setColPropsHeader(colInfo) {
	
	var $header = $('#colPropsColHeader')
	setFormComponentLabel($header,colInfo.properties.fieldID,
			colInfo.properties.labelFormat)
	
}


function initTableColPropsPageConent(pageContext,columnInfo) {
	
	
	initFieldInfo(pageContext.databaseID, function() {
				
		switch (columnInfo.colType) {
		case 'numberInput':
			initNumberInputColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'rating':
			initRatingColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'textInput':
			initTextInputColProperties(pageContext,columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'textSelection':
			initTextSelectionColProperties(pageContext,columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'datePicker':
			initDatePickerColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'userSelection':
			initUserSelectionColProperties(pageContext,columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'checkbox':
			initCheckBoxColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'toggle':
			initToggleColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'button':
			initFormButtonColProperties(pageContext,columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'attachment':
			initAttachmentColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'note':
			initNoteColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'comment':
			initCommentColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'progress':
			initProgressColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'socialButton':
			initSocialButtonColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'tags':
			initTagColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'emailAddr':
			initEmailAddrColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'urlLink':
			initUrlLinkColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'file':
			initFileColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		case 'image':
			initImageColProperties(columnInfo.parentTableID, columnInfo.columnID)
			break
		default:
			console.log("Unknown column type: " + columnInfo.colType)
		}
		
	})
	
	
	const tableSettingsLinkID = 'table-' + columnInfo.parentTableID
	initSettingsPageButtonLink('#colPropsBackToTableSettingsLink',tableSettingsLinkID)
	
	
}

function navigateToNewColumnSettingsPage(pageContext,columnInfo) {
	var contentURL = '/admin/tablecol/'+columnInfo.columnID
	setSettingsPageContent(contentURL, function() {
		initTableColPropsPageConent(pageContext,columnInfo)
	})
}
