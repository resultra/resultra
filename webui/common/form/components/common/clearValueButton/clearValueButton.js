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


function smallClearButtonHTML(iconClass, className) {
	
	// className is to uniquely identify the button with other HTML elements,
	// such that it can be found with jQuery's find() function.
	
	var buttonHTML = '<button type="button" tabindex="-1" class="btn btn-default btn-sm clearButton input-group-addon' + 
			className + 
			'"><span class="'+iconClass+'"></span></button>'
	
	return buttonHTML
}

function smallClearDeleteButtonHTML(className) {
	return smallClearButtonHTML("glyphicon glyphicon-remove",className)
}
