// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initFileViewProperties(componentSelectionParams) {
	console.log("Init text box properties panel")

	var fileRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()		

	var elemPrefix = "file_"
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#fileViewProps')
	
	
}