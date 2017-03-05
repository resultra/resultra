

function htmlInputFromHTMLEditorContainer($htmlEditor) {
	return 	$htmlEditor.find(".htmlEditorInput")
}


function htmlEditorContainerHTML(elementID)
{	
	var containerHTML = ''+
	'<div class=" layoutContainer htmlEditorContainer">' +
		'<div class="htmlEditorHeader">' +
			'<label>Editor Label</label>' +
		'</div>' +
		'<div class="htmlEditorContent">' +
			'<div class="htmlEditorInput">'+
			'</div>' +
		'</div>'+
	'</div>';
	
		
	return containerHTML
}
