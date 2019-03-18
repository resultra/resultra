// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function loadLabelProperties($label,labelRef) {
	console.log("Loading user selection properties")
	
	var elemPrefix = "tag_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for user selection")
		var formatParams = {
			parentFormID: labelRef.parentFormID,
			labelID: labelRef.labelID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/label/setLabelFormat", formatParams, function(updatedLabel) {
			setLabelComponentLabel($label,updatedLabel)
			setContainerComponentInfo($label,updatedLabel,updatedLabel.labelID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: labelRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: labelRef.parentFormID,
			labelID: labelRef.labelID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/label/setVisibility",params,function(updatedLabel) {
			setContainerComponentInfo($label,updatedLabel,updatedLabel.labelID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: labelRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
		
	var permissionParams = {
		elemPrefix: elemPrefix,
		initialVal: labelRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: labelRef.parentFormID,
				labelID: labelRef.labelID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/label/setPermissions",params,function(updatedLabel) {
				setContainerComponentInfo($label,updatedLabel,updatedLabel.labelID)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(permissionParams)
		

	var validationParams = {
		setValidation: function(validationProps) {
			var validationParams = {
				parentFormID: labelRef.parentFormID,
				labelID: labelRef.labelID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("frm/label/setValidation",validationParams,function(updatedLabel) {
				setContainerComponentInfo($label,updatedLabel,updatedLabel.labelID)
			})
		}
	}
	initLabelValidationProperties(validationParams)
	
	var helpPopupParams = {
		initialMsg: labelRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: labelRef.parentFormID,
				labelID: labelRef.labelID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/label/setHelpPopupMsg",params,function(updatedLabel) {
				setContainerComponentInfo($label,updatedLabel,updatedLabel.labelID)
				updateComponentHelpPopupMsg($label, updatedLabel)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)


	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: labelRef.parentFormID,
		componentID: labelRef.labelID,
		componentLabel: 'label',
		$componentContainer: $label
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#labelProps')
		
	toggleFormulaEditorForField(labelRef.properties.fieldID)
	
}