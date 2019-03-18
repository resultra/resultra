// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function loadUrlLinkProperties($urlLink,urlLinkRef) {
	console.log("loading text box properties")
	
	var elemPrefix = "urlLink_"
	
	
	var validationParams = {
		elemPrefix:elemPrefix,
		initialValidationProps: urlLinkRef.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentFormID: urlLinkRef.parentFormID,
				urlLinkID: urlLinkRef.urlLinkID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("frm/urlLink/setValidation",validationParams,function(updatedUrlLink) {
				setContainerComponentInfo($urlLink,updatedUrlLink,updatedUrlLink.urlLinkID)
			})
		
		}
	}
	initValueRequiredValidationProperties(validationParams)
	
	
		
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: urlLinkRef.parentFormID,
			urlLinkID: urlLinkRef.urlLinkID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/urlLink/setLabelFormat", formatParams, function(updatedUrlLink) {
			setUrlLinkComponentLabel($urlLink,updatedUrlLink)
			setContainerComponentInfo($urlLink,updatedUrlLink,updatedUrlLink.urlLinkID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: urlLinkRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: urlLinkRef.parentFormID,
			urlLinkID: urlLinkRef.urlLinkID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/urlLink/setVisibility",params,function(updatedUrlLink) {
			setContainerComponentInfo($urlLink,updatedUrlLink,updatedUrlLink.urlLinkID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: urlLinkRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: urlLinkRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: urlLinkRef.parentFormID,
				urlLinkID: urlLinkRef.urlLinkID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/urlLink/setPermissions",params,function(updatedUrlLink) {
				setContainerComponentInfo($urlLink,updatedUrlLink,updatedUrlLink.urlLinkID)
				initUrlLinkClearValueControl($urlLink,updatedUrlLink)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var clearValueParams = {
		initialVal: urlLinkRef.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentFormID: urlLinkRef.parentFormID,
				urlLinkID: urlLinkRef.urlLinkID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("frm/urlLink/setClearValueSupported",formatParams,function(updatedUrlLink) {
				setContainerComponentInfo($urlLink,updatedUrlLink,updatedUrlLink.urlLinkID)
				initUrlLinkClearValueControl($urlLink,updatedUrlLink)
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: urlLinkRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: urlLinkRef.parentFormID,
				urlLinkID: urlLinkRef.urlLinkID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/urlLink/setHelpPopupMsg",params,function(updatedUrlLink) {
				setContainerComponentInfo($urlLink,updatedUrlLink,updatedUrlLink.urlLinkID)
				updateComponentHelpPopupMsg($urlLink, updatedUrlLink)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: urlLinkRef.parentFormID,
		componentID: urlLinkRef.urlLinkID,
		componentLabel: 'url link input',
		$componentContainer: $urlLink
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#urlLinkProps')
		
	toggleFormulaEditorForField(urlLinkRef.properties.fieldID)
		
}