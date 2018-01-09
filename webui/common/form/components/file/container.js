function fileContainerInputControl() {

	return '<div class="input-group">'+
				'<div class="formInputStaticInputContainer">' +
					'<a class="fileDisplay"></a>' +
				'</div>' + 
				'<span class="btn btn-default input-group-addon fileinput-button fileEditLinkButton">' +
             	   '<span class="glyphicon glyphicon-file"></span>' +
				   '<input class="uploadInput" type="file">' + 
            	'</span>' +
				clearValueButtonHTML("fileComponentClearValueButton") +
			'</div>'


}

function fileContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer fileComponent fileFormComponent">' +
			'<div class="container-fluid componentHeader">' + 
				'<div class="row">' +
					'<div class="col-xs-9 componentHeaderLabelCol">' +
						'<label>New File</label>' +
					'</div>' +
					'<div class="col-xs-3 componentHeaderButtonCol">' +
						componentHelpPopupButtonHTML() +
					'</div>' +
				'</div>' +
			'</div>' +
			fileContainerInputControl() +
		'</div>';
	return containerHTML
}

function fileTableViewContainerHTML() {
	var containerHTML = ''+
		'<div class="layoutContainer fileComponent fileTableCellComponent">' +
			fileContainerInputControl() +
		'</div>';
	return containerHTML
}


function fileEditPopupViewContainerHTML() {
	var containerHTML = ''+
		'<div class="filePopupContainer">' +
			'<div class="fileEditorHeader">' +
				'<button type="button" class="close closeFileEditorPopup" data-dismiss="modal" aria-hidden="true">x</button>' +
			'</div>' +
			'<div class="marginTop5">' +
				'<label>Email address</label>' + 
				'<input type="text" name="symbol" class="fileComponentInput form-control" placeholder="">'+
			'</div>' +
		'</div>'
	return containerHTML
	
}



function setFileComponentLabel($fileContainer, fileRef) {

	var $label = $fileContainer.find('label')
	
	setFormComponentLabel($label,fileRef.properties.fieldID,
			fileRef.properties.labelFormat)	
}

function initFileClearValueControl($fileContainer, fileRef) {
	initClearValueControl($fileContainer,fileRef,".fileComponentClearValueButton")	
}

function initFileEditAddrControl($fileContainer, fileRef) {
	
	var $editAddrButton = $fileContainer.find(".fileEditLinkButton")
	
	if(formComponentIsReadOnly(fileRef.properties.permissions)) {
		$editAddrButton.css("display","none")
	} else {
		$editAddrButton.css("display","")
	}
	
}

function initFileFormComponentContainer($container,fileRef) {
	setFileComponentLabel($container,fileRef)
	initFileClearValueControl($container, fileRef)
	initComponentHelpPopupButton($container, fileRef)
	initFileEditAddrControl($container,fileRef)
}