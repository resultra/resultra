
// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript file needs to included after the other text box related files, so all the functions
// are already defined.


function initDesignFormFile() {
	initNewFileDialog()
}

function selectFormFile ($container,fileRef) {
	loadFileProperties($container,fileRef)
}

function resizeFileComponent($container,geometry) {
	
	var fileRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		fileID: fileRef.fileID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/file/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.fileID)
	})	
}


var fileDesignFormConfig = {
	draggableHTMLFunc:	fileContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewFileDialog,
	resizeConstraints: elemResizeConstraints(75,600,400,400),
	resizeFunc: resizeFileComponent,
	initFunc: initDesignFormFile,
	selectionFunc: selectFormFile
}
