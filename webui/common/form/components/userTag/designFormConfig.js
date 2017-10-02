
function initUserTagDesignControlBehavior(userTagObjectRef) {
// no-op	
}


function initDesignFormUserTag() {
	initUserTagDialog()
}

function selectFormUserTag($container,userTagObjectRef) {
	loadUserTagProperties($container,userTagObjectRef)
}


function resizeUserTag($container,geometry) {
	
	var userSelRef = getContainerObjectRef($container)
	
	initDummiedUpUserTagControl($container,geometry.sizeWidth)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		userTagID: userSelRef.userTagID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/userTag/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.userTagID)
	})	
}

function startUserTagPaletteDrag($paletteItemContainer) {
	initDummiedUpUserTagControl($paletteItemContainer,250)
}


var userTagDesignFormConfig = {
	draggableHTMLFunc:	userTagContainerHTML,
	initDummyDragAndDropComponentContainer: startUserTagPaletteDrag,
	createNewItemAfterDropFunc: openNewUserTagDialog,
	resizeConstraints: elemResizeConstraints(200,1200,75,75),
	resizeFunc: resizeUserTag,
	initFunc: initDesignFormUserTag,
	selectionFunc: selectFormUserTag
}
