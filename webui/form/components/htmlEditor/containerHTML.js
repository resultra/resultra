

function htmlInputIDFromContainerElemID(htmlEditorElemID) {
	return 	htmlEditorElemID + '_htmlInput'
}


function htmlEditorContainerHTML(elementID)
{	
	var htmlInputID = htmlInputIDFromContainerElemID(elementID)
	
	var containerHTML = ''+
	'<div class="ui-widget-content layoutContainer htmlEditorContainer  draggable resizable" id="'+elementID+'">' +
		'<label>Editor Label</label>' +
		'<div id="' + htmlInputID + '" contenteditable="true" class="htmlEditorInput">'+
		'</div>'+
	'</div>';
	
		
	return containerHTML
}
