
function initUserSelectionDesignControlBehavior(userSelectionObjectRef) {
// no-op	
}


function initDesignFormUserSelection() {
	initUserSelectionDialog()
}

function selectFormUserSelection(userSelectionObjectRef) {
	loadUserSelectionProperties(userSelectionObjectRef)
}

function resizeUserSelection(userSelectionID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		userSelectionID: userSelectionID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/userSelection/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(userSelectionID,updatedObjRef)
	})	
}

function startUserSelectionPaletteDrag(placeholderID,$paletteItemContainer) {
// no-op
}


var userSelectionDesignFormConfig = {
	draggableHTMLFunc:	userSelectionContainerHTML,
	startPaletteDrag: startUserSelectionPaletteDrag,
	createNewItemAfterDropFunc: openNewUserSelectionDialog,
	resizeConstraints: elemResizeConstraints(200,400,75,75),
	resizeFunc: resizeUserSelection,
	initFunc: initDesignFormUserSelection,
	selectionFunc: selectFormUserSelection
}
