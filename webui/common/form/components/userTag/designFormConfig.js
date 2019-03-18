// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function initUserTagDesignControlBehavior(userTagObjectRef) {
// no-op	
}


function initDesignFormUserTag() {
	initUserTagDialog()
}

function selectFormUserTag($container,userTagObjectRef) {
	loadUserTagProperties($container,userTagObjectRef)
}


function resizeUserTag($container,geometry) {
	
	var userSelRef = getContainerObjectRef($container)
	
	initDummiedUpUserTagControl($container,geometry.sizeWidth)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		userTagID: userSelRef.userTagID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/userTag/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.userTagID)
	})	
}

function startUserTagPaletteDrag($paletteItemContainer) {
	initDummiedUpUserTagControl($paletteItemContainer,250)
}


var userTagDesignFormConfig = {
	draggableHTMLFunc:	userTagContainerHTML,
	initDummyDragAndDropComponentContainer: startUserTagPaletteDrag,
	createNewItemAfterDropFunc: openNewUserTagDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(200,1200),
	resizeFunc: resizeUserTag,
	initFunc: initDesignFormUserTag,
	selectionFunc: selectFormUserTag
}
