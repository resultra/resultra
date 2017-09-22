function imageContainerInputControl() {

	return '<div class="input-group">'+
				'<div class="form-control-static imageDisplayContainer">' +
					'<a class="imageDisplay"></a>' +
				'</div>' + 
				'<span class="btn btn-default input-group-addon fileinput-button imageEditLinkButton">' +
             	   '<span class="glyphicon glyphicon-picture"></span>' +
				   '<input class="uploadInput" type="file">' + 
            	'</span>' +
				clearValueButtonHTML("imageComponentClearValueButton") +
			'</div>'


}

function imageContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer imageComponent imageFormComponent">' +
				'<label>New Image</label>'+ componentHelpPopupButtonHTML() +
				imageContainerInputControl() +
		'</div>';
	return containerHTML
}

function imageTableViewContainerHTML() {
	var containerHTML = ''+
		'<div class="layoutContainer imageComponent imageTableCellComponent">' +
			imageContainerInputControl() +
		'</div>';
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