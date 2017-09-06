function urlLinkContainerInputControl() {
	return '<div class="input-group">'+
					'<input type="text" name="symbol" class="urlLinkComponentInput form-control" placeholder="">'+
					clearValueButtonHTML("urlLinkComponentClearValueButton") +
				'</div>'
}

function urlLinkContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer urlLinkComponent urlLinkFormComponent">' +
				'<label>New Email Address</label>'+ componentHelpPopupButtonHTML() +
				urlLinkContainerInputControl() +
		'</div>';
	return containerHTML
}

function urlLinkTableViewContainerHTML() {
	var containerHTML = ''+
		'<div class="layoutContainer urlLinkComponent urlLinkTableCellComponent">' +
			urlLinkContainerInputControl() +
		'</div>';
	return containerHTML
}


function setUrlLinkComponentLabel($urlLinkContainer, urlLinkRef) {

	var $label = $urlLinkContainer.find('label')
	
	setFormComponentLabel($label,urlLinkRef.properties.fieldID,
			urlLinkRef.properties.labelFormat)	
}

function initUrlLinkClearValueControl($urlLinkContainer, urlLinkRef) {
	initClearValueControl($urlLinkContainer,urlLinkRef,".urlLinkComponentClearValueButton")	
}

function initUrlLinkFormComponentContainer($container,urlLinkRef) {
	setUrlLinkComponentLabel($container,urlLinkRef)
	initUrlLinkClearValueControl($container, urlLinkRef)
	initComponentHelpPopupButton($container, urlLinkRef)
}