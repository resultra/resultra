var colorThresholdSchemeDefault = 'default'

// Given a list of threshold values, convert these thresholds into zones for each color scheme. This is useful drawing 
// controls from a minimum and maximum threshold value, rather than just a minimum. 
function convertStartingThresholdsToZones(thresholdVals, minVal, maxVal) {
	
	var zones = []
	
	
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
	
	if(thresholdVals.length == 0) {
		var zone = { min:minVal,max:maxVal,colorScheme:"none" }
		zones.push(zone)
	} else {
		var firstThreshold = thresholdVals[0]
		if(firstThreshold.startingVal > minVal) {
			// If the first threshold starts after the minimum value,
			// Fill the remainder with the "none" color scheme.
			var firstBarEndVal = currThresholdEndVal(-1)
			var zone = { min: minVal, max: firstBarEndVal, colorScheme: "none" }
			zones.push(zone)
		}
		for(var thresholdIndex = 0; thresholdIndex < thresholdVals.length; thresholdIndex++) {
			var currThreshold = thresholdVals[thresholdIndex]
			if(currThreshold.startingVal < maxVal) {
				var currBarEndVal = currThresholdEndVal(thresholdIndex)
				var zone = { min: currThreshold.startingVal, max: currBarEndVal, colorScheme:currThreshold.colorScheme }
				zones.push(zone)
			}
			
		}
	}
	
	return zones
	
}