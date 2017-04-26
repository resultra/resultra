
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

	initCheckboxChangeHandler('#adminAttachmentComponentValidationRequired', 
				attachmentRef.properties.validation.valueRequired, function (newVal) {
		var validationProps = {
			valueRequired: newVal
		}		
		var validationParams = {
			parentFormID: attachmentRef.parentFormID,
			imageID: attachmentRef.imageID,
			validation: validationProps
		}
		console.log("Setting new validation settings: " + JSON.stringify(validationParams))

		jsonAPIRequest("frm/image/setValidation",validationParams,function(updatedAttachment) {
			setContainerComponentInfo($attachmentContainer,updatedAttachment,updatedAttachment.imageID)
		})
	})


	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: attachmentRef.parentFormID,
		componentID: attachmentRef.imageID,
		componentLabel: 'attachment box',
		$componentContainer: $attachmentContainer
	}
	initDeleteFormComponentPropertyPanel(deleteParams)


	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#imageProps')

	toggleFormulaEditorForField(attachmentRef.properties.fieldID)
	
}