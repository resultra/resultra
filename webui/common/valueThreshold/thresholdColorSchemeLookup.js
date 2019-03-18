// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function getThresholdColorScheme(thresholdVals,val) {
	var colorScheme = colorThresholdSchemeDefault
	if(thresholdVals.length > 0) {
		// make sure the given threshold values are sorted. The algorithm to populate the 
		// stacked progress bars depends on it.
		thresholdVals.sort(function(a,b) { return a.startingVal-b.startingVal })
		
		for(var thresholdIndex in thresholdVals) {
			var currThreshold = thresholdVals[thresholdIndex]
			if(val > currThreshold.startingVal) {
				colorScheme = currThreshold.colorScheme
			}
		}
		return colorScheme
	
	} else {
		return colorScheme
	}		
}