
// Definition of parameters and callbacks for a date picker to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormHtmlEditor() {
	console.log("Init html editor design form behavior")
	initNewHtmlEditorDialog()
}

function selectFormHtmlEditor(htmlEditorObjRef) {
	console.log("Selected html editor: " + JSON.stringify(htmlEditorObjRef))
	loadHtmlEditorProperties(htmlEditorObjRef)
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
	})	
}


var htmlEditorDesignFormConfig = {
	draggableHTMLFunc:	htmlEditorContainerHTML,
	startPaletteDrag: function(placeholderID,$paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewHtmlEditorDialog,
	resizeConstraints: elemResizeConstraints(100,640,100,500),
	resizeFunc: resizeHtmlEditor,
	initFunc: initDesignFormHtmlEditor,
	selectionFunc: selectFormHtmlEditor
}
