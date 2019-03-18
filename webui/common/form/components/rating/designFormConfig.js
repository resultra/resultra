// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

function initRatingDesignControlBehavior($ratingContainer,ratingObjectRef) {
	
	
	var $ratingControl = getRatingControlFromRatingContainer($ratingContainer)

	$ratingControl.rating()
	
}


function initDesignFormRating() {
	initNewRatingDialog()
}

function selectFormRating($container,ratingObjRef) {
	loadRatingProperties($container,ratingObjRef)
}

function resizeRating($container,geometry) {
	
	var ratingRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: ratingRef.parentFormID,
		ratingID: ratingRef.ratingID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/rating/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.ratingID)
	})	
}

function startRatingPaletteDrag($paletteItemContainer) {
	
	var $ratingControl = getRatingControlFromRatingContainer($paletteItemContainer)
	 $ratingControl.rating()
}


var ratingDesignFormConfig = {
	draggableHTMLFunc:	ratingContainerHTML,
	initDummyDragAndDropComponentContainer: startRatingPaletteDrag,
	createNewItemAfterDropFunc: openNewRatingDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(100,640),
	resizeFunc: resizeRating,
	initFunc: initDesignFormRating,
	selectionFunc: selectFormRating
}
