// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Definition of parameters and callbacks for a date picker to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormDatePicker() {
	console.log("Init checkbox design form behavior")
	initNewDatePickerDialog()
}

function selectFormDatePicker($container,datePickerObjRef) {
	console.log("Selected date picker: " + JSON.stringify(datePickerObjRef))
	loadDatePickerProperties($container,datePickerObjRef)
}


function resizeDatePicker($container,geometry) {
	
	var datePickerRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		datePickerID: datePickerRef.datePickerID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/datePicker/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.datePickerID)
	})	
}


var datePickerDesignFormConfig = {
	draggableHTMLFunc:	datePickerContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewDatePickerDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(125,640),
	resizeFunc: resizeDatePicker,
	initFunc: initDesignFormDatePicker,
	selectionFunc: selectFormDatePicker
}
