function urlLinkContainerInputControl() {
		
	return '<div class="input-group">'+
				'<div class="formInputStaticInputContainer">' +
					'<a class="urlLinkDisplay"></a>' +
				'</div>' + 
				'<span class="input-group-addon urlLinkEditLinkButton">' +
             	   '<span class="glyphicon glyphicon-link"></span>' +
            	'</span>' +
				clearValueButtonHTML("urlLinkComponentClearValueButton") +
			'</div>'
}

function urlLinkEditPopupViewContainerHTML() {
	var containerHTML = ''+
		'<div class="urlLinkPopupContainer">' +
			'<div class="linkEditorHeader">' +
				'<button type="button" class="close closeLinkEditorPopup" data-dismiss="modal" aria-hidden="true">x</button>' +
			'</div>' +
			'<div class="marginTop5">' +
				'<label>Link address</label>' + 
				'<input type="text" name="symbol" class="urlLinkComponentInput form-control" placeholder="">'+
			'</div>' +
		'</div>'
	return containerHTML
	
}


function urlLinkContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class="layoutContainer urlLinkComponent urlLinkFormComponent">' +
				'<label>New link input</label>'+ componentHelpPopupButtonHTML() +
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


function calcURLMinTableCellColWidth(urlLinkRef,urlText) {
	
	var addrWidth = calcTextWidth(urlText)
	var paddingWidth = 10
	
	var addonWidth = 26
	if(formComponentIsReadOnly(urlLinkRef.properties.permissions)) { addonWidth +=26  }
	
	if (clearValueControlIsEnabled(urlLinkRef)) {
		addonWidth += 17
	}
	
	var unconstrainedWidth = addrWidth + paddingWidth + addonWidth
	
	return calcContrainedPxVal(unconstrainedWidth,300,400)
}


