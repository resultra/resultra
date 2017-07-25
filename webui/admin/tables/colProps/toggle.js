function initToggleColPropertiesImpl(toggleCol) {
	
	setColPropsHeader(toggleCol)
	
	var elemPrefix = "toggle_"
	hideSiblingsShowOne("#toggleColProps")
	
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: toggleCol.parentTableID,
			toggleID: toggleCol.toggleID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/toggle/setLabelFormat", formatParams, function(updatedToggle) {
			setColPropsHeader(updatedToggle)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: toggleCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: toggleCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: toggleCol.parentTableID,
				toggleID: toggleCol.toggleID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/toggle/setPermissions",params,function(updatedToggle) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var clearValueParams = {
		initialVal: toggleCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: toggleCol.parentTableID,
				toggleID: toggleCol.toggleID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/toggle/setClearValueSupported",formatParams,function(updatedToggle) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	
	var formatParams = {
		initialVals: toggleCol.properties,
		setOffLabel: function(newLabel) {	
			var labelParams = {
				parentTableID: toggleCol.parentTableID,
				toggleID: toggleCol.toggleID,
				label: newLabel
			}
			jsonAPIRequest("tableView/toggle/setOffLabel",labelParams,function(updatedToggleRef) {
			})
		},
		setOnLabel: function(newLabel) {
			var labelParams = {
				parentTableID: toggleCol.parentTableID,
				toggleID: toggleCol.toggleID,
				label: newLabel
			}
			jsonAPIRequest("tableView/toggle/setOnLabel",labelParams,function(updatedToggleRef) {
			})
		},
		setOffColorScheme: function(newColorScheme) {
			var colorSchemeParams = {
				parentTableID: toggleCol.parentTableID,
				toggleID: toggleCol.toggleID,
				colorScheme: newColorScheme
			}
			jsonAPIRequest("tableView/toggle/setOffColorScheme",colorSchemeParams,function(updatedToggleRef) {
			})
		},
		setOnColorScheme: function(newColorScheme) {
			var colorSchemeParams = {
				parentTableID: toggleCol.parentTableID,
				toggleID: toggleCol.toggleID,
				colorScheme: newColorScheme
			}
			jsonAPIRequest("tableView/toggle/setOnColorScheme",colorSchemeParams,function(updatedToggleRef) {
			})		
		}
	}
	initToggleFormatProperties(formatParams)
	
	var helpPopupParams = {
		initialMsg: toggleCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: toggleCol.parentTableID,
				toggleID: toggleCol.toggleID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/toggle/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
}


function initToggleColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		toggleID: columnID
	}
	jsonAPIRequest("tableView/toggle/get", getColParams, function(toggleCol) { 
		initToggleColPropertiesImpl(toggleCol)
	})
}