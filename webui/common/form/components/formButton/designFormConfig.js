
function initDesignFormButton() {
	console.log("Init button design form behavior")
}

function selectFormButton(buttonObjRef) {
	console.log("Selected button: " + JSON.stringify(buttonObjRef))
	loadFormButtonProperties(buttonObjRef)
}

function resizeFormButton(buttonID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		buttonID: buttonID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/formButton/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(buttonID,updatedObjRef)
	})	
}

var formButtonDesignFormConfig = {
	draggableHTMLFunc:	formButtonContainerHTML,
	startPaletteDrag: function(placeholderID) {},
	createNewItemAfterDropFunc: openNewFormButtonDialog,
	resizeConstraints: elemResizeConstraints(320,640,50,50),
	resizeFunc: resizeFormButton,
	initFunc: initDesignFormButton,
	selectionFunc: selectFormButton
}
