
// Definition of parameters and callbacks for a toggle control to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormToggle() {
	console.log("Init toggle design form behavior")
	initNewToggleDialog()
}

function selectFormToggle($container,toggleObjRef) {
	console.log("Selected toggle: " + JSON.stringify(toggleObjRef))
	loadToggleProperties($container,toggleObjRef)
}

function resizeToggle($container,geometry) {
	
	var toggleRef = getContainerObjectRef($container)
	
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		toggleID: toggleRef.toggleID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/toggle/resize", resizeParams, function(updatedObjRef) {
		reInitToggleComponentControl($container,updatedObjRef)
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.toggleID)
	})	
}


var toggleDesignFormConfig = {
	draggableHTMLFunc:	toggleContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {
		initDummyToggleControlForDragAndDrop($paletteItemContainer)
	},
	createNewItemAfterDropFunc: openNewToggleDialog,
	resizeConstraints: elemResizeConstraints(100,640,30,30),
	resizeFunc: resizeToggle,
	initFunc: initDesignFormToggle,
	selectionFunc: selectFormToggle
}
