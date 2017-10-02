function initUserTagColPropertiesImpl(userTagInputCol) {
	
	setColPropsHeader(userTagInputCol)
	hideSiblingsShowOne("#userTagColProps")
	
	var elemPrefix = "userTag_"
		
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: userTagInputCol.parentTableID,
			userTagID: userTagInputCol.userTagID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/userTag/setLabelFormat", formatParams, function(updatedUserTag) {
			setColPropsHeader(updatedUserTag)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: userTagInputCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: userTagInputCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: userTagInputCol.parentTableID,
				userTagID: userTagInputCol.userTagID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/userTag/setPermissions",params,function(updatedUserTag) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	
	var validationParams = {
		initialValidationProps: userTagInputCol.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentTableID: userTagInputCol.parentTableID,
				userTagID: userTagInputCol.userTagID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("tableView/userTag/setValidation",validationParams,function(updatedUserTag) {
			})
		
		}
	}
	initUserTagValidationProperties(validationParams)
	
	
	var clearValueParams = {
		initialVal: userTagInputCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: userTagInputCol.parentTableID,
				userTagID: userTagInputCol.userTagID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/userTag/setClearValueSupported",formatParams,function(updatedDatePicker) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: userTagInputCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: userTagInputCol.parentTableID,
				userTagID: userTagInputCol.userTagID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/userTag/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
	
	var currUserParams = {
		currUserSelectable: userTagInputCol.properties.currUserSelectable,
		setCurrUserSelectable: function(isSelectable) {
			var params = {
				parentTableID: userTagInputCol.parentTableID,
				userTagID: userTagInputCol.userTagID,
				currUserSelectable: isSelectable
			}
			jsonAPIRequest("tableView/userTag/setCurrUserSelectable",params,function(updateCol) {
			})	
		}
	}
	initSelectionCurrUserProperties(currUserParams)
	
	var selectRoleProps = {
		databaseID: colPropsAdminContext.databaseID,
		initialRoles: userTagInputCol.properties.selectableRoles,
		setRolesCallback: function(selectableRoles) {
			var params = {
				parentTableID: userTagInputCol.parentTableID,
				userTagID: userTagInputCol.userTagID,
				selectableRoles: selectableRoles
			}
			jsonAPIRequest("tableView/userTag/setSelectableRoles",params,function(updateCol) {
			})
		}
	}
	initUserTagRoleProps(selectRoleProps)
	
	
}

function initUserTagColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		userTagID: columnID
	}
	jsonAPIRequest("tableView/userTag/get", getColParams, function(userTagInputCol) { 
		initUserTagColPropertiesImpl(userTagInputCol)
	})
	
	
	
}