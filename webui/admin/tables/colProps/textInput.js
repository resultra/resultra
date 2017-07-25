function initTextInputColPropertiesImpl(textInputCol) {
	
	setColPropsHeader(textInputCol)
	hideSiblingsShowOne("#textInputColProps")
	
	var elemPrefix = "textInput_"
		
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: textInputCol.parentTableID,
			textInputID: textInputCol.textInputID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/textInput/setLabelFormat", formatParams, function(updatedNumberInput) {
			setColPropsHeader(updatedNumberInput)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: textInputCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: textInputCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: textInputCol.parentTableID,
				textInputID: textInputCol.textInputID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/textInput/setPermissions",params,function(updatedTextInput) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	function saveValueList(valueListID) {
		var setValueListParams = {
			parentTableID: textInputCol.parentTableID,
			textInputID: textInputCol.textInputID,
			valueListID: valueListID
		}
		jsonAPIRequest("tableView/textInput/setValueList", setValueListParams, function(updatedTextBox) {
		})			
	}
	var valueListPropertyParams = {
		databaseID: colPropsAdminContext.databaseID,
		saveValueListCallback: saveValueList,
		defaultValueListID: textInputCol.properties.valueListID
	}
	initValueListSelectionPropertyPanel(valueListPropertyParams)
	
	
	var validationParams = {
		initialValidationProps: textInputCol.properties.validation,
		setValidation: function(validationProps) {
			var validationParams = {
				parentTableID: textInputCol.parentTableID,
				textInputID: textInputCol.textInputID,
				validation: validationProps
			}
			console.log("Setting new validation settings: " + JSON.stringify(validationParams))

			jsonAPIRequest("tableView/textInput/setValidation",validationParams,function(updatedTextInput) {
			})
		
		}
	}
	initTextInputValidationProperties(validationParams)
	
	var clearValueParams = {
		initialVal: textInputCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: textInputCol.parentTableID,
				textInputID: textInputCol.textInputID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/textInput/setClearValueSupported",formatParams,function(updatedTextInput) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: textInputCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: textInputCol.parentTableID,
				textInputID: textInputCol.textInputID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/textInput/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
}

function initTextInputColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		textInputID: columnID
	}
	jsonAPIRequest("tableView/textInput/get", getColParams, function(textInputCol) { 
		initTextInputColPropertiesImpl(textInputCol)
	})
	
	
	
}