// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function initLabelDesignControlBehavior(labelObjectRef) {
// no-op	
}


function initDesignFormLabel() {
	initLabelDialog()
}

function selectFormLabel($container,labelObjectRef) {
	loadLabelProperties($container,labelObjectRef)
}

function resizeLabel($container,geometry) {
	
	var userSelRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		labelID: userSelRef.labelID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/label/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.labelID)
		var labelWidth = updatedObjRef.properties.geometry.sizeWidth - 15
		initLabelSelectionControl($container, updatedObjRef,labelWidth)
	})	
}

function startLabelPaletteDrag($paletteItemContainer) {
	initLabelSelectionControl($paletteItemContainer)
}


var labelDesignFormConfig = {
	draggableHTMLFunc:	labelContainerHTML,
	initDummyDragAndDropComponentContainer: startLabelPaletteDrag,
	createNewItemAfterDropFunc: openNewLabelDialog,
	resizeConstraints: elemResizeConstraints(200,800,75,400),
	resizeFunc: resizeLabel,
	initFunc: initDesignFormLabel,
	selectionFunc: selectFormLabel
}
