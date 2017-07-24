function componentHelpPopupButtonHTML() {

	// className is to uniquely identify the button with other HTML elements,
	// such that it can be found with jQuery's find() function.

	var buttonHTML = '<button tabindex="-1" class="btn btn-default btn-sm clearButton ' + 
			'componentHelpPopupButton pull-right' + 
			'"><span class="glyphicon glyphicon-question-sign"></span></button>'

	return buttonHTML
}

function updateComponentHelpPopupMsg($container, componentRef) {
	
	var $popupButton = $container.find(".componentHelpPopupButton")
	
	var popupMsgHTML = ""
	if(nonEmptyStringVal(componentRef.properties.helpPopupMsg)) {
		popupMsgHTML = '' +
			'<div class="inlineContent helpPopupContent">' + 
				formatInlineContentHTMLDisplay(componentRef.properties.helpPopupMsg) +
			'</div>'
		$popupButton.show()
	} else {
		$popupButton.hide()
	}
	
	// Bootstrap's popover uses an asyncrhonous call to destroy the popup. Therefore,
	// it's not possible to reliably destroy and re-initialize the popup just for puprposes
	// of changing the HTML to be displayed. So, to dynamically change the popup's HTML, we
	// can use a level of indirection to set & retrieve the message's HTML.
	$popupButton.data('componentPopupMsg',popupMsgHTML)
	
}

function initComponentHelpPopupButton($container, componentRef) {
	
	var $popupButton = $container.find(".componentHelpPopupButton")
	
	updateComponentHelpPopupMsg($container,componentRef)
	
	$popupButton.popover({
		html: 'true',
		delay: { "show": 100, "hide": 3000 },
		content: function() { return $popupButton.data('componentPopupMsg') },
		trigger: 'click hover',
		placement: 'auto top'
	})
	
}


