
// Definition of parameters and callbacks for a check box to be editable within the form editor.
// this javascript file needs to included after the other check box related files, so all the functions
// are already defined.

var checkBoxEditConfig = {
	draggableHTMLFunc:	checkBoxContainerHTML,
	createNewItemAfterDropFunc: openNewCheckboxDialog,
	resizeConstraints: elemResizeConstraints(100,600,400,400),
	resizeAPIName: "frm/checkBox/resize",
	reposAPIName: "frm/checkBox/reposition"
}
