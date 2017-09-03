
function initLabelDesignControlBehavior(labelObjectRef) {
// no-op	
}


function initDesignFormLabel() {
	initLabelDialog()
}

function selectFormLabel($container,labelObjectRef) {
	loadLabelProperties($container,labelObjectRef)
}

function resizeLabel($container,geometry) {
	
	var userSelRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		labelID: userSelRef.labelID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/label/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.labelID)
	})	
}

function startLabelPaletteDrag(placeholderID,$paletteItemContainer) {
// no-op
}


var labelDesignFormConfig = {
	draggableHTMLFunc:	labelContainerHTML,
	initDummyDragAndDropComponentContainer: startLabelPaletteDrag,
	createNewItemAfterDropFunc: openNewLabelDialog,
	resizeConstraints: elemResizeConstraints(200,400,75,75),
	resizeFunc: resizeLabel,
	initFunc: initDesignFormLabel,
	selectionFunc: selectFormLabel
}
