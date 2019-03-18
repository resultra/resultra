// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function loadImageProperties($image,imageRef) {
	console.log("loading text box properties")
	
	var elemPrefix = "image_"
	
	
	var validationParams = {
		elemPrefix: elemPrefix,
		initialValidationProps: imageRef.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentFormID: imageRef.parentFormID,
				imageID: imageRef.imageID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("frm/image/setValidation",validationParams,function(updatedImage) {
				setContainerComponentInfo($image,updatedImage,updatedImage.imageID)
			})
		
		}
	}
	initValueRequiredValidationProperties(validationParams)
	
	
		
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: imageRef.parentFormID,
			imageID: imageRef.imageID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/image/setLabelFormat", formatParams, function(updatedImage) {
			setImageComponentLabel($image,updatedImage)
			setContainerComponentInfo($image,updatedImage,updatedImage.imageID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: imageRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: imageRef.parentFormID,
			imageID: imageRef.imageID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/image/setVisibility",params,function(updatedImage) {
			setContainerComponentInfo($image,updatedImage,updatedImage.imageID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: imageRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: imageRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: imageRef.parentFormID,
				imageID: imageRef.imageID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/image/setPermissions",params,function(updatedImage) {
				setContainerComponentInfo($image,updatedImage,updatedImage.imageID)
				initImageClearValueControl($image,updatedImage)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var clearValueParams = {
		initialVal: imageRef.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentFormID: imageRef.parentFormID,
				imageID: imageRef.imageID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("frm/image/setClearValueSupported",formatParams,function(updatedImage) {
				setContainerComponentInfo($image,updatedImage,updatedImage.imageID)
				initImageClearValueControl($image,updatedImage)
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: imageRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: imageRef.parentFormID,
				imageID: imageRef.imageID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/image/setHelpPopupMsg",params,function(updatedImage) {
				setContainerComponentInfo($image,updatedImage,updatedImage.imageID)
				updateComponentHelpPopupMsg($image, updatedImage)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: imageRef.parentFormID,
		componentID: imageRef.imageID,
		componentLabel: 'text box',
		$componentContainer: $image
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#imageProps')
		
	toggleFormulaEditorForField(imageRef.properties.fieldID)
		
}