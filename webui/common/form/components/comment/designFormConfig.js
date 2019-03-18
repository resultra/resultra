// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function initDesignFormComment() {
	console.log("Init comment component design form behavior")
	initNewCommentComponentDialog()
}

function selectFormComment($container,commentObjRef) {
	console.log("Selected date picker: " + JSON.stringify(commentObjRef))
	loadCommentComponentProperties($container,commentObjRef)
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
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewCommentComponentDialog,
	resizeConstraints: elemResizeConstraints(200,1280,200,1280),
	resizeHandles: 'e,s,se',
	resizeFunc:resizeCommentComponent,
	initFunc: initDesignFormComment,
	selectionFunc: selectFormComment
}
