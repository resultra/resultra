
// Definition of parameters and callbacks for a check box to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

function initDesignFormCheckBox() {
	console.log("Init checkbox design form behavior")
	initNewCheckBoxDialog()
}

function selectFormCheckbox(checkboxObjRef) {
	console.log("Selected checkbox: " + JSON.stringify(checkboxObjRef))
}


var checkBoxDesignFormConfig = {
	draggableHTMLFunc:	checkBoxContainerHTML,
	createNewItemAfterDropFunc: openNewCheckboxDialog,
	resizeConstraints: elemResizeConstraints(100,640,30,30),
	resizeAPIName: "frm/checkBox/resize",
	reposAPIName: "frm/checkBox/reposition",
	initFunc: initDesignFormCheckBox,
	selectionFunc: selectFormCheckbox
}
