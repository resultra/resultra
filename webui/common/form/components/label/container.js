

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


function labelTableCellContainerHTML() {
	var containerHTML = ''+
		'<div class=" layoutContainer labelTableCellContainer">' +
			labelControlContainerHTML() +
		'</div>';									
	return containerHTML
	
}

function initSelectionControlFormDimensions($container, labelRef) {
	
	var $labelControl = labelControlFromLabelComponentContainer($container)
	
	var overallHeight = labelRef.properties.geometry.sizeHeight
	var overallWidth = labelRef.properties.geometry.sizeWidth
	
	$container.css('height',overallHeight + "px")
	$container.css('width', overallWidth + 'px')
	
	var $labelContainer = $container.find(".select2-selection--multiple")
	
	var headerHeight = 35
	var labelControlHeightPx = (overallHeight - headerHeight) + "px"
	$labelContainer.css('min-height',labelControlHeightPx)
	$labelContainer.css('max-height',labelControlHeightPx)
	
}

function initLabelSelectionControl($container, labelRef,labelWidth) {
		
	var $labelControl = labelControlFromLabelComponentContainer($container)
	

	$labelControl.select2({
		placeholder: "Enter labels", // TODO - Allow a property to configure the placeholder.
		width: labelWidth,
		tags:true,
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
	initLabelSelectionControl($container, label)
	
	
	var labelWidth = labelRef.properties.geometry.sizeWidth - 15
	
	initSelectionControlFormDimensions($container, label,labelWidth)
	
	/* setElemFixedWidthFlexibleHeight($container,
				label.properties.geometry.sizeWidth) */
	
}


