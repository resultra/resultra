
function initDesignFormButton() {
	console.log("Init button design form behavior")
}

function selectFormButton($container,buttonObjRef) {
	console.log("Selected button: " + JSON.stringify(buttonObjRef))
	loadFormButtonProperties($container,buttonObjRef)
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
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewFormButtonDialog,
	resizeConstraints: elemResizeConstraints(50,640,50,50),
	resizeFunc: resizeFormButton,
	initFunc: initDesignFormButton,
	selectionFunc: selectFormButton
}
