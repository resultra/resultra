
function clearFormComponentValidationPrompt($container) {
		$container.popover('destroy')
		
		$container.css({ border:"" })
		$container.find("label").css("color","")
		$container.css('background-color', '')
	
}

function setupFormComponentValidationPrompt($container,validationResult,validationCompleteCallback) {
	if (validationResult.validationSucceeded) {
		clearFormComponentValidationPrompt($container)
		
		validationCompleteCallback(true)
	} else {
		$container.popover({
			html: 'true',
			content: function() { return escapeHTML(validationResult.errorMsg) },
			trigger: 'hover',
			placement: 'auto left'
		})
		
		$container.css({ border:"2px red" })
		$container.find("label").css("color","red")
		$container.css('background-color', 'rgba(255,0,0,0.1)')
		
		validationCompleteCallback(false)
	}
}