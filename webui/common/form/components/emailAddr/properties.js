function loadEmailAddrProperties($emailAddr,emailAddrRef) {
	console.log("loading text box properties")
	
	var elemPrefix = "emailAddr_"
	
	
	var validationParams = {
		initialValidationProps: emailAddrRef.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentFormID: emailAddrRef.parentFormID,
				emailAddrID: emailAddrRef.emailAddrID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("frm/emailAddr/setValidation",validationParams,function(updatedEmailAddr) {
				setContainerComponentInfo($emailAddr,updatedEmailAddr,updatedEmailAddr.emailAddrID)
			})
		
		}
	}
	initTextInputValidationProperties(validationParams)
	
	
		
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: emailAddrRef.parentFormID,
			emailAddrID: emailAddrRef.emailAddrID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/emailAddr/setLabelFormat", formatParams, function(updatedEmailAddr) {
			setEmailAddrComponentLabel($emailAddr,updatedEmailAddr)
			setContainerComponentInfo($emailAddr,updatedEmailAddr,updatedEmailAddr.emailAddrID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: emailAddrRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: emailAddrRef.parentFormID,
			emailAddrID: emailAddrRef.emailAddrID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/emailAddr/setVisibility",params,function(updatedEmailAddr) {
			setContainerComponentInfo($emailAddr,updatedEmailAddr,updatedEmailAddr.emailAddrID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: emailAddrRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: emailAddrRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: emailAddrRef.parentFormID,
				emailAddrID: emailAddrRef.emailAddrID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/emailAddr/setPermissions",params,function(updatedEmailAddr) {
				setContainerComponentInfo($emailAddr,updatedEmailAddr,updatedEmailAddr.emailAddrID)
				initEmailAddrClearValueControl($emailAddr,updatedEmailAddr)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var clearValueParams = {
		initialVal: emailAddrRef.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentFormID: emailAddrRef.parentFormID,
				emailAddrID: emailAddrRef.emailAddrID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("frm/emailAddr/setClearValueSupported",formatParams,function(updatedEmailAddr) {
				setContainerComponentInfo($emailAddr,updatedEmailAddr,updatedEmailAddr.emailAddrID)
				initEmailAddrClearValueControl($emailAddr,updatedEmailAddr)
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: emailAddrRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: emailAddrRef.parentFormID,
				emailAddrID: emailAddrRef.emailAddrID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/emailAddr/setHelpPopupMsg",params,function(updatedEmailAddr) {
				setContainerComponentInfo($emailAddr,updatedEmailAddr,updatedEmailAddr.emailAddrID)
				updateComponentHelpPopupMsg($emailAddr, updatedEmailAddr)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: emailAddrRef.parentFormID,
		componentID: emailAddrRef.emailAddrID,
		componentLabel: 'text box',
		$componentContainer: $emailAddr
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#emailAddrProps')
		
	toggleFormulaEditorForField(emailAddrRef.properties.fieldID)
		
}