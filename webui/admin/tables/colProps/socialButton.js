function initSocialButtonColPropertiesImpl(socialButtonCol) {
	
	setColPropsHeader(socialButtonCol)
	
	var elemPrefix = "socialButton_"
	hideSiblingsShowOne("#socialButtonColProps")
		
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: socialButtonCol.parentTableID,
			socialButtonID: socialButtonCol.socialButtonID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/socialButton/setLabelFormat", formatParams, function(updateRating) {
			setColPropsHeader(updateRating)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: socialButtonCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: socialButtonCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: socialButtonCol.parentTableID,
				socialButtonID: socialButtonCol.socialButtonID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/socialButton/setPermissions",params,function(updatedRating) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	
	var iconParams = {
		initialIcon: socialButtonCol.properties.icon,
		setIcon: function(newIcon) {
			var iconParams = {
				parentTableID: socialButtonCol.parentTableID,
				socialButtonID: socialButtonCol.socialButtonID,
				icon: newIcon
			}
			jsonAPIRequest("tableView/socialButton/setIcon",iconParams,function(updatedRating) {
			})
		}
	}
	initSocialButtonIconProps(iconParams)
	
	
	var helpPopupParams = {
		initialMsg: socialButtonCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: socialButtonCol.parentTableID,
				socialButtonID: socialButtonCol.socialButtonID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/socialButton/setHelpPopupMsg",params,function(updateRating) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
}

function initSocialButtonColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		socialButtonID: columnID
	}
	jsonAPIRequest("tableView/socialButton/get", getColParams, function(socialButtonCol) { 
		initSocialButtonColPropertiesImpl(socialButtonCol)
	})
	
	
	
}