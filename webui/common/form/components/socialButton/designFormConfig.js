
function initSocialButtonDesignControlBehavior($socialButtonContainer,socialButtonObjectRef) {
	
	
	var $socialButtonControl = getSocialButtonControlFromContainer($socialButtonContainer)
	
}


function initDesignFormRating() {
	initNewSocialButtonDialog()
}

function selectFormRating($container,ratingObjRef) {
	loadRatingProperties($container,ratingObjRef)
}

function resizeRating($container,geometry) {
	
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
	
	var $socialButtonControl = getSocialButtonControlFromContainer($paletteItemContainer)
	 $socialButtonControl.rating()
}


var socialButtonDesignFormConfig = {
	draggableHTMLFunc:	socialButtonContainerHTML,
	initDummyDragAndDropComponentContainer: startSocialButtonPaletteDrag,
	createNewItemAfterDropFunc: openNewSocialButtonDialog,
	resizeConstraints: elemResizeConstraints(100,640,30,30),
	resizeFunc: resizeRating,
	initFunc: initDesignFormRating,
	selectionFunc: selectFormRating
}
