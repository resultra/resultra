function loadRatingProperties($socialButton,socialButtonRef) {
	console.log("Loading rating properties")
	
	var iconParams = {
		initialIcon: socialButtonRef.properties.icon,
		setIcon: function(newIcon) {
			var iconParams = {
				parentFormID: socialButtonRef.parentFormID,
				socialButtonID: socialButtonRef.socialButtonID,
				icon: newIcon
			}
			jsonAPIRequest("frm/socialButton/setIcon",iconParams,function(updatedSocialButton) {
				setContainerComponentInfo($socialButton,updatedSocialButton,updatedSocialButton.socialButtonID)
				reInitRatingFormComponentControl($socialButton,updatedSocialButton)
			})
		}
	}
	initRatingIconProps(iconParams)
	
	
	var tooltipParams = {
		initialTooltips: socialButtonRef.properties.tooltips,
		setTooltips: function(updatedTooltips) {
			var tooltipParams = {
				parentFormID: socialButtonRef.parentFormID,
				socialButtonID: socialButtonRef.socialButtonID,
				tooltips: updatedTooltips
			}
			
			jsonAPIRequest("frm/socialButton/setTooltips", tooltipParams, function(updateRating) {
				setContainerComponentInfo($socialButton,updateRating,updateRating.ratingID)
			})	
		}
	}
	initRatingTooltipProperties(tooltipParams)


	var elemPrefix = "socialButton_"
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: socialButtonRef.parentFormID,
			socialButtonID: socialButtonRef.socialButtonID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/socialButton/setLabelFormat", formatParams, function(updatedSocialButton) {
			setRatingComponentLabel($socialButton,updatedSocialButton)
			setContainerComponentInfo($socialButton,updatedSocialButton,updatedSocialButton.socialButtonID)
		})	
	}
	
	
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: socialButtonRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)

	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: socialButtonRef.parentFormID,
			socialButtonID: socialButtonRef.socialButtonID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/socialButton/setVisibility",params,function(updatedSocialButton) {
			setContainerComponentInfo($socialButton,updatedSocialButton,updatedSocialButton.socialButtonID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: socialButtonRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: socialButtonRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: socialButtonRef.parentFormID,
				socialButtonID: socialButtonRef.socialButtonID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/socialButton/setPermissions",params,function(updatedSocialButton) {
				setContainerComponentInfo($socialButton,updatedSocialButton,updatedSocialButton.socialButtonID)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
		
	var helpPopupParams = {
		initialMsg: socialButtonRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: socialButtonRef.parentFormID,
				socialButtonID: socialButtonRef.socialButtonID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/socialButton/setHelpPopupMsg",params,function(updatedSocialButton) {
				setContainerComponentInfo($socialButton,updatedSocialButton,updatedSocialButton.socialButtonID)
				updateComponentHelpPopupMsg($socialButton, updatedSocialButton)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: socialButtonRef.parentFormID,
		componentID: socialButtonRef.socialButtonID,
		componentLabel: 'socialButton',
		$componentContainer: $socialButton
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#socialButtonProps')
		
	toggleFormulaEditorForField(socialButtonRef.properties.fieldID)
	
}