
function initUserSelectionDesignControlBehavior(userSelectionObjectRef) {
	var userSelectionControlSelector = '#' + userSelectionIDFromElemID(userSelectionObjectRef.userSelectionID)

//	$(userSelectionControlSelector).rating()
	
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

function startUserSelectionPaletteDrag(placeholderID) {
	var userSelectionControlSelector = '#'+userSelectionIDFromElemID(placeholderID)
//	$(userSelectionControlSelector).rating()
}


var userSelectionDesignFormConfig = {
	draggableHTMLFunc:	userSelectionContainerHTML,
	startPaletteDrag: startUserSelectionPaletteDrag,
	createNewItemAfterDropFunc: openNewUserSelectionDialog,
	resizeConstraints: elemResizeConstraints(100,640,30,30),
	resizeFunc: resizeUserSelection,
	initFunc: initDesignFormUserSelection,
	selectionFunc: selectFormUserSelection
}
