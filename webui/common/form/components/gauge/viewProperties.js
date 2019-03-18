// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initGaugeViewProperties(componentSelectionParams) {
	console.log("Init gauge properties panel")
		
	var gaugeRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()	

	var elemPrefix = "gauge_"		
	
	// Toggle to the gauges properties, hiding the other property panels
	hideSiblingsShowOne('#gaugeViewProps')
	
	
}