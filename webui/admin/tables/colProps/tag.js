function initTagColPropertiesImpl(tagInputCol) {
	
	setColPropsHeader(tagInputCol)
	hideSiblingsShowOne("#tagColProps")
	
	var elemPrefix = "tag_"
		
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for tag column")
		var formatParams = {
			parentTableID: tagInputCol.parentTableID,
			tagID: tagInputCol.tagID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/tag/setLabelFormat", formatParams, function(updatedTag) {
			setColPropsHeader(updatedTag)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: tagInputCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: tagInputCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: tagInputCol.parentTableID,
				tagID: tagInputCol.tagID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/tag/setPermissions",params,function(updatedTag) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	
	var validationParams = {
		initialValidationProps: tagInputCol.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentTableID: tagInputCol.parentTableID,
				tagID: tagInputCol.tagID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("tableView/tag/setValidation",validationParams,function(updatedTag) {
			})
		
		}
	}
	initLabelValidationProperties(validationParams)
		
	var helpPopupParams = {
		initialMsg: tagInputCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: tagInputCol.parentTableID,
				tagID: tagInputCol.tagID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/tag/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
		
	
}

function initTagColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		tagID: columnID
	}
	jsonAPIRequest("tableView/tag/get", getColParams, function(tagInputCol) { 
		initTagColPropertiesImpl(tagInputCol)
	})
	
	
	
}