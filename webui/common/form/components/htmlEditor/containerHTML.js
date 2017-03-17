

function htmlInputFromHTMLEditorContainer($htmlEditor) {
	return 	$htmlEditor.find(".htmlEditorInput")
}


function htmlEditorContainerHTML(elementID)
{
	
	function editButtonHTML() {
	
		// className is to uniquely identify the button with other HTML elements,
		// such that it can be found with jQuery's find() function.
	
		var buttonHTML = '<button class="btn btn-default btn-sm clearButton ' + 
				'startEditButton' + 
				'"><span class="glyphicon glyphicon-pencil"></span></button>'
	
		return buttonHTML
	}
		
	var containerHTML = ''+
	'<div class=" layoutContainer htmlEditorContainer">' +
		'<div class="htmlEditorHeader">' +
			'<label>Editor Label</label>' +
			editButtonHTML() +
		'</div>' +
		'<div class="htmlEditorContent lightGreyBorder">' +
			'<div class="htmlEditorInput inlineContent htmlEditorDefaultBackground">'+
			'</div>' +
		'</div>'+
	'</div>';
	
		
	return containerHTML
}


function setEditorComponentLabel($editor,editorRef) {
	var $label = $editor.find('label')
	
	setFormComponentLabel($label,editorRef.properties.fieldID,
			editorRef.properties.labelFormat)	
	
}