function emailAddrContainerInputControl() {

	return '<div class="input-group">'+
				'<div class="formInputStaticInputContainer">' +
					'<a class="emailAddrDisplay"></a>' +
				'</div>' + 
				'<span class="input-group-addon emailAddrEditLinkButton">' +
             	   '<span class="glyphicon glyphicon-envelope"></span>' +
            	'</span>' +
				clearValueButtonHTML("emailAddrComponentClearValueButton") +
			'</div>'


}

function emailAddrContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer emailAddrComponent emailAddrFormComponent">' +
				'<label>New Email Address</label>'+ componentHelpPopupButtonHTML() +
				emailAddrContainerInputControl() +
		'</div>';
	return containerHTML
}

function emailAddrTableViewContainerHTML() {
	var containerHTML = ''+
		'<div class="layoutContainer emailAddrComponent emailAddrTableCellComponent">' +
			emailAddrContainerInputControl() +
		'</div>';
	return containerHTML
}


function emailAddrEditPopupViewContainerHTML() {
	var containerHTML = ''+
		'<div class="emailAddrPopupContainer">' +
			'<div class="emailAddrEditorHeader">' +
				'<button type="button" class="close closeEmailAddrEditorPopup" data-dismiss="modal" aria-hidden="true">x</button>' +
			'</div>' +
			'<div class="marginTop5">' +
				'<label>Email address</label>' + 
				'<input type="text" name="symbol" class="emailAddrComponentInput form-control" placeholder="">'+
			'</div>' +
		'</div>'
	return containerHTML
	
}



function setEmailAddrComponentLabel($emailAddrContainer, emailAddrRef) {

	var $label = $emailAddrContainer.find('label')
	
	setFormComponentLabel($label,emailAddrRef.properties.fieldID,
			emailAddrRef.properties.labelFormat)	
}

function initEmailAddrClearValueControl($emailAddrContainer, emailAddrRef) {
	initClearValueControl($emailAddrContainer,emailAddrRef,".emailAddrComponentClearValueButton")	
}

function initEmailAddrEditAddrControl($emailAddrContainer, emailAddrRef) {
	
	var $editAddrButton = $emailAddrContainer.find(".emailAddrEditLinkButton")
	
	if(formComponentIsReadOnly(emailAddrRef.properties.permissions)) {
		$editAddrButton.css("display","none")
	} else {
		$editAddrButton.css("display","")
	}
	
}

function calcEmailAddrMinTableCellColWidth(emailAddrRef,emailAddrText) {
	
	var addrWidth = calcTextWidth(emailAddrText)
	var paddingWidth = 10
	
	var addonWidth = 26
	if(formComponentIsReadOnly(emailAddrRef.properties.permissions)) { addonWidth +=26  }
	
	if (clearValueControlIsEnabled(emailAddrRef)) {
		addonWidth += 17
	}
	
	var unconstrainedWidth = addrWidth + paddingWidth + addonWidth
	
	return calcContrainedPxVal(unconstrainedWidth,250,400)
}

function initEmailAddrFormComponentContainer($container,emailAddrRef) {
	setEmailAddrComponentLabel($container,emailAddrRef)
	initEmailAddrClearValueControl($container, emailAddrRef)
	initComponentHelpPopupButton($container, emailAddrRef)
	initEmailAddrEditAddrControl($container,emailAddrRef)
}