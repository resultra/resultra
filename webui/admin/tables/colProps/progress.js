
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