
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
	startPaletteDrag: function(placeholderID) {},
	createNewItemAfterDropFunc: openNewTextBoxDialog,
	resizeConstraints: elemResizeConstraints(100,600,400,400),
	resizeFunc: resizeTextBox,
	initFunc: initDesignFormTextBox,
	selectionFunc: selectFormTextBox
}
