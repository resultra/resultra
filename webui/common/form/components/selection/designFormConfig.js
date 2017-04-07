
// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript file needs to included after the other text box related files, so all the functions
// are already defined.


function initDesignFormSelection() {
	console.log("Init text box design form behavior")
	initNewSelectionDialog()
}

function selectFormSelection ($container,selectionRef) {
	console.log("Select selection component: " + JSON.stringify(selectionRef))
	loadSelectionProperties($container,selectionRef)
}

function resizeSelection($container,geometry) {
	
	var selectionRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		selectionID: selectionRef.selectionID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/selection/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.selectionID)
	})	
}


var selectionDesignFormConfig = {
	draggableHTMLFunc:	selectionContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewSelectionDialog,
	resizeConstraints: elemResizeConstraints(100,600,400,400),
	resizeFunc: resizeSelection,
	initFunc: initDesignFormSelection,
	selectionFunc: selectFormSelection
}
