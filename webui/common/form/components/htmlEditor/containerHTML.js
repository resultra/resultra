

function htmlInputIDFromContainerElemID(htmlEditorElemID) {
	return 	htmlEditorElemID + '_htmlInput'
}


function htmlEditorContainerHTML(elementID)
{	
	var htmlInputID = htmlInputIDFromContainerElemID(elementID)
	
	var containerHTML = ''+
	'<div class="ui-widget-content layoutContainer htmlEditorContainer  draggable resizable" id="'+elementID+'">' +
		'<div class="htmlEditorHeader">' +
			'<label>Editor Label</label>' +
		'</div>' +
		'<div class="htmlEditorContent">' +
			'<div id="' + htmlInputID + '" contenteditable="true" class="htmlEditorInput">'+
			'</div>' +
		'</div>'+
	'</div>';
	
		
	return containerHTML
}
