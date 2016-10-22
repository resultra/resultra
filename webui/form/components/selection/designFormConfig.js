
// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript file needs to included after the other text box related files, so all the functions
// are already defined.


function initDesignFormSelection() {
	console.log("Init text box design form behavior")
	initNewSelectionDialog()
}

function selectFormSelection (selectionRef) {
	console.log("Select selection component: " + JSON.stringify(selectionRef))
	loadSelectionProperties(selectionRef)
}

function resizeSelection(selectionID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		selectionID: selectionID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/selection/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(selectionID,updatedObjRef)
	})	
}


var selectionDesignFormConfig = {
	draggableHTMLFunc:	selectionContainerHTML,
	startPaletteDrag: function(placeholderID) {},
	createNewItemAfterDropFunc: openNewSelectionDialog,
	resizeConstraints: elemResizeConstraints(100,600,400,400),
	resizeFunc: resizeSelection,
	initFunc: initDesignFormSelection,
	selectionFunc: selectFormSelection
}
