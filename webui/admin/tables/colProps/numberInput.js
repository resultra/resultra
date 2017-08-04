function initNumberInputColPropertiesImpl(numberInputCol) {
	
	setColPropsHeader(numberInputCol)
	
	var elemPrefix = "numberInput_"
	hideSiblingsShowOne("#numberInputColProps")
	
	var formatSelectionParams = {
		elemPrefix: elemPrefix,
		initialFormat: numberInputCol.properties.valueFormat.format,
		formatChangedCallback: function (newFormat) {
		
			console.log("Format changed for text box: " + newFormat)
		
			var newValueFormat = {
				format: newFormat
			}
			var formatParams = {
				parentTableID: numberInputCol.parentTableID,
				numberInputID: numberInputCol.numberInputID,
				valueFormat: newValueFormat
			}
			jsonAPIRequest("tableView/numberInput/setValueFormat", formatParams, function(updatedNumberInput) { 
			})	
		
		}
	}
	initNumberFormatSelection(formatSelectionParams)
	
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: numberInputCol.parentTableID,
			numberInputID: numberInputCol.numberInputID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/numberInput/setLabelFormat", formatParams, function(updatedNumberInput) {
			setColPropsHeader(updatedNumberInput)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: numberInputCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)
	
	
	var readOnlyParams = {
		elemPrefix: elemPrefix,
		initialVal: numberInputCol.properties.permissions,
		permissionsChangedCallback: function(updatedPermissions) {
			var params = {
				parentTableID: numberInputCol.parentTableID,
				numberInputID: numberInputCol.numberInputID,
				permissions: updatedPermissions
			}
			jsonAPIRequest("tableView/numberInput/setPermissions",params,function(updatedNumberInput) {
			})
		}
	}
	initFormComponentPermissionsPropertyPanel(readOnlyParams)
	
	var clearValueParams = {
		initialVal: numberInputCol.properties.clearValueSupported,
		elemPrefix: elemPrefix,
		setClearValueSupported: function(clearValueSupported) {
			var formatParams = {
				parentTableID: numberInputCol.parentTableID,
				numberInputID: numberInputCol.numberInputID,
				clearValueSupported: clearValueSupported
			}
			jsonAPIRequest("tableView/numberInput/setClearValueSupported",formatParams,function(updatedNumberInput) {
			})
		}
	}
	initClearValueProps(clearValueParams)
	
	var helpPopupParams = {
		initialMsg: numberInputCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: numberInputCol.parentTableID,
				numberInputID: numberInputCol.numberInputID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/numberInput/setHelpPopupMsg",params,function(updateCol) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	var conditionalFormatParams = {
		initialFormats: numberInputCol.properties.conditionalFormats,
		setConditionalFormats: function(formats) {
			var params = {
				parentTableID: numberInputCol.parentTableID,
				numberInputID: numberInputCol.numberInputID,
				conditionalFormats: formats
			}
			jsonAPIRequest("tableView/numberInput/setConditionalFormats",params,function(updateCol) {
			})	
		}
	}
	initNumberConditionalFormatPropertyPanel(conditionalFormatParams)
	
	
	
}

function initNumberInputColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		numberInputID: columnID
	}
	jsonAPIRequest("tableView/numberInput/get", getColParams, function(numberInputCol) { 
		initNumberInputColPropertiesImpl(numberInputCol)
	})
	
	
	
}