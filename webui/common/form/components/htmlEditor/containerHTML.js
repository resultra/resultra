

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
		'</div>' +
		'<div class="htmlEditorContent lightGreyBorder">' +
			'<div class="htmlEditorInput inlineContent htmlEditorDefaultBackground">'+
			'</div>' +
		'</div>'+
		'<div class="editorFooter pull-right componentHoverFooter initiallyHidden">' +
			smallClearDeleteButtonHTML("editorComponentClearValueButton") + 
			editButtonHTML() +
		'</div>' +
	
	'</div>';
	
		
	return containerHTML
}

function initEditorFormComponentViewModeGeometry($container,editorRef) {
	// In view mode, the height will be flexible, up the maximum set in the form designer.
	// This ensures there isn't any "dead space" when there aren't enough attachments to
	// fill up the attachment area below the header.
	setElemFixedWidthFlexibleHeight($container,editorRef.properties.geometry.sizeWidth)
	
	var $header = $container.find(".htmlEditorHeader")
	
	// Set the maximum height of the attachment area to be the remainder after the header
	// is accounted for.
	var headerBottom = $header.position().top + $header.outerHeight(true);
	
	var containerHeightPx = (editorRef.properties.geometry.sizeHeight - headerBottom) + "px"
	var editorHeightPx = (editorRef.properties.geometry.sizeHeight - headerBottom - 4) + "px"
	
	var $contentContainer = $container.find(".htmlEditorContent")
	var $editorContainer = $container.find(".htmlEditorInput")
	
	$contentContainer.css('height',containerHeightPx)
	$editorContainer.css('height',editorHeightPx)
	
}



function setEditorComponentLabel($editor,editorRef) {
	var $label = $editor.find('label')
	
	setFormComponentLabel($label,editorRef.properties.fieldID,
			editorRef.properties.labelFormat)	
	
}