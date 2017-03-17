
// Definition of parameters and callbacks for a date picker to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormImage() {
	console.log("Init html editor design form behavior")
	initNewImageDialog()
}

function selectFormImage($container,attachmentObjRef) {
	console.log("Selected attachment form component: " + JSON.stringify(attachmentObjRef))
	loadImageProperties($container,attachmentObjRef)
}


function resizeImage($container,geometry) {
	
	var imageRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		imageID: imageRef.imageID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/image/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.imageID)
	})	
}


var imageDesignFormConfig = {
	draggableHTMLFunc:	imageContainerHTML,
	startPaletteDrag: function(placeholderID,$paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewImageDialog,
	resizeConstraints: elemResizeConstraints(200,1280,125,1280),
	resizeHandles: 'e,s,se',
	resizeFunc: resizeImage,
	initFunc: initDesignFormImage,
	selectionFunc: selectFormImage
}
