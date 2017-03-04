
function initUserSelectionDesignControlBehavior(userSelectionObjectRef) {
// no-op	
}


function initDesignFormUserSelection() {
	initUserSelectionDialog()
}

function selectFormUserSelection($container,userSelectionObjectRef) {
	loadUserSelectionProperties(userSelectionObjectRef)
}

function resizeUserSelection($container,geometry) {
	
	var userSelRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		userSelectionID: userSelRef.userSelectionID,
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
