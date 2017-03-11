function initDesignFormComment() {
	console.log("Init comment component design form behavior")
	initNewCommentComponentDialog()
}

function selectFormComment($container,commentObjRef) {
	console.log("Selected date picker: " + JSON.stringify(commentObjRef))
	loadCommentComponentProperties(commentObjRef)
}


function resizeCommentComponent($container,geometry) {
	
	var commentRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		commentID: commentRef.commentID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/comment/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.commentID)
	})	
}


var commentDesignFormConfig = {
	draggableHTMLFunc:	commentContainerHTML,
	startPaletteDrag: function(placeholderID,$paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewCommentComponentDialog,
	resizeConstraints: elemResizeConstraints(200,1280,200,1280),
	resizeHandles: 'e,s,se',
	resizeFunc:resizeCommentComponent,
	initFunc: initDesignFormComment,
	selectionFunc: selectFormComment
}
