// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initAttachmentColPropertiesImpl(attachmentCol) {
	
	setColPropsHeader(attachmentCol)
	
	var elemPrefix = "attachment_"
	hideSiblingsShowOne("#attachmentColProps")
	
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: attachmentCol.parentTableID,
			attachmentID: attachmentCol.attachmentID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/attachment/setLabelFormat", formatParams, function(updatedAttachment) {
			setColPropsHeader(updatedAttachment)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: attachmentCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: attachmentCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: attachmentCol.parentTableID,
			attachmentID: attachmentCol.attachmentID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/attachment/setPermissions",params,function(updatedAttachment) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)

	var helpPopupParams = {
		initialMsg: attachmentCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: attachmentCol.parentTableID,
				attachmentID: attachmentCol.attachmentID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/attachment/setHelpPopupMsg",params,function(updatedAttachment) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)

	
}


function initAttachmentColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		attachmentID: columnID
	}
	jsonAPIRequest("tableView/attachment/get", getColParams, function(attachmentCol) { 
		initAttachmentColPropertiesImpl(attachmentCol)
	})
}