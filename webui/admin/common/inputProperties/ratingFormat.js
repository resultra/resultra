

function initRatingIconProps(params) {
	var $iconSelection = $('#adminRatingComponentIconSelection')
	$iconSelection.val(params.initialIcon)
	initSelectControlChangeHandler($iconSelection,function(newIcon) {
		params.setIcon(newIcon)
	})
	
}
