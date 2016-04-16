
// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript file needs to included after the other text box related files, so all the functions
// are already defined.

var textBoxDesignFormConfig = {
	draggableHTMLFunc:	textBoxContainerHTML,
	createNewItemAfterDropFunc: newLayoutContainer,
	resizeConstraints: elemResizeConstraints(100,600,400,400),
	resizeAPIName: "frm/textBox/resize",
	reposAPIName: "frm/textBox/reposition"
}
