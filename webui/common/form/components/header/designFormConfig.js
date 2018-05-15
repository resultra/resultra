
function initDesignFormHeader() {
	console.log("Init header design form behavior")
//	initNewCheckBoxDialog()
}

function selectFormHeader($container,headerObjRef) {
	console.log("Selected header: " + JSON.stringify(headerObjRef))
	loadFormHeaderProperties($container,headerObjRef)
}

function resizeFormHeader($container,geometry) {
	
	var headerRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		headerID: headerRef.headerID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/header/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.headerID)
	})	
}

var formHeaderDesignFormConfig = {
	draggableHTMLFunc:	formHeaderContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewFormHeaderDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(80,720),
	resizeFunc: resizeFormHeader,
	initFunc: initDesignFormHeader,
	selectionFunc: selectFormHeader
}
