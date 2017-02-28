
function initDesignFormCaption() {
	console.log("Init caption design form behavior")
//	initNewCheckBoxDialog()
}

function selectFormCaption(captionObjRef) {
	console.log("Selected caption: " + JSON.stringify(captionObjRef))
	loadFormCaptionProperties(captionObjRef)
}

function resizeFormCaption(captionID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		captionID: captionID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/caption/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(captionID,updatedObjRef)
	})	
}

var formCaptionDesignFormConfig = {
	draggableHTMLFunc:	formCaptionContainerHTML,
	startPaletteDrag: function(placeholderID,$paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewFormCaptionDialog,
	resizeConstraints: elemResizeConstraints(320,640,50,50),
	resizeFunc: resizeFormCaption,
	initFunc: initDesignFormCaption,
	selectionFunc: selectFormCaption
}
