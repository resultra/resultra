

function initCheckBoxFormatProps(params) {
	
	var $colorSchemeSelection = $('#adminCheckboxComponentColorSchemeSelection')
	$colorSchemeSelection.val(params.initialColorScheme)
	initSelectControlChangeHandler($colorSchemeSelection,function(newColorScheme) {
		params.setColorScheme(newColorScheme)
	})

	initCheckboxChangeHandler('#adminCheckboxComponentStrikethrough', 
			params.initialStrikethrough, function (newVal) {
			params.setStrikethrough(newVal)	
	})
	
}

