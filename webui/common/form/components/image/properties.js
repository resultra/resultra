
function loadImageProperties($attachmentContainer, attachmentRef) {
	console.log("Loading html editor properties")
	
	var elemPrefix = "image_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for attachment form component")
		var formatParams = {
			parentFormID: attachmentRef.parentFormID,
			imageID: attachmentRef.imageID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/image/setLabelFormat", formatParams, function(updatedAttachment) {
			setAttachmentComponentLabel($attachmentContainer,updatedAttachment)
			setContainerComponentInfo($attachmentContainer,updatedAttachment,updatedAttachment.imageID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: attachmentRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)

	var permissionParams = {
		elemPrefix: elemPrefix,
		initialVal: attachmentRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: attachmentRef.parentFormID,
				imageID: attachmentRef.imageID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/image/setPermissions",params,function(updatedAttachment) {
				setContainerComponentInfo($attachmentContainer,updatedAttachment,updatedAttachment.imageID)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(permissionParams)

	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#imageProps')

	toggleFormulaEditorForField(attachmentRef.properties.fieldID)
	
}