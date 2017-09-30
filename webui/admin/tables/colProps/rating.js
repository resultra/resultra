function initRatingColPropertiesImpl(ratingCol) {
	
	setColPropsHeader(ratingCol)
	
	var elemPrefix = "rating_"
	hideSiblingsShowOne("#ratingColProps")
		
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: ratingCol.parentTableID,
			ratingID: ratingCol.ratingID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/rating/setLabelFormat", formatParams, function(updateRating) {
			setColPropsHeader(updateRating)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: ratingCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: ratingCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: ratingCol.parentTableID,
				ratingID: ratingCol.ratingID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/rating/setPermissions",params,function(updatedRating) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	
	var iconParams = {
		initialIcon: ratingCol.properties.icon,
		setIcon: function(newIcon) {
			var iconParams = {
				parentTableID: ratingCol.parentTableID,
				ratingID: ratingCol.ratingID,
				icon: newIcon
			}
			jsonAPIRequest("tableView/rating/setIcon",iconParams,function(updatedRating) {
			})
		}
	}
	initRatingIconProps(iconParams)
	
	function createTooltipParams(ratingCol) {
		var tooltipParams = {
			initialTooltips: ratingCol.properties.tooltips,
			minVal: ratingCol.properties.minVal,
			maxVal: ratingCol.properties.maxVal,
			setTooltips: function(updatedTooltips) {
				var tooltipParams = {
					parentTableID: ratingCol.parentTableID,
					ratingID: ratingCol.ratingID,
					tooltips: updatedTooltips
				}
			
				jsonAPIRequest("tableView/rating/setTooltips", tooltipParams, function(updateRating) {
				})	
			}
		}
		return tooltipParams
	}
	
	initRatingTooltipProperties(createTooltipParams(ratingCol))
	
	
	function setRatingRange(minVal,maxVal) {
		console.log("Saving range propeties rating")
		var formatParams = {
			parentTableID: ratingCol.parentTableID,
			ratingID: ratingCol.ratingID,
			minVal: minVal,
			maxVal: maxVal
		}
		jsonAPIRequest("tableView/rating/setRange", formatParams, function(updatedRating) {
			initRatingTooltipProperties(createTooltipParams(updatedRating))
		})	
	}
	var ratingRangeParams = {
		setRangeCallback: setRatingRange,
		initialMinVal: ratingCol.properties.minVal,
		initialMaxVal: ratingCol.properties.maxVal
	}
	initRatingRangeProperties(ratingRangeParams)
	
	
	var clearValueParams = {
		initialVal: ratingCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: ratingCol.parentTableID,
				ratingID: ratingCol.ratingID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/rating/setClearValueSupported",formatParams,function(updateRating) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: ratingCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: ratingCol.parentTableID,
				ratingID: ratingCol.ratingID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/rating/setHelpPopupMsg",params,function(updateRating) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
}

function initRatingColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		ratingID: columnID
	}
	jsonAPIRequest("tableView/rating/get", getColParams, function(ratingCol) { 
		initRatingColPropertiesImpl(ratingCol)
	})
	
	
	
}