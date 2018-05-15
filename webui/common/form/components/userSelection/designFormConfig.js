
function initUserSelectionDesignControlBehavior(userSelectionObjectRef) {
// no-op	
}


function initDesignFormUserSelection() {
	initUserSelectionDialog()
}

function selectFormUserSelection($container,userSelectionObjectRef) {
	loadUserSelectionProperties($container,userSelectionObjectRef)
}


function resizeUserSelection($container,geometry) {
	
	var userSelRef = getContainerObjectRef($container)
	
	initDummiedUpUserSelectionControl($container,geometry.sizeWidth)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		userSelectionID: userSelRef.userSelectionID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/userSelection/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.userSelectionID)
	})	
}

function startUserSelectionPaletteDrag($paletteItemContainer) {
	initDummiedUpUserSelectionControl($paletteItemContainer,250)
}


var userSelectionDesignFormConfig = {
	draggableHTMLFunc:	userSelectionContainerHTML,
	initDummyDragAndDropComponentContainer: startUserSelectionPaletteDrag,
	createNewItemAfterDropFunc: openNewUserSelectionDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(200,1200),
	resizeFunc: resizeUserSelection,
	initFunc: initDesignFormUserSelection,
	selectionFunc: selectFormUserSelection
}
