
function initRatingDesignControlBehavior(ratingObjectRef) {
	var ratingControlSelector = '#' + ratingControlIDFromElemID(ratingObjectRef.ratingID)

	$(ratingControlSelector).rating()
	
}


function initDesignFormRating() {
	initNewRatingDialog()
}

function selectFormRating(ratingObjRef) {
	loadRatingProperties(ratingObjRef)
}

function resizeRating(ratingID,geometry) {
	var resizeParams = {
		parentFormID: ratingObjectRef.parentFormID,
		ratingID: ratingID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/rating/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(ratingID,updatedObjRef)
	})	
}

function startRatingPaletteDrag(placeholderID) {
	var ratingControlSelector = '#'+ratingControlIDFromElemID(placeholderID)
	$(ratingControlSelector).rating()
}


var ratingDesignFormConfig = {
	draggableHTMLFunc:	ratingContainerHTML,
	startPaletteDrag: startRatingPaletteDrag,
	createNewItemAfterDropFunc: openNewRatingDialog,
	resizeConstraints: elemResizeConstraints(100,640,30,30),
	resizeFunc: resizeRating,
	initFunc: initDesignFormRating,
	selectionFunc: selectFormRating
}
