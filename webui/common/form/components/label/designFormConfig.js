
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
		var labelWidth = updatedObjRef.properties.geometry.sizeWidth - 15
		initLabelSelectionControl($container, updatedObjRef,labelWidth)
	})	
}

function startLabelPaletteDrag(placeholderID,$paletteItemContainer) {
// no-op
}


var labelDesignFormConfig = {
	draggableHTMLFunc:	labelContainerHTML,
	initDummyDragAndDropComponentContainer: startLabelPaletteDrag,
	createNewItemAfterDropFunc: openNewLabelDialog,
	resizeConstraints: elemResizeConstraints(200,800,75,400),
	resizeHandles: 'e,s,se',
	resizeFunc: resizeLabel,
	initFunc: initDesignFormLabel,
	selectionFunc: selectFormLabel
}
