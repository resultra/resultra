function initNumberInputColPropertiesImpl(numberInputCol) {
	
	setColPropsHeader(numberInputCol)
	
	var elemPrefix = "numberInput_"
	
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