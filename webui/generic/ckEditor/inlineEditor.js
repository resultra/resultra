// This file is part of the Resultra project.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
function inlineCKEditorEnabled($editorContainer) {
	
	var editableAttr = $editorContainer.attr("contenteditable")
	if (editableAttr === "true") {
		return true
	} else {
		return false
	}
}

function disableInlineCKEditor($editorContainer,editor) {
	editor.destroy()
	var rawHTML = $editorContainer.html()
	var formattedHTML = formatInlineContentHTMLDisplay(rawHTML)
	$editorContainer.html(formattedHTML)
	$editorContainer.attr("contenteditable","false")
}


function enableInlineCKEditor($editorContainer) {
	
	// Must be set for editing to occur
	$editorContainer.attr("contenteditable","true")
	
	// Styles which will appear in the "Styles" menu
	  var bootstrapStyles = [{
	  		name: 'Highlight',
	  		element: 'mark'
	  	},
	  	{
	  		name: 'Big',
	  		element: 'big'
	  	}, {
	  		name: 'Small',
	  		element: 'small'
	  	}
	  ]

	  var allowedContent = 'h1 h2 h3 h4 h5 h6;' +
	  	'ol ul li;' +
	  	's u p strong em mark big small blockquote del ins hr;' +
	  	'a[!href];'

	  var toolbarConfig = [
		  { name: 'paragraph', items: ['NumberedList', 'BulletedList','-','Outdent','Indent','-','Blockquote']}, 
		  { name: 'styles', items: ['Format','Styles'] },
		  { name: 'insert', items: ['HorizontalRule']},
	  	   '/', 
		  { name: 'paragraph', items: ['Bold', 'Italic', 'Strike', 'Underline', '-', 'RemoveFormat'] }, 
		  { name: 'links', items: ['Link', 'Unlink'] },
		  { name: 'tools', items: ['Maximize']}
	  ]


	  // Styles which will appear in the 'Format" menu
	  var formats = 'p;h1;h2;h3;h4;h5;h6;pre'

	  var editorInputDOMElem = $editorContainer.get(0)
	  
	  var editor = CKEDITOR.inline(editorInputDOMElem, {
	  	toolbar: toolbarConfig,
	  	stylesSet: bootstrapStyles,
	  	allowedContent: allowedContent,
	  	format_tags: formats,
		   // Override the default for removed buttons. Notably, this allows underline to appear.
		  removeButtons: 'Subscript,Superscript',
		  removePlugins: 'magicline', // remove the red 'new page' control which by default appears at the bottom of the editing area.
		  title:false // Disable the "Rich Text Editor" popup which appears by default when hovering over the editing area
	  })
	  
	  return editor
	
}