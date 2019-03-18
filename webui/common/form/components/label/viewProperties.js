// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initLabelViewProperties(componentSelectionParams) {
	
	console.log("Init user selection properties panel")
	
	var labelRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()

	var elemPrefix = "label_"

	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#labelViewProps')
	
	
}