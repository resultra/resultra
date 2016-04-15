
// Definition of parameters and callbacks for a text box to be editable within the form editor.
// this javascript file needs to included after the other text box related files, so all the functions
// are already defined.

var textBoxEditConfig = {
	draggableHTMLFunc:	textBoxContainerHTML,
	createNewItemAfterDropFunc: newLayoutContainer,
	resizeConstraints: elemResizeConstraints(100,600,400,400),
	resizeFunc:  function (resizeParams) {
		console.log("resizing textBox: " + JSON.stringify(resizeParams))
	},
	reposFunc:  function (reposParams) {
		console.log("reposition textBox: " + JSON.stringify(reposParams))
	}	
}
