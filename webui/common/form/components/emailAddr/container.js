function emailAddrContainerInputControl() {
	return '<div class="input-group">'+
					'<input type="text" name="symbol" class="emailAddrComponentInput form-control" placeholder="">'+
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


function setEmailAddrComponentLabel($emailAddrContainer, emailAddrRef) {

	var $label = $emailAddrContainer.find('label')
	
	setFormComponentLabel($label,emailAddrRef.properties.fieldID,
			emailAddrRef.properties.labelFormat)	
}

function initEmailAddrClearValueControl($emailAddrContainer, emailAddrRef) {
	initClearValueControl($emailAddrContainer,emailAddrRef,".emailAddrComponentClearValueButton")	
}

function initEmailAddrFormComponentContainer($container,emailAddrRef) {
	setEmailAddrComponentLabel($container,emailAddrRef)
	initEmailAddrClearValueControl($container, emailAddrRef)
	initComponentHelpPopupButton($container, emailAddrRef)
}