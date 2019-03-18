// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Definition of parameters and callbacks for a date picker to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormImage() {
	console.log("Init html editor design form behavior")
	initNewAttachmentDialog()
}

function selectFormAttachment($container,attachmentObjRef) {
	console.log("Selected attachment form component: " + JSON.stringify(attachmentObjRef))
	loadAttachmentProperties($container,attachmentObjRef)
}


function resizeImage($container,geometry) {
	
	var imageRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		imageID: imageRef.imageID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/attachment/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.imageID)
	})	
}


var attachmentDesignFormConfig = {
	draggableHTMLFunc:	attachmentContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewAttachmentDialog,
	resizeConstraints: elemResizeConstraints(200,1280,125,1280),
	resizeHandles: 'e,s,se',
	resizeFunc: resizeImage,
	initFunc: initDesignFormImage,
	selectionFunc: selectFormAttachment
}
