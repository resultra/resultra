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