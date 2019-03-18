// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initCommentViewProperties(componentSelectionParams) {
	console.log("Init comment component properties panel")
	
	var commentRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()
	
	// Comment boxes are not linked to the timeline with a ComponentLink,
	// However, a ComponentLink can be synthesized with just a field ID.
	var elemPrefix = "comment_"

	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		fieldID: commentRef.properties.fieldID
	}
	initFormComponentTimelinePane(timelineParams)
	
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#commentViewProps')
	
	
}