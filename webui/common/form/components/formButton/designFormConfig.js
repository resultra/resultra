
function initDesignFormButton() {
	console.log("Init button design form behavior")
}

function selectFormButton(buttonObjRef) {
	console.log("Selected button: " + JSON.stringify(buttonObjRef))
	loadFormButtonProperties(buttonObjRef)
}

function resizeFormButton($container,geometry) {
	
	var buttonRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		buttonID: buttonRef.buttonID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/formButton/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.buttonID)
	})	
}

var formButtonDesignFormConfig = {
	draggableHTMLFunc:	formButtonContainerHTML,
	startPaletteDrag: function(placeholderID,$paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewFormButtonDialog,
	resizeConstraints: elemResizeConstraints(320,640,50,50),
	resizeFunc: resizeFormButton,
	initFunc: initDesignFormButton,
	selectionFunc: selectFormButton
}
