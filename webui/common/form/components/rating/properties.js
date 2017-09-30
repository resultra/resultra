function loadRatingProperties($rating,ratingRef) {
	console.log("Loading rating properties")
	
	var iconParams = {
		initialIcon: ratingRef.properties.icon,
		setIcon: function(newIcon) {
			var iconParams = {
				parentFormID: ratingRef.parentFormID,
				ratingID: ratingRef.ratingID,
				icon: newIcon
			}
			jsonAPIRequest("frm/rating/setIcon",iconParams,function(updatedRating) {
				setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
				reInitRatingFormComponentControl($rating,updatedRating)
			})
		}
	}
	initRatingIconProps(iconParams)
	
	
	function createTooltipParams(ratingRef) {
		
		function updateTooltips(updatedTooltips) {
			var tooltipParams = {
				parentFormID: ratingRef.parentFormID,
				ratingID: ratingRef.ratingID,
				tooltips: updatedTooltips,
			}
		
			jsonAPIRequest("frm/rating/setTooltips", tooltipParams, function(updateRating) {
				setContainerComponentInfo($rating,updateRating,updateRating.ratingID)
			})
		}
		
		var tooltipParams = {
			initialTooltips: ratingRef.properties.tooltips,
			minVal: ratingRef.properties.minVal,
			maxVal: ratingRef.properties.maxVal,
			setTooltips: updateTooltips
		}
		return tooltipParams
	}
	
	initRatingTooltipProperties(createTooltipParams(ratingRef))

	function setRatingRange(minVal,maxVal) {
		console.log("Saving range propeties rating")
		var formatParams = {
			parentFormID: ratingRef.parentFormID,
			ratingID: ratingRef.ratingID,
			minVal: minVal,
			maxVal: maxVal
		}
		jsonAPIRequest("frm/rating/setRange", formatParams, function(updatedRating) {
			setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
			reInitRatingFormComponentControl($rating,updatedRating)
			initRatingTooltipProperties(createTooltipParams(updatedRating))
		})	
	}
	var ratingRangeParams = {
		setRangeCallback: setRatingRange,
		initialMinVal: ratingRef.properties.minVal,
		initialMaxVal: ratingRef.properties.maxVal
	}
	initRatingRangeProperties(ratingRangeParams)

	var elemPrefix = "rating_"
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentFormID: ratingRef.parentFormID,
			ratingID: ratingRef.ratingID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/rating/setLabelFormat", formatParams, function(updatedRating) {
			setRatingComponentLabel($rating,updatedRating)
			setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
		})	
	}
	
	
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: ratingRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	

	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: ratingRef.parentFormID,
			ratingID: ratingRef.ratingID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/rating/setVisibility",params,function(updatedRating) {
			setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: ratingRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)

	initCheckboxChangeHandler('#adminRatingComponentValidationRequired', 
				ratingRef.properties.validation.valueRequired, function (newVal) {
			
		var validationProps = {
			valueRequired: newVal
		}		
			
		var validationParams = {
			parentFormID: ratingRef.parentFormID,
			ratingID: ratingRef.ratingID,
			validation: validationProps
		}
		console.log("Setting new validation settings: " + JSON.stringify(validationParams))

		jsonAPIRequest("frm/rating/setValidation",validationParams,function(updatedRating) {
				setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
		})
	})
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: ratingRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: ratingRef.parentFormID,
				ratingID: ratingRef.ratingID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/rating/setPermissions",params,function(updatedRating) {
				setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	
	var clearValueParams = {
		initialVal: ratingRef.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentFormID: ratingRef.parentFormID,
				ratingID: ratingRef.ratingID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("frm/rating/setClearValueSupported",formatParams,function(updatedRating) {
				setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: ratingRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: ratingRef.parentFormID,
				ratingID: ratingRef.ratingID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/rating/setHelpPopupMsg",params,function(updatedRating) {
				setContainerComponentInfo($rating,updatedRating,updatedRating.ratingID)
				updateComponentHelpPopupMsg($rating, updatedRating)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: ratingRef.parentFormID,
		componentID: ratingRef.ratingID,
		componentLabel: 'rating',
		$componentContainer: $rating
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#ratingProps')
		
	toggleFormulaEditorForField(ratingRef.properties.fieldID)
	
}