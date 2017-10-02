function loadUserTagProperties($userTag,userTagRef) {
	console.log("Loading user selection properties")
	
	var elemPrefix = "userTag_"
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for user selection")
		var formatParams = {
			parentFormID: userTagRef.parentFormID,
			userTagID: userTagRef.userTagID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("frm/userTag/setLabelFormat", formatParams, function(updatedUserTag) {
			setUserTagComponentLabel($userTag,updatedUserTag)
			setContainerComponentInfo($userTag,updatedUserTag,updatedUserTag.userTagID)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: userTagRef.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	function saveVisibilityConditions(updatedConditions) {
		var params = {
			parentFormID: userTagRef.parentFormID,
			userTagID: userTagRef.userTagID,
			visibilityConditions: updatedConditions
		}
		jsonAPIRequest("frm/userTag/setVisibility",params,function(updatedUserTag) {
			setContainerComponentInfo($userTag,updatedUserTag,updatedUserTag.userTagID)
		})
	}
	var visibilityParams = {
		elemPrefix: elemPrefix,
		// TODO - pass in database ID as part of the component's context, rather than reference a global.
		databaseID: designFormContext.databaseID,
		initialConditions: userTagRef.properties.visibilityConditions,
		saveVisibilityConditionsCallback:saveVisibilityConditions
	}
	initFormComponentVisibilityPropertyPanel(visibilityParams)
		
	var permissionParams = {
		elemPrefix: elemPrefix,
		initialVal: userTagRef.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentFormID: userTagRef.parentFormID,
				userTagID: userTagRef.userTagID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("frm/userTag/setPermissions",params,function(updatedUserTag) {
				setContainerComponentInfo($userTag,updatedUserTag,updatedUserTag.userTagID)
				initUserTagClearValueButton($userTag,updatedUserTag)
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(permissionParams)
	
	var clearValueParams = {
		initialVal: userTagRef.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentFormID: userTagRef.parentFormID,
				userTagID: userTagRef.userTagID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("frm/userTag/setClearValueSupported",formatParams,function(updatedUserTag) {
					setContainerComponentInfo($userTag,updatedUserTag,updatedUserTag.userTagID)
					initUserTagClearValueButton($userTag,updatedUserTag)
			})
		}
	}
	initClearValueProps(clearValueParams)
	

	initCheckboxChangeHandler('#adminUserTagComponentValidationRequired', 
				userTagRef.properties.validation.valueRequired, function (newVal) {
		
		var validationProps = {
			valueRequired: newVal
		}		
		
		var validationParams = {
			parentFormID: userTagRef.parentFormID,
			userTagID: userTagRef.userTagID,
			validation: validationProps
		}
		console.log("Setting new validation settings: " + JSON.stringify(validationParams))

		jsonAPIRequest("frm/userTag/setValidation",validationParams,function(updatedUserTag) {
			setContainerComponentInfo($userTag,updatedUserTag,updatedUserTag.userTagID)
		})
	})
	
	var helpPopupParams = {
		initialMsg: userTagRef.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentFormID: userTagRef.parentFormID,
				userTagID: userTagRef.userTagID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("frm/userTag/setHelpPopupMsg",params,function(updatedUserTag) {
				setContainerComponentInfo($userTag,updatedUserTag,updatedUserTag.userTagID)
				updateComponentHelpPopupMsg($userTag, updatedUserTag)
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)


	var deleteParams = {
		elemPrefix: elemPrefix,
		parentFormID: userTagRef.parentFormID,
		componentID: userTagRef.userTagID,
		componentLabel: 'user selection',
		$componentContainer: $userTag
	}
	initDeleteFormComponentPropertyPanel(deleteParams)
	
	var currUserParams = {
		elemPrefix: elemPrefix,
		currUserSelectable: userTagRef.properties.currUserSelectable,
		setCurrUserSelectable: function(isSelectable) {
			var params = {
				parentFormID: userTagRef.parentFormID,
				userTagID: userTagRef.userTagID,
				currUserSelectable: isSelectable
			}
			jsonAPIRequest("frm/userTag/setCurrUserSelectable",params,function(updatedUserTag) {
				setContainerComponentInfo($userTag,updatedUserTag,updatedUserTag.userTagID)
			})	
		}
	}
	initSelectionCurrUserProperties(currUserParams)
	
	var selectRoleProps = {
		elemPrefix: elemPrefix,
		databaseID: designFormContext.databaseID,
		initialRoles: userTagRef.properties.selectableRoles,
		setRolesCallback: function(selectableRoles) {
			var params = {
				parentFormID: userTagRef.parentFormID,
				userTagID: userTagRef.userTagID,
				selectableRoles: selectableRoles
			}
			jsonAPIRequest("frm/userTag/setSelectableRoles",params,function(updatedUserTag) {
				setContainerComponentInfo($userTag,updatedUserTag,updatedUserTag.userTagID)
			})
		}
	}
	initUserSelectionRoleProps(selectRoleProps)

	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#userTagProps')
		
	toggleFormulaEditorForField(userTagRef.properties.fieldID)
	
}