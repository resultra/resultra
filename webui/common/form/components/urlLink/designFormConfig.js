
// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript file needs to included after the other text box related files, so all the functions
// are already defined.


function initDesignFormUrlLink() {
	initNewUrlLinkDialog()
}

function selectFormUrlLink ($container,urlLinkRef) {
	loadUrlLinkProperties($container,urlLinkRef)
}

function resizeUrlLink($container,geometry) {
	
	var urlLinkRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		urlLinkID: urlLinkRef.urlLinkID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/urlLink/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.urlLinkID)
	})	
}


var urlLinkDesignFormConfig = {
	draggableHTMLFunc:	urlLinkContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewUrlLinkDialog,
	resizeConstraints: elemResizeConstraintsWidthOnly(75,600),
	resizeFunc: resizeUrlLink,
	initFunc: initDesignFormUrlLink,
	selectionFunc: selectFormUrlLink
}
