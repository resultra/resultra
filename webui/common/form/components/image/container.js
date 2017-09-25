

function imageContainerHTML(elementID)
{
	
	var imageButton = '<span class="btn btn-default btn-sm clearButton pull-right fileinput-button imageEditLinkButton">' +
             	   '<span class="glyphicon glyphicon-picture"></span>' +
				   '<input class="uploadInput" type="file">' + 
            	'</span>'
	
	var containerHTML = '' +
		'<div class="layoutContainer imageComponent imageFormComponent">' +
			'<div class="row">' +
				'<div class="col-xs-7 componentHeaderLabelCol">' +
					'<label>New Image</label>' +
				'</div>' +
				'<div class="col-xs-5 componentHeaderButtonCol">' +
					smallClearComponentValHeaderButton("imageComponentClearValueButton") +
					imageButton +
					componentHelpPopupButtonHTML() +
				'</div>' +
			'</div>' +
			'<div class="imageDisplayContainer">' +
				'<a class="imageDisplay"></a>' +
			'</div>' +
		'</div>';
	return containerHTML
}

function imageTableViewContainerHTML() {
	
	var imageButton = '<span class="btn btn-default input-group-addon fileinput-button imageEditLinkButton">' +
             	   '<span class="glyphicon glyphicon-picture"></span>' +
				   '<input class="uploadInput" type="file">' + 
            	'</span>'
	
	var containerHTML = ''+
		'<div class="layoutContainer imageComponent imageTableCellComponent">' +
			'<div class="input-group">'+
				'<div class="form-control-static imageDisplayContainer">' +
					'<a class="imageDisplay"></a>' +
				'</div>' +
				imageButton + 
				clearValueButtonHTML("imageComponentClearValueButton") +
			'</div>' +
		'</div>'
		
	return containerHTML
}


function imageEditPopupViewContainerHTML() {
	var containerHTML = ''+
		'<div class="imagePopupContainer">' +
			'<div class="imageEditorHeader">' +
				'<button type="button" class="close closeImageEditorPopup" data-dismiss="modal" aria-hidden="true">x</button>' +
			'</div>' +
			'<div class="marginTop5">' +
				'<label>Email address</label>' + 
				'<input type="text" name="symbol" class="imageComponentInput form-control" placeholder="">'+
			'</div>' +
		'</div>'
	return containerHTML
	
}



function setImageComponentLabel($imageContainer, imageRef) {

	var $label = $imageContainer.find('label')
	
	setFormComponentLabel($label,imageRef.properties.fieldID,
			imageRef.properties.labelFormat)	
}

function initImageClearValueControl($imageContainer, imageRef) {
	initClearValueControl($imageContainer,imageRef,".imageComponentClearValueButton")	
}

function initImageEditAddrControl($imageContainer, imageRef) {
	
	var $editAddrButton = $imageContainer.find(".imageEditLinkButton")
	
	if(formComponentIsReadOnly(imageRef.properties.permissions)) {
		$editAddrButton.css("display","none")
	} else {
		$editAddrButton.css("display","")
	}
	
}

function initImageFormComponentContainer($container,imageRef) {
	setImageComponentLabel($container,imageRef)
	initImageClearValueControl($container, imageRef)
	initComponentHelpPopupButton($container, imageRef)
	initImageEditAddrControl($container,imageRef)
}