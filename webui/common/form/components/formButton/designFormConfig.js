// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function initDesignFormButton() {
	console.log("Init button design form behavior")
}

function selectFormButton($container,buttonObjRef) {
	console.log("Selected button: " + JSON.stringify(buttonObjRef))
	loadFormButtonProperties($container,buttonObjRef)
}

function resizeFormButton($container,geometry) {
	
	var buttonRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		buttonID: buttonRef.buttonID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/formButton/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.buttonID)
	})	
}

var formButtonDesignFormConfig = {
	draggableHTMLFunc:	formButtonContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewFormButtonDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(50,640),
	resizeFunc: resizeFormButton,
	initFunc: initDesignFormButton,
	selectionFunc: selectFormButton
}
