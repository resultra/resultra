// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function loadFileProperties($file,fileRef) {
	console.log("loading text box properties")
	
	var elemPrefix = "file_"
	
	
	var validationParams = {
		elemPrefix: elemPrefix,
		initialValidationProps: fileRef.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentFormID: fileRef.parentFormID,
				fileID: fileRef.fileID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("frm/file/setValidation",validationParams,function(updatedFile) {
				setContainerComponentInfo($file,updatedFile,updatedFile.fileID)
			})
		
		}
	}
	initValueRequiredValidationProperties(validationParams)
	
	
		
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: fileRef.parentFormID,
			fileID: fileRef.fileID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/file/setLabelFormat", formatParams, function(updatedFile) {
			setFileComponentLabel($file,updatedFile)
			setContainerComponentInfo($file,updatedFile,updatedFile.fileID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: fileRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: fileRef.parentFormID,
			fileID: fileRef.fileID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/file/setVisibility",params,function(updatedFile) {
			setContainerComponentInfo($file,updatedFile,updatedFile.fileID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: fileRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: fileRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: fileRef.parentFormID,
				fileID: fileRef.fileID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/file/setPermissions",params,function(updatedFile) {
				setContainerComponentInfo($file,updatedFile,updatedFile.fileID)
				initFileClearValueControl($file,updatedFile)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var clearValueParams = {
		initialVal: fileRef.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentFormID: fileRef.parentFormID,
				fileID: fileRef.fileID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("frm/file/setClearValueSupported",formatParams,function(updatedFile) {
				setContainerComponentInfo($file,updatedFile,updatedFile.fileID)
				initFileClearValueControl($file,updatedFile)
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: fileRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: fileRef.parentFormID,
				fileID: fileRef.fileID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/file/setHelpPopupMsg",params,function(updatedFile) {
				setContainerComponentInfo($file,updatedFile,updatedFile.fileID)
				updateComponentHelpPopupMsg($file, updatedFile)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: fileRef.parentFormID,
		componentID: fileRef.fileID,
		componentLabel: 'text box',
		$componentContainer: $file
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#fileProps')
		
	toggleFormulaEditorForField(fileRef.properties.fieldID)
		
}