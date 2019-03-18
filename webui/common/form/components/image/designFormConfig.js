// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript image needs to included after the other text box related images, so all the functions
// are already defined.


function initDesignFormImage() {
	initNewImageDialog()
}

function selectFormImage ($container,imageRef) {
	loadImageProperties($container,imageRef)
}

function resizeImageComponent($container,geometry) {
	
	var imageRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		imageID: imageRef.imageID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/image/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.imageID)
	})	
}


var imageDesignFormConfig = {
	draggableHTMLFunc:	imageContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewImageDialog,
	resizeConstraints: elemResizeConstraints(100,1280,100,1280),
	resizeFunc: resizeImageComponent,
	resizeHandles: 'e,s,se',
	initFunc: initDesignFormImage,
	selectionFunc: selectFormImage
}
