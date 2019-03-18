// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Definition of parameters and callbacks for a date picker to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormHtmlEditor() {
	console.log("Init html editor design form behavior")
	initNewHtmlEditorDialog()
}

function selectFormHtmlEditor($container,htmlEditorObjRef) {
	console.log("Selected html editor: " + JSON.stringify(htmlEditorObjRef))
	loadHtmlEditorProperties($container,htmlEditorObjRef)
}


function resizeHtmlEditor($container,geometry) {
	
	var editorRef = getContainerObjectRef($container)
	
	var resizeParams = {
		parentFormID: designFormContext.formID,
		htmlEditorID: editorRef.htmlEditorID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/htmlEditor/resize", resizeParams, function(updatedObjRef) {
		setContainerComponentInfo($container,updatedObjRef,updatedObjRef.htmlEditorID)
		initEditorFormComponentViewModeGeometry($container,updatedObjRef)
	})	
}

function resizeInProgressHTMLContainer($container,geometry) {
	initHTMLEditorComponentViewModeGeometry($container,geometry.sizeWidth,geometry.sizeHeight)
}


var htmlEditorDesignFormConfig = {
	draggableHTMLFunc:	htmlEditorContainerHTML,
	initDummyDragAndDropComponentContainer: function($paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewHtmlEditorDialog,
	resizeConstraints: elemResizeConstraints(125,1280,125,1280),
	resizeHandles: 'e,s,se',
	resizeFunc: resizeHtmlEditor,
	resizeInProgressFunc: resizeInProgressHTMLContainer,
	initFunc: initDesignFormHtmlEditor,
	selectionFunc: selectFormHtmlEditor
}
