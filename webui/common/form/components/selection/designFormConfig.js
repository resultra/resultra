// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript file needs to included after the other text box related files, so all the functions
// are already defined.


function initDesignFormSelection() {
	console.log("Init text box design form behavior")
	initNewSelectionDialog()
}

function selectFormSelection ($container,selectionRef) {
	console.log("Select selection component: " + JSON.stringify(selectionRef))
	loadSelectionProperties($container,selectionRef)
}

function resizeSelection($container,geometry) {
	
	var selectionRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		selectionID: selectionRef.selectionID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/selection/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.selectionID)
	})	
}


var selectionDesignFormConfig = {
	draggableHTMLFunc:	selectionContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewSelectionDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(100,600),
	resizeFunc: resizeSelection,
	initFunc: initDesignFormSelection,
	selectionFunc: selectFormSelection
}
