// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Definition of parameters and callbacks for a toggle control to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormToggle() {
	console.log("Init toggle design form behavior")
	initNewToggleDialog()
}

function selectFormToggle($container,toggleObjRef) {
	console.log("Selected toggle: " + JSON.stringify(toggleObjRef))
	loadToggleProperties($container,toggleObjRef)
}

function resizeToggle($container,geometry) {
	
	var toggleRef = getContainerObjectRef($container)
	
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		toggleID: toggleRef.toggleID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/toggle/resize", resizeParams, function(updatedObjRef) {
		reInitToggleComponentControl($container,updatedObjRef)
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.toggleID)
	})	
}


var toggleDesignFormConfig = {
	draggableHTMLFunc:	toggleContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {
		initDummyToggleControlForDragAndDrop($paletteItemContainer)
	},
	createNewItemAfterDropFunc: openNewToggleDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(100,640),
	resizeFunc: resizeToggle,
	initFunc: initDesignFormToggle,
	selectionFunc: selectFormToggle
}
