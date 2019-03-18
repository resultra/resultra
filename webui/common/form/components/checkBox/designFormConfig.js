// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Definition of parameters and callbacks for a check box to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormCheckBox() {
	console.log("Init checkbox design form behavior")
	initNewCheckBoxDialog()
}

function selectFormCheckbox($container,checkboxObjRef) {
	console.log("Selected checkbox: " + JSON.stringify(checkboxObjRef))
	loadCheckboxProperties($container,checkboxObjRef)
}

function resizeCheckBox($container,geometry) {
	
	var checkboxRef = getContainerObjectRef($container)
	
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		checkBoxID: checkboxRef.checkBoxID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/checkBox/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.checkBoxID)
	})	
}


var checkBoxDesignFormConfig = {
	draggableHTMLFunc:	checkBoxContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewCheckboxDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(100,640),
	resizeFunc: resizeCheckBox,
	initFunc: initDesignFormCheckBox,
	selectionFunc: selectFormCheckbox
}
