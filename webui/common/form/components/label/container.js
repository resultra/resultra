

function labelControlContainerHTML() {
		
	return  '<div class="formLabelControl">' + 
					'<select class="form-control labelCompSelectionControl" multiple="multiple"></select>' +
				'</div>'
}

function labelContainerHTML(elementID)
{
	var containerHTML = ''+
		'<div class=" layoutContainer labelFormContainer">' +
			'<div class="form-group marginBottom0">'+
				'<div class="labelHeader">' +
					'<label class="componentLabel">New Label</label>' + 
					componentHelpPopupButtonHTML() +
				'</div>' +
				labelControlContainerHTML() +
			'</div>'+
		'</div>';
										
	return containerHTML
}

function labelControlFromLabelComponentContainer($labelContainer) {
	return $labelContainer.find(".labelCompSelectionControl")
}


function labelTablePopupViewContainerHTML() {
	var containerHTML = ''+
		'<div class=" layoutContainer labelTableCellContainer tagPopupContainer">' +
			'<div class="tagEditorHeader">' +
				'<button type="button" class="close closeTagEditorPopup" data-dismiss="modal" aria-hidden="true">x</button>' +
			'</div>' +
			'<div class="marginTop5">' +
				labelControlContainerHTML() +
			'</div>' +
		'</div>';									
	return containerHTML
	
}


function labelTableCellContainerHTML() {
	return '<div class="layoutContainer tagEditTableCell">' +
			'<div>' +
				'<a class="btn tagEditPopop"></a>'+
			'</div>' +
		'</div>'
}

function initTagContainerDimensions($container,overallWidth,overallHeight) {
	
	$container.css('height',overallHeight + "px")
	$container.css('width', overallWidth + 'px')
	
	var $labelContainer = $container.find(".select2-selection--multiple")
	
	var headerHeight = 35
	var labelControlHeightPx = (overallHeight - headerHeight) + "px"
	$labelContainer.css('min-height',labelControlHeightPx)
	$labelContainer.css('max-height',labelControlHeightPx)
	
}

function initTagTablePopupDimensions($container) {
	initTagContainerDimensions($container,250,150)
}

function initSelectionControlFormDimensions($container, labelRef) {
	
	var $labelControl = labelControlFromLabelComponentContainer($container)
	
	var overallHeight = labelRef.properties.geometry.sizeHeight
	var overallWidth = labelRef.properties.geometry.sizeWidth
	
	initTagContainerDimensions($labelControl,overallWidth,overallHeight)
	
	$container.css('height',overallHeight + "px")
	$container.css('width', overallWidth + 'px')
	
}

function initLabelSelectionControl($container, labelRef,labelWidth) {
		
	var $labelControl = labelControlFromLabelComponentContainer($container)
	
	$labelControl.select2({
		placeholder: "Enter tags", // TODO - Allow a property to configure the placeholder.
		width: labelWidth,
		tags:true,
		minimumInputLength: 1,
		maximumInputLength:32,
		tokenSeparators: [',']
	});

}	
	
	

function setLabelComponentLabel($label,label) {
	var $compLabel = $label.find('.componentLabel')
	
	setFormComponentLabel($compLabel,label.properties.fieldID,
			label.properties.labelFormat)	
	
}

function initLabelFormComponentContainer($container,label) {
	
	setLabelComponentLabel($container,label)
	initComponentHelpPopupButton($container, label)	
	
	var labelWidth = label.properties.geometry.sizeWidth - 15

	initLabelSelectionControl($container, label,labelWidth)
	
	//initSelectionControlFormDimensions($container, label,labelWidth)
		
}


