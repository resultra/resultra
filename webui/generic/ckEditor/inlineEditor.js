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
	$editorContainer.attr("contenteditable","false")
}


function enableInlineCKEditor($editorContainer) {
	
	// Must be set for editing to occur
	$editorContainer.attr("contenteditable","true")
	
	// Styles which will appear in the "Styles" menu
	  var bootstrapStyles = [{
	  		name: 'Marker',
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
	  	's u p strong em mark big small blockquote del ins;' +
	  	'a[!href];'

	  var toolbarConfig = [{
	  	name: 'links',
	  	items: ['Link', 'Unlink']
	  }, {
	  	name: 'paragraph',
	  	items: ['Bold', 'Italic', 'Strike', 'Underline', '-', 'RemoveFormat']
	  }, {
	  	name: 'styles',
	  	items: ['Styles', 'Format']
	  }, {
	  	name: 'paragraph',
	  	items: ['NumberedList', 'BulletedList']
	  }]

	  // Styles which will appear in the 'Format" menu
	  var formats = 'p;h1;h2;h3;h4;h5;h6;pre'

	  var editorInputDOMElem = $editorContainer.get(0)
	  
	  var editor = CKEDITOR.inline(editorInputDOMElem, {
	  	toolbar: toolbarConfig,
	  	stylesSet: bootstrapStyles,
	  	allowedContent: allowedContent,
	  	format_tags: formats
	  })
	  
	  return editor
	
}