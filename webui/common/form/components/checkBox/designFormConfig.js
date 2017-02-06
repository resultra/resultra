
// Definition of parameters and callbacks for a check box to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormCheckBox() {
	console.log("Init checkbox design form behavior")
	initNewCheckBoxDialog()
}

function selectFormCheckbox(checkboxObjRef) {
	console.log("Selected checkbox: " + JSON.stringify(checkboxObjRef))
	loadCheckboxProperties(checkboxObjRef)
}

function resizeCheckBox(checkBoxID,geometry) {
	var resizeParams = {
		parentFormID: designFormContext.formID,
		checkBoxID: checkBoxID,
		geometry: geometry
	}
	
	jsonAPIRequest("frm/checkBox/resize", resizeParams, function(updatedObjRef) {
		setElemObjectRef(checkBoxID,updatedObjRef)
	})	
}


var checkBoxDesignFormConfig = {
	draggableHTMLFunc:	checkBoxContainerHTML,
	startPaletteDrag: function(placeholderID,$paletteItemContainer) {},
	createNewItemAfterDropFunc: openNewCheckboxDialog,
	resizeConstraints: elemResizeConstraints(100,640,30,30),
	resizeFunc: resizeCheckBox,
	initFunc: initDesignFormCheckBox,
	selectionFunc: selectFormCheckbox
}
