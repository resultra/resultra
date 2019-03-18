// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function initUserSelectionDesignControlBehavior(userSelectionObjectRef) {
// no-op	
}


function initDesignFormUserSelection() {
	initUserSelectionDialog()
}

function selectFormUserSelection($container,userSelectionObjectRef) {
	loadUserSelectionProperties($container,userSelectionObjectRef)
}


function resizeUserSelection($container,geometry) {
	
	var userSelRef = getContainerObjectRef($container)
	
	initDummiedUpUserSelectionControl($container,geometry.sizeWidth)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		userSelectionID: userSelRef.userSelectionID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/userSelection/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.userSelectionID)
	})	
}

function startUserSelectionPaletteDrag($paletteItemContainer) {
	initDummiedUpUserSelectionControl($paletteItemContainer,250)
}


var userSelectionDesignFormConfig = {
	draggableHTMLFunc:	userSelectionContainerHTML,
	initDummyDragAndDropComponentContainer: startUserSelectionPaletteDrag,
	createNewItemAfterDropFunc: openNewUserSelectionDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(200,1200),
	resizeFunc: resizeUserSelection,
	initFunc: initDesignFormUserSelection,
	selectionFunc: selectFormUserSelection
}
