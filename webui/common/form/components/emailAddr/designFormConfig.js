// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript file needs to included after the other text box related files, so all the functions
// are already defined.


function initDesignFormEmailAddr() {
	initNewEmailAddrDialog()
}

function selectFormEmailAddr ($container,emailAddrRef) {
	loadEmailAddrProperties($container,emailAddrRef)
}

function resizeEmailAddr($container,geometry) {
	
	var emailAddrRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		emailAddrID: emailAddrRef.emailAddrID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/emailAddr/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.emailAddrID)
	})	
}


var emailAddrDesignFormConfig = {
	draggableHTMLFunc:	emailAddrContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewEmailAddrDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(75,600),
	resizeFunc: resizeEmailAddr,
	initFunc: initDesignFormEmailAddr,
	selectionFunc: selectFormEmailAddr
}
