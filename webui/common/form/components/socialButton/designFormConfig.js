// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function initSocialButtonDesignControlBehavior($socialButtonContainer,socialButtonObjectRef) {
	var $socialButtonControl = getSocialButtonControlFromContainer($socialButtonContainer)
}


function initDesignFormSocialButton() {
	initNewSocialButtonDialog()
}

function selectFormSocialButton($container,socialButtonObjRef) {
	loadSocialButtonProperties($container,socialButtonObjRef)
}

function resizeSocialButton($container,geometry) {
	
	var socialButtonRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: socialButtonRef.parentFormID,
		socialButtonID: socialButtonRef.socialButtonID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/socialButton/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.socialButtonID)
	})	
}

function startSocialButtonPaletteDrag($paletteItemContainer) {
	initDummySocialButtonFormComponentControl($paletteItemContainer)
}


var socialButtonDesignFormConfig = {
	draggableHTMLFunc:	socialButtonContainerHTML,
	initDummyDragAndDropComponentContainer: startSocialButtonPaletteDrag,
	createNewItemAfterDropFunc: openNewSocialButtonDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(100,640),
	resizeFunc: resizeSocialButton,
	initFunc: initDesignFormSocialButton,
	selectionFunc: selectFormSocialButton
}
