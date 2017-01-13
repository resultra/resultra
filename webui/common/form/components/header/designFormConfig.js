
function initDesignFormHeader() {
	console.log("Init header design form behavior")
//	initNewCheckBoxDialog()
}

function selectFormHeader(headerObjRef) {
	console.log("Selected header: " + JSON.stringify(headerObjRef))
	loadFormHeaderProperties(headerObjRef)
}

function resizeFormHeader(headerID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		headerID: headerID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/header/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(headerID,updatedObjRef)
	})	
}

var formHeaderDesignFormConfig = {
	draggableHTMLFunc:	formHeaderContainerHTML,
	startPaletteDrag: function(placeholderID) {},
	createNewItemAfterDropFunc: openNewFormHeaderDialog,
	resizeConstraints: elemResizeConstraints(320,640,50,50),
	resizeFunc: resizeFormHeader,
	initFunc: initDesignFormHeader,
	selectionFunc: selectFormHeader
}
