
function initRatingDesignControlBehavior($ratingContainer,ratingObjectRef) {
	
	
	var $ratingControl = getRatingControlFromRatingContainer($ratingContainer)

	$ratingControl.rating()
	
}


function initDesignFormRating() {
	initNewRatingDialog()
}

function selectFormRating(ratingObjRef) {
	loadRatingProperties(ratingObjRef)
}

function resizeRating($container,geometry) {
	
	var ratingRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: ratingObjectRef.parentFormID,
		ratingID: ratingRef.ratingID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/rating/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.ratingID)
	})	
}

function startRatingPaletteDrag(placeholderID,$paletteItemContainer) {
	
	var $ratingControl = getRatingControlFromRatingContainer($paletteItemContainer)
	 $ratingControl.rating()
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
