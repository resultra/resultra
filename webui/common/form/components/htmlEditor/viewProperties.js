// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initHTMLEditorViewProperties(componentSelectionParams) {
	console.log("Init html editor properties panel")

	var htmlEditorRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()	

	var elemPrefix = "htmlEditor_"

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		fieldID: htmlEditorRef.properties.fieldID
	}
	initFormComponentTimelinePane(timelineParams)

	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#htmlEditorViewProps')
	
	
}