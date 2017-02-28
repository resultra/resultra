// Definition of parameters and callbacks for a progess indicator to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormProgress() {
	console.log("Init progress indicator design form behavior")
}

function selectFormProgress(progressObjRef) {
	console.log("Selected progress indicator: " + JSON.stringify(progressObjRef))
	loadProgressProperties(progressObjRef)
}

function resizeProgress(progressID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		progressID: progressID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/progress/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(progressID,updatedObjRef)
	})	
}


var progressDesignFormConfig = {
	draggableHTMLFunc:	progressContainerHTML,
	startPaletteDrag: function(placeholderID,$paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewCheckboxDialog,
	resizeConstraints: elemResizeConstraints(100,640,30,30),
	resizeFunc: resizeProgress,
	initFunc: initDesignFormProgress,
	selectionFunc: selectFormProgress
}
