
function initDesignFormHeader() {
	console.log("Init header design form behavior")
//	initNewCheckBoxDialog()
}

function selectFormHeader($container,headerObjRef) {
	console.log("Selected header: " + JSON.stringify(headerObjRef))
	loadFormHeaderProperties(headerObjRef)
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
	startPaletteDrag: function(placeholderID,$paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewFormHeaderDialog,
	resizeConstraints: elemResizeConstraints(320,640,50,50),
	resizeFunc: resizeFormHeader,
	initFunc: initDesignFormHeader,
	selectionFunc: selectFormHeader
}
