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
			jsonAPIRequest("tableView/attachment/setPermissions",params,function(updatedCheckBox) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
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