
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
	resizeConstraints: elemResizeConstraints(100,640,30,30),
	resizeFunc: resizeSocialButton,
	initFunc: initDesignFormSocialButton,
	selectionFunc: selectFormSocialButton
}
