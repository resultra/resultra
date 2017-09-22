
function loadAttachmentProperties($attachmentContainer, attachmentRef) {
	console.log("Loading html editor properties")
	
	var elemPrefix = "image_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for attachment form component")
		var formatParams = {
			parentFormID: attachmentRef.parentFormID,
			imageID: attachmentRef.imageID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/attachment/setLabelFormat", formatParams, function(updatedAttachment) {
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
			jsonAPIRequest("frm/attachment/setPermissions",params,function(updatedAttachment) {
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

		jsonAPIRequest("frm/attachment/setValidation",validationParams,function(updatedAttachment) {
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
	
	var helpPopupParams = {
		initialMsg: attachmentRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: attachmentRef.parentFormID,
				imageID: attachmentRef.imageID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/attachment/setHelpPopupMsg",params,function(updatedAttachment) {
				setContainerComponentInfo($attachmentContainer,updatedAttachment,updatedAttachment.imageID)
				updateComponentHelpPopupMsg($attachmentContainer, updatedAttachment)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)


	// Toggle to the properties, hiding the other property panels
	hideSiblingsShowOne('#attachmentProps')

	toggleFormulaEditorForField(attachmentRef.properties.fieldID)
	
}