
// Definition of parameters and callbacks for a date picker to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormImage() {
	console.log("Init html editor design form behavior")
	initNewImageDialog()
}

function selectFormImage(imageObjRef) {
	console.log("Selected html editor: " + JSON.stringify(imageObjRef))
	loadImageProperties(imageObjRef)
}


function resizeImage(imageID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		imageID: imageID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/image/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(imageID,updatedObjRef)
	})	
}


var imageDesignFormConfig = {
	draggableHTMLFunc:	imageContainerHTML,
	createNewItemAfterDropFunc: openNewImageDialog,
	resizeConstraints: elemResizeConstraints(100,640,100,500),
	resizeFunc: resizeImage,
	initFunc: initDesignFormImage,
	selectionFunc: selectFormImage
}
