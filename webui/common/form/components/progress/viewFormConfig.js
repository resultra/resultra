function loadRecordIntoProgressIndicator($progressContainer, recordRef) {
	
	console.log("loadRecordIntoProgressIndicator: loading record into progress indicator: " + JSON.stringify(recordRef))
	
	var progressObjectRef = getContainerObjectRef($progressContainer)
	
	var progressFieldID = progressObjectRef.properties.fieldID
	
	
	function setProgressVal(val) {
		var $progressBarContainer = $progressContainer.find(".progress")
	
		$progressBarContainer.find(".progress-bar").each(function() {
			var $progressBar = $(this)
			var barRange = $progressBar.data("barRange")
			
			if (val < barRange.minVal) {
				$progressBar.css('width','0%')
			} else if (val >= barRange.maxVal) {
				$progressBar.css('width',barRange.maxPerc + "%")
			} else {
				var partialProgressPerc = (val-barRange.minVal)/(barRange.maxVal-barRange.minVal)*barRange.maxPerc
				$progressBar.css('width',partialProgressPerc + "%")
			}
		})
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
	
	
	
	function populateStackedProgressBars($progressBarContainer,minVal,maxVal,thresholdVals) {
		
		// make sure the given threshold values are sorted. The algorithm to populate the 
		// stacked progress bars depends on it.
		thresholdVals.sort(function(a,b) { return a.startingVal-b.startingVal })
		
		// Populate one stacked progress bar in the parent DOM element $progressBarContainer
		function populateOneStackedProgressBar($progressBarContainer,minBarVal,maxBarVal,colorScheme) {
			var $progressBar = $('<div class="progress-bar" role="progressbar">' +
    				'<span class="sr-only"></span>' +
				'</div>')
			$progressBar.addClass("progress-bar-"+colorScheme)
			$progressBar.css('width',"0%")
			var barRangeInfo = {
				minVal: minBarVal,
				maxVal: maxBarVal,
				maxPerc: (maxBarVal-minBarVal)/(maxVal-minVal)*100.0
			}
			$progressBar.data("barRange",barRangeInfo)
			
			$progressBarContainer.append($progressBar)
		}
		
		// Given the index of the current threshold value, inspect
		// the next threshold value's starting value. The next
		// threshold value's starting value, or the maximum value for
		// the progress bar as a whole serves the the maximum value for the
		// current threshold value.
		function currThresholdEndVal(currIndex) {
			var nextIndex = currIndex+1
			if (nextIndex >= thresholdVals.length) {
				return maxVal
			} else {
				var nextVal = thresholdVals[nextIndex]
				if(nextVal.startingVal >= maxVal) {
					return maxVal
				} else {
					return nextVal.startingVal
				}
			}
		}
		
		// Replace the boiler-plate progress bar with one or more actual progress bars, depending
		// on the range and thresholds.
		$progressBarContainer.empty()
		if(thresholdVals.length == 0) {
			// If no threshold values are defined, then populate a single (stacked) progress
			// bar encompassing the entire value range of the progress indicator.
			populateOneStackedProgressBar($progressBarContainer,minVal,maxVal,"info")
		} else {
			var firstThreshold = thresholdVals[0]
			if(firstThreshold.startingVal > minVal) {
				// If the first threshold starts after the minimum value for the progres bar as a whole,
				// Fill in the remainder with an "info" stacked progress bar.
				var firstBarEndVal = currThresholdEndVal(-1)
				populateOneStackedProgressBar($progressBarContainer,minVal,firstBarEndVal,"info")
			}
			for(var thresholdIndex = 0; thresholdIndex < thresholdVals.length; thresholdIndex++) {
				var currThreshold = thresholdVals[thresholdIndex]
				if(currThreshold.startingVal < maxVal) {
					var currBarEndVal = currThresholdEndVal(thresholdIndex)
					populateOneStackedProgressBar($progressBarContainer,currThreshold.startingVal,currBarEndVal,currThreshold.colorScheme)		
				}
				
			}
		}
	}
	var $progressBarContainer = $progress.find(".progress")
	populateStackedProgressBars($progressBarContainer,
			progressObjectRef.properties.minVal,progressObjectRef.properties.maxVal,
			progressObjectRef.properties.thresholdVals)
	
	
	
}