
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

function repositionCheckBox(checkBoxID,position) {
	
	var reposParams = {
		parentFormID: designFormContext.formID,
		checkBoxID: checkBoxID,
		position: position
	}
	
	jsonAPIRequest("frm/checkBox/reposition", reposParams, function(updatedObjRef) {
		setElemObjectRef(checkBoxID,updatedObjRef)
	})
	
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
	createNewItemAfterDropFunc: openNewCheckboxDialog,
	resizeConstraints: elemResizeConstraints(100,640,30,30),
	resizeFunc: resizeCheckBox,
	repositionFunc:repositionCheckBox,
	initFunc: initDesignFormCheckBox,
	selectionFunc: selectFormCheckbox
}
