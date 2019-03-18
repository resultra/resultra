// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function loadRecordIntoProgressIndicator($progressContainer, recordRef) {
	
	console.log("loadRecordIntoProgressIndicator: loading record into progress indicator: " + JSON.stringify(recordRef))
	
	var progressObjectRef = getContainerObjectRef($progressContainer)
	
	var progressFieldID = progressObjectRef.properties.fieldID
	
	
	function setProgressVal(val) {
		var $progressBarContainer = $progressContainer.find(".progress")
			
		// Put the formatted value on the first progress bar.
		var $progressBar = $progressBarContainer.find(".progress-bar")
		var formattedVal = formatNumberValue(progressObjectRef.properties.valueFormat.format,val)
		$progressBar.text(formattedVal)
		
		function getProgressBarColorScheme() {
			var thresholdVals = progressObjectRef.properties.thresholdVals
			if(thresholdVals.length > 0) {
				var firstThreshold = thresholdVals[0]
				if(firstThreshold.startingVal > val) {
					// If the first threshold starts after the current value,
					// color the bar with the default.
					return "info"
				}
				var thresholdColorScheme = "info" // default
				for(var thresholdIndex = 0; thresholdIndex < thresholdVals.length; thresholdIndex++) {
					var currThreshold = thresholdVals[thresholdIndex]
					if(currThreshold.startingVal <= val) {
						thresholdColorScheme = currThreshold.colorScheme
					}
				}
				return thresholdColorScheme
			} else {
				return "info"
			}
		}
		
		var progressColorSchemeClass = "progress-bar-" + getProgressBarColorScheme()
		$progressBar.removeClass()
		$progressBar.addClass("progress-bar")		
		$progressBar.addClass(progressColorSchemeClass)
		
		var minVal = progressObjectRef.properties.minVal
		var maxVal = progressObjectRef.properties.maxVal
		var currProgress = ((val - minVal) / (maxVal-minVal)) * 100
		if (currProgress < 0) { currProgress = 0 }
		if (currProgress > 100) { currProgress = 100 }
		var currProgressPerc = currProgress + '%'
		$progressBar.css("width",currProgressPerc)
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