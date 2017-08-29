
function initProgressColPropertiesImpl(progressCol) {
	
	setColPropsHeader(progressCol)
	
	var elemPrefix = "progress_"
	hideSiblingsShowOne("#progressColProps")
	
	
	function saveLabelProps(updatedLabelProps) {
		console.log("Saving label propeties for text box")
		var formatParams = {
			parentTableID: progressCol.parentTableID,
			progressID: progressCol.progressID,
			labelFormat: updatedLabelProps
		}
		jsonAPIRequest("tableView/progress/setLabelFormat", formatParams, function(updatedProgress) {
			setColPropsHeader(updatedProgress)
		})	
	}
	var labelParams = {
		elemPrefix: elemPrefix,
		initialVal: progressCol.properties.labelFormat,
		saveLabelPropsCallback: saveLabelProps
	}
	initComponentLabelPropertyPanel(labelParams)


	var formatSelectionParams = {
		elemPrefix: elemPrefix,
		initialFormat: progressCol.properties.valueFormat.format,
		formatChangedCallback: function (newFormat) {
			console.log("Format changed for progress bar: " + newFormat)
			
			var newValueFormat = {
				format: newFormat
			}
			var formatParams = {
				parentTableID: progressCol.parentTableID,
				progressID: progressCol.progressID,
				valueFormat: newValueFormat
			}
			jsonAPIRequest("tableView/progress/setValueFormat", formatParams, function(updatedProgress) {
			})	
			
		}
	}
	initNumberFormatSelection(formatSelectionParams)
	
	
	function saveProgressThresholds(newThresholdVals) {
		var setThresholdParams = {
			parentTableID: progressCol.parentTableID,
			progressID: progressCol.progressID,
			thresholdVals: newThresholdVals
		}
		jsonAPIRequest("tableView/progress/setThresholds", setThresholdParams, function(updatedProgress) {
		})	
	}	
	var thresholdParams = {
		elemPrefix: elemPrefix,
		saveThresholdsCallback: saveProgressThresholds,
		initialThresholdVals: progressCol.properties.thresholdVals
	}
	initThresholdValuesPropertyPanel(thresholdParams)
	
	
	function setProgressRange(minVal,maxVal) {
		var setRangeParams = {
			parentTableID: progressCol.parentTableID,
			progressID: progressCol.progressID,
			minVal: minVal,
			maxVal: maxVal
		}
		jsonAPIRequest("tableView/progress/setRange", setRangeParams, function(updatedProgress) {
		})	
	}
	var progressRangeParams = {
		setRangeCallback: setProgressRange,
		initialMinVal: progressCol.properties.minVal,
		initialMaxVal: progressCol.properties.maxVal
	}
	initProgressRangeProperties(progressRangeParams)
	
	
	
	var helpPopupParams = {
		initialMsg: progressCol.properties.helpPopupMsg,
		elemPrefix: elemPrefix,	
		setMsg: function(popupMsg) {
			var params = {
				parentTableID: progressCol.parentTableID,
				progressID: progressCol.progressID,
				popupMsg: popupMsg
			}
			jsonAPIRequest("tableView/progress/setHelpPopupMsg",params,function(updatedProgress) {
			})
		}	
	}
	initComponentHelpPopupPropertyPanel(helpPopupParams)
	
	
}


function initProgressColProperties(tableID,columnID) {
	
	var getColParams = {
		parentTableID: tableID,
		progressID: columnID
	}
	jsonAPIRequest("tableView/progress/get", getColParams, function(progressCol) { 
		initProgressColPropertiesImpl(progressCol)
	})
}