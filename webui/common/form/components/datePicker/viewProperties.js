// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initDatePickerViewProperties(componentSelectionParams) {
	console.log("Init date picker properties panel")
	
	var datePickerRef = componentSelectionParams.selectedObjRef
	var currRecordRef = componentSelectionParams.getCurrentRecordFunc()	

	var elemPrefix = "datePicker_"
	
	var timelineParams = {
		elemPrefix: elemPrefix,
		recordID: currRecordRef.recordID,
		fieldID: datePickerRef.properties.fieldID
	}
	initFormComponentTimelinePane(timelineParams)
	
	// Toggle to the check box properties, hiding the other property panels
	hideSiblingsShowOne('#datePickerViewProps')
	
	
}