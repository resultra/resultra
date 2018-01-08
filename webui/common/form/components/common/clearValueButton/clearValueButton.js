

function smallClearButtonHTML(iconClass, className) {
	
	// className is to uniquely identify the button with other HTML elements,
	// such that it can be found with jQuery's find() function.
	
	var buttonHTML = '<button type="button" tabindex="-1" class="btn btn-default btn-sm clearButton input-group-addon ' + 
			className + 
			'"><span class="'+iconClass+'"></span></button>'
	
	return buttonHTML
}

function smallClearDeleteButtonHTML(className) {
	return smallClearButtonHTML("glyphicon glyphicon-remove",className)
}

function clearValueButtonHTML(className) {
	
	// className is to uniquely identify the button with other HTML elements,
	// such that it can be found with jQuery's find() function.
	
	var buttonHTML = '' +
		'<span class="input-group-btn">' +
			'<button type="button" tabindex="-1" class="btn btn-default btn-sm clearValueButton ' + className + '">' + 
				'<small><i class="glyphicon glyphicon-remove"></i></small>' +
			'</button>' +
	    '</span'
	
	return buttonHTML
}

function smallClearComponentValHeaderButton(className) {
	
	var buttonHTML = '<button tabindex="-1" class="btn btn-default btn-sm clearButton pull-right ' + 
			className + '"><span class="glyphicon glyphicon-remove"></span></button>'

	return buttonHTML
}

function clearValueControlIsEnabled(controlRef) {
	var fieldID = controlRef.properties.fieldID
	var fieldRef = getFieldRef(fieldID)
	if(fieldRef.isCalcField) {
		return false
	}
	if(formComponentIsReadOnly(controlRef.properties.permissions)) {
		return false
	}
	if(controlRef.properties.clearValueSupported) {
		return true
	}
	return false
	
}


function initClearValueControl($container,controlRef,buttonClassName) {
	var $clearValueButton = $container.find(buttonClassName)
	
	if (clearValueControlIsEnabled(controlRef)) {
		$clearValueButton.css("display","")
	} else {
		$clearValueButton.css("display","none")
	}	
}