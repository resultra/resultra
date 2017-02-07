

function htmlInputFromHTMLEditorContainer($htmlEditor) {
	return 	$htmlEditor.find(".htmlEditorInput")
}


function htmlEditorContainerHTML(elementID)
{	
	var containerHTML = ''+
	'<div class="ui-widget-content layoutContainer htmlEditorContainer  draggable resizable">' +
		'<div class="htmlEditorHeader">' +
			'<label>Editor Label</label>' +
		'</div>' +
		'<div class="htmlEditorContent">' +
			'<div contenteditable="true" class="htmlEditorInput">'+
			'</div>' +
		'</div>'+
	'</div>';
	
		
	return containerHTML
}
