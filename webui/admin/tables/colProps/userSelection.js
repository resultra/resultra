// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initUserSelectionColPropertiesImpl(pageContext,userSelectionInputCol) {
	
	setColPropsHeader(userSelectionInputCol)
	hideSiblingsShowOne("#userSelectionColProps")
	
	var elemPrefix = "userSelection_"
		
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: userSelectionInputCol.parentTableID,
			userSelectionID: userSelectionInputCol.userSelectionID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/userSelection/setLabelFormat", formatParams, function(updatedUserSelection) {
			setColPropsHeader(updatedUserSelection)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: userSelectionInputCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: userSelectionInputCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: userSelectionInputCol.parentTableID,
				userSelectionID: userSelectionInputCol.userSelectionID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/userSelection/setPermissions",params,function(updatedUserSelection) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	
	var validationParams = {
		valueRequired: userSelectionInputCol.properties.validation.valueRequired,
		setValidation: function(validationProps) {
			var validationParams = {
				parentTableID: userSelectionInputCol.parentTableID,
				userSelectionID: userSelectionInputCol.userSelectionID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("tableView/userSelection/setValidation",validationParams,function(updatedUserSelection) {
			})
		
		}
	}
	initUserSelectionValidationProperties(validationParams)
	
	
	var clearValueParams = {
		initialVal: userSelectionInputCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: userSelectionInputCol.parentTableID,
				userSelectionID: userSelectionInputCol.userSelectionID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/userSelection/setClearValueSupported",formatParams,function(updatedDatePicker) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: userSelectionInputCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: userSelectionInputCol.parentTableID,
				userSelectionID: userSelectionInputCol.userSelectionID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/userSelection/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
	
	var currUserParams = {
		elemPrefix: elemPrefix,	
		currUserSelectable: userSelectionInputCol.properties.currUserSelectable,
		setCurrUserSelectable: function(isSelectable) {
			var params = {
				parentTableID: userSelectionInputCol.parentTableID,
				userSelectionID: userSelectionInputCol.userSelectionID,
				currUserSelectable: isSelectable
			}
			jsonAPIRequest("tableView/userSelection/setCurrUserSelectable",params,function(updateCol) {
			})	
		}
	}
	initSelectionCurrUserProperties(currUserParams)
	
	var selectRoleProps = {
		elemPrefix: elemPrefix,	
		databaseID: pageContext.databaseID,
		initialRoles: userSelectionInputCol.properties.selectableRoles,
		setRolesCallback: function(selectableRoles) {
			var params = {
				parentTableID: userSelectionInputCol.parentTableID,
				userSelectionID: userSelectionInputCol.userSelectionID,
				selectableRoles: selectableRoles
			}
			jsonAPIRequest("tableView/userSelection/setSelectableRoles",params,function(updateCol) {
			})
		}
	}
	initUserSelectionRoleProps(selectRoleProps)
	
	
}

function initUserSelectionColProperties(pageContext,tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		userSelectionID: columnID
	}
	jsonAPIRequest("tableView/userSelection/get", getColParams, function(userSelectionInputCol) { 
		initUserSelectionColPropertiesImpl(pageContext,userSelectionInputCol)
	})
	
	
	
}