function initImageColPropertiesImpl(imageCol) {
	
	setColPropsHeader(imageCol)
	hideSiblingsShowOne("#imageColProps")
	
	var elemPrefix = "image_"
		
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for image column")
		var formatParams = {
			parentTableID: imageCol.parentTableID,
			imageID: imageCol.imageID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/image/setLabelFormat", formatParams, function(updatedNumberInput) {
			setColPropsHeader(updatedNumberInput)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: imageCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: imageCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: imageCol.parentTableID,
				imageID: imageCol.imageID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/image/setPermissions",params,function(updatedImage) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
		
	
	var validationParams = {
		initialValidationProps: imageCol.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentTableID: imageCol.parentTableID,
				imageID: imageCol.imageID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("tableView/image/setValidation",validationParams,function(updatedImage) {
			})
		
		}
	}
	initTextInputValidationProperties(validationParams)
	
	var clearValueParams = {
		initialVal: imageCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: imageCol.parentTableID,
				imageID: imageCol.imageID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/image/setClearValueSupported",formatParams,function(updatedImage) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: imageCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: imageCol.parentTableID,
				imageID: imageCol.imageID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/image/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
}

function initImageColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		imageID: columnID
	}
	jsonAPIRequest("tableView/image/get", getColParams, function(imageCol) { 
		initImageColPropertiesImpl(imageCol)
	})
	
	
	
}