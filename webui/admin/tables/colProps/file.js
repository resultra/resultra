function initFileColPropertiesImpl(fileCol) {
	
	setColPropsHeader(fileCol)
	hideSiblingsShowOne("#fileColProps")
	
	var elemPrefix = "file_"
		
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: fileCol.parentTableID,
			fileID: fileCol.fileID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/file/setLabelFormat", formatParams, function(updatedNumberInput) {
			setColPropsHeader(updatedNumberInput)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: fileCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: fileCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: fileCol.parentTableID,
				fileID: fileCol.fileID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/file/setPermissions",params,function(updatedFile) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
		
	
	var validationParams = {
		initialValidationProps: fileCol.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentTableID: fileCol.parentTableID,
				fileID: fileCol.fileID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("tableView/file/setValidation",validationParams,function(updatedFile) {
			})
		
		}
	}
	initTextInputValidationProperties(validationParams)
	
	var clearValueParams = {
		initialVal: fileCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: fileCol.parentTableID,
				fileID: fileCol.fileID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/file/setClearValueSupported",formatParams,function(updatedFile) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: fileCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: fileCol.parentTableID,
				fileID: fileCol.fileID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/file/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
}

function initFileColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		fileID: columnID
	}
	jsonAPIRequest("tableView/file/get", getColParams, function(fileCol) { 
		initFileColPropertiesImpl(fileCol)
	})
	
	
	
}