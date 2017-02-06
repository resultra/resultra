
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

function startUserSelectionPaletteDrag(placeholderID,$paletteItemContainer) {
	var userSelectionControlSelector = '#'+userSelectionIDFromElemID(placeholderID)
//	$(userSelectionControlSelector).rating()
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
