

function htmlInputFromHTMLEditorContainer($htmlEditor) {
	return 	$htmlEditor.find(".htmlEditorInput")
}

function noteEditButtonHTML() {

	// className is to uniquely identify the button with other HTML elements,
	// such that it can be found with jQuery's find() function.

	var buttonHTML = '<button tabindex="-1" class="btn btn-default btn-sm clearButton ' + 
			'startEditButton' + 
			'"><span class="glyphicon glyphicon-pencil"></span></button>'

	return buttonHTML
}



function htmlEditorContainerHTML(elementID)
{	
	var containerHTML = ''+
	'<div class=" layoutContainer htmlEditorContainer">' +
		'<div class="htmlEditorHeader">' +
			'<label>Editor Label</label>' +
		'</div>' +
		'<div class="htmlEditorContent lightGreyBorder">' +
			'<div class="htmlEditorInput inlineContent htmlEditorDefaultBackground">'+
			'</div>' +
		'</div>'+
		'<div class="editorFooter componentHoverFooter">' +
			smallClearDeleteButtonHTML("editorComponentClearValueButton") + 
			noteEditButtonHTML() +
		'</div>' +	
	'</div>';
			
	return containerHTML
}

function noteEditorTableViewContainerHTML() {
	var containerHTML = ''+
	'<div class="noteEditorPopupContainer">' +
		'<div class="htmlEditorHeader">' +
			'<button type="button" class="close closeEditorPopup" data-dismiss="modal" aria-hidden="true">x</button>' +
		'</div>' +
		'<div class="htmlEditorContent lightGreyBorder marginTop5">' +
			'<div class="htmlEditorInput inlineContent htmlEditorDefaultBackground">'+
			'</div>' +
		'</div>'+
		'<div class="editorFooter">' +
			smallClearDeleteButtonHTML("editorComponentClearValueButton") + 
			noteEditButtonHTML() +
		'</div>' +	
	'</div>';
			
	return containerHTML
	
}

function noteEditorTableViewCellContainerHTML() {
	return '<div class="layoutContainer noteEditTableCell">' +
			'<div>' +
				'<a class="btn noteEditPopop">Show note</a>'+
			'</div>' +
		'</div>'
}


function initHTMLEditorTextCellComponentViewModeGeometry($container) {
	
	var width = 250
	var height = 230
	
	setElemFixedWidthFlexibleHeight($container,width)
	
	var $header = $container.find(".htmlEditorHeader")
	var headerBottom = $header.position().top + $header.outerHeight(true);
	
	var containerHeightPx = (height - headerBottom) + "px"
	var editorHeightPx = (height - headerBottom - 4) + "px"
	
	var $contentContainer = $container.find(".htmlEditorContent")
	var $editorContainer = $container.find(".htmlEditorInput")
	
	$contentContainer.css('height',containerHeightPx)
	$editorContainer.css('height',editorHeightPx)
	
}

function initHTMLEditorComponentViewModeGeometry($container,width,height) {
	// In view mode, the height will be flexible, up the maximum set in the form designer.
	// This ensures there isn't any "dead space" when there aren't enough attachments to
	// fill up the attachment area below the header.
	setElemFixedWidthFlexibleHeight($container,width)
	
	var $header = $container.find(".htmlEditorHeader")
	
	// Set the maximum height of the attachment area to be the remainder after the header
	// is accounted for.
	var headerBottom = $header.position().top + $header.outerHeight(true);
	
	var containerHeightPx = (height - headerBottom) + "px"
	var editorHeightPx = (height - headerBottom - 4) + "px"
	
	var $contentContainer = $container.find(".htmlEditorContent")
	var $editorContainer = $container.find(".htmlEditorInput")
	
	$contentContainer.css('height',containerHeightPx)
	$editorContainer.css('height',editorHeightPx)
	
}


function initEditorFormComponentViewModeGeometry($container,editorRef) {
	
	var width = editorRef.properties.geometry.sizeWidth
	var height = editorRef.properties.geometry.sizeHeight
	
	initHTMLEditorComponentViewModeGeometry($container,width,height)
		
}



function setEditorComponentLabel($editor,editorRef) {
	var $label = $editor.find('label')
	
	setFormComponentLabel($label,editorRef.properties.fieldID,
			editorRef.properties.labelFormat)	
	
}