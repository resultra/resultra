
// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript image needs to included after the other text box related images, so all the functions
// are already defined.


function initDesignFormImage() {
	initNewImageDialog()
}

function selectFormImage ($container,imageRef) {
	loadImageProperties($container,imageRef)
}

function resizeImageComponent($container,geometry) {
	
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
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewImageDialog,
	resizeConstraints: elemResizeConstraints(75,600,400,400),
	resizeFunc: resizeImageComponent,
	initFunc: initDesignFormImage,
	selectionFunc: selectFormImage
}
