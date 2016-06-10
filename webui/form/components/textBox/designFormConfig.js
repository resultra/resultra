
// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript file needs to included after the other text box related files, so all the functions
// are already defined.


function initDesignFormTextBox() {
	console.log("Init text box design form behavior")
	initNewTextBoxDialog()
}

function selectFormTextBox (textBoxRef) {
	console.log("Select text box: " + JSON.stringify(textBoxRef))
	loadTextBoxProperties(textBoxRef)
}

function repositionTextBox(textBoxID,position) {
	
	var reposParams = {
		parentFormID: designFormContext.formID,
		textBoxID: textBoxID,
		position: position
	}
	
	jsonAPIRequest("frm/textBox/reposition", reposParams, function(updatedObjRef) {
		setElemObjectRef(textBoxID,updatedObjRef)
	})
	
}

function resizeTextBox(textBoxID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		textBoxID: textBoxID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/textBox/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(textBoxID,updatedObjRef)
	})	
}


var textBoxDesignFormConfig = {
	draggableHTMLFunc:	textBoxContainerHTML,
	createNewItemAfterDropFunc: openNewTextBoxDialog,
	resizeConstraints: elemResizeConstraints(100,600,400,400),
	resizeFunc: resizeTextBox,
	repositionFunc: repositionTextBox,
	initFunc: initDesignFormTextBox,
	selectionFunc: selectFormTextBox
}
