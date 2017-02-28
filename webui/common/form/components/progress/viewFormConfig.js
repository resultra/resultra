function loadRecordIntoProgressIndicator($progressContainer, recordRef) {
	
	console.log("loadRecordIntoProgressIndicator: loading record into progress indicator: " + JSON.stringify(recordRef))
	
	var progressObjectRef = getContainerObjectRef($progressContainer)
	
	var progressFieldID = progressObjectRef.properties.fieldID

	
	function setProgressVal(val) {
		var $progressBar = $progressContainer.find(".progress-bar")
		
		var range = progressObjectRef.properties.maxVal - progressObjectRef.properties.minVal
		var progressWidthPerc = (val -  progressObjectRef.properties.minVal)/range * 100.0
		if (progressWidthPerc > 100) {
			progressWidthPerc = 100
		} else if (progressWidthPerc < 0) {
			progressWidthPerc = 0
		}
		
		$progressBar.css('width', progressWidthPerc+'%')
		$progressBar.attr('aria-valuenow', val)
	}
	
	// Populate the "intersection" of field values in the record
	// with the fields shown by the layout's containers.
	if(recordRef.fieldValues.hasOwnProperty(progressFieldID)) {

		var fieldVal = recordRef.fieldValues[progressFieldID]
		setProgressVal(fieldVal)

	} // If record has a value for the current container's associated field ID.
	else
	{
		setProgressVal(0.0)
	}	
	
}

function initProgressRecordEditBehavior($progress,componentContext,recordProxy, progressObjectRef) {
		
	$progress.data("viewFormConfig", {
		loadRecord: loadRecordIntoProgressIndicator
	})	
}