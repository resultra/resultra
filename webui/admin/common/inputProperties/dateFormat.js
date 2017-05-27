
function initDateFormatProperties(params) {
	
	var $formatSelection = $('#adminDateComponentFormatSelection')
	$formatSelection.val(params.initialFormat)
	initSelectControlChangeHandler($formatSelection,function(newFormat) {
		params.setFormat(newFormat)
	})
}
