// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
// Definition of parameters and callbacks for a progess indicator to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormProgress() {
	console.log("Init progress indicator design form behavior")
}

function selectFormProgress($container,progressObjRef) {
	console.log("Selected progress indicator: " + JSON.stringify(progressObjRef))
	loadProgressProperties($container,progressObjRef)
}

function resizeProgress($container,geometry) {
	
	var progressRef = getContainerObjectRef($container)
	 
	var resizeParams = {
		parentFormID: designFormContext.formID,
		progressID: progressRef.progressID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/progress/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.progressID)
	})	
}


var progressDesignFormConfig = {
	draggableHTMLFunc:	progressContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewProgressDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(75,640),
	resizeFunc: resizeProgress,
	initFunc: initDesignFormProgress,
	selectionFunc: selectFormProgress
}
