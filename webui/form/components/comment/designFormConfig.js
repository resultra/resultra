function initDesignFormComment() {
	console.log("Init comment component design form behavior")
	initNewCommentComponentDialog()
}

function selectFormComment(commentObjRef) {
	console.log("Selected date picker: " + JSON.stringify(commentObjRef))
	loadCommentComponentProperties(commentObjRef)
}


function resizeCommentComponent(commentID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		commentID: commentID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/comment/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(commentID,updatedObjRef)
	})	
}


var commentDesignFormConfig = {
	draggableHTMLFunc:	commentContainerHTML,
	startPaletteDrag: function(placeholderID) {},
	createNewItemAfterDropFunc: openNewCommentComponentDialog,
	resizeConstraints: elemResizeConstraints(100,640,30,30),
	resizeFunc: resizeCommentComponent,
	initFunc: initDesignFormComment,
	selectionFunc: selectFormComment
}
